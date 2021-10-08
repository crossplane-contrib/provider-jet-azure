/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/pipeline"

	genConfig "github.com/crossplane-contrib/provider-tf-azure/cmd/generator/config"
)

// Constants to use in generated artifacts.
const (
	modulePath  = "github.com/crossplane-contrib/provider-tf-azure"
	groupSuffix = ".azure.tf.crossplane.io"
)

var skipList = map[string]struct{}{
	"azurerm_mssql_server_extended_auditing_policy": {},
	// group prefix collision
	"azurerm_api_management_group":              {},
	"azurerm_api_management_product_group":      {},
	"azurerm_dedicated_host_group":              {},
	"azurerm_storage_sync_group":                {},
	"azurerm_virtual_desktop_application_group": {},
	// associated with non-generated
	"azurerm_virtual_desktop_workspace_application_group_association": {},
	// generated name too long
	"azurerm_network_interface_application_gateway_backend_address_pool_association": {},
	"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection":  {},
	// deprecated
	"azurerm_virtual_machine_scale_set":                       {},
	"azurerm_virtual_machine_configuration_policy_assignment": {},
	"azurerm_virtual_machine":                                 {},
	"azurerm_virtual_machine_extension":                       {},
	"azurerm_virtual_machine_data_disk_attachment":            {},
	"azurerm_virtual_machine_scale_set_extension":             {},
	// doc not found in Terraform Azurerm provider
	"azurerm_virtual_network_dns_servers": {},
}

var includeList = []string{
	"azurerm_virtual_.+",
	"azurerm_kubernetes_.+",
	"azurerm_postgresql_.+",
	"azurerm_cosmosdb_.+",
	"azurerm_resource_group",
}

// "make prepare.azurerm" should be run before running this generator pipeline
func main() { // nolint:gocyclo
	// Cyclomatic complexity of this function is above our goal of 10,
	// and it establishes a Terrajet code generation pipeline that's very similar
	// to other Terrajet based providers.
	// delete API dirs
	deleteGenDirs("apis", map[string]struct{}{
		"v1alpha1": {},
		"rconfig":  {},
	})
	// delete controller dirs
	deleteGenDirs("internal/controller", map[string]struct{}{
		"config": {},
	})

	genConfig.SetResourceConfigurations()
	wd, err := os.Getwd()
	if err != nil {
		panic(errors.Wrap(err, "cannot get working directory"))
	}
	groups := map[string]map[string]*schema.Resource{}
	useIncludeList := os.Getenv("USE_INCLUDE_LIST") == "true"
	for name, resource := range xpprovider.Provider().ResourcesMap {
		if len(resource.Schema) == 0 {
			// There are resources with no schema, that we will address later.
			fmt.Printf("Skipping resource %s because it has no schema\n", name)
			continue
		}
		if _, ok := skipList[name]; ok {
			continue
		}
		if useIncludeList {
			match := false
			for _, r := range includeList {
				ok, err := regexp.MatchString(r, name)
				if err != nil {
					panic(errors.Wrap(err, "cannot match regular expression"))
				}
				if ok {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		fmt.Printf("Generating code for resource: %s\n", name)

		words := strings.Split(name, "_")
		groupName := words[1]
		if len(groups[groupName]) == 0 {
			groups[groupName] = map[string]*schema.Resource{}
		}
		groups[groupName][name] = resource
	}
	count := 0
	versionPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-azure/apis/v1alpha1",
	}
	controllerPkgList := []string{
		"github.com/crossplane-contrib/provider-tf-azure/internal/controller/config",
	}
	for group, resources := range groups {
		version := "v1alpha1"
		versionGen := pipeline.NewVersionGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix, version)

		crdGen := pipeline.NewCRDGenerator(versionGen.Package(), versionGen.DirectoryPath(), strings.ToLower(group)+groupSuffix, "azure")
		tfGen := pipeline.NewTerraformedGenerator(versionGen.Package(), versionGen.DirectoryPath())
		ctrlGen := pipeline.NewControllerGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix)

		keys := make([]string, len(resources))
		i := 0
		for k := range resources {
			keys[i] = k
			i++
		}
		sort.Strings(keys)

		for _, name := range keys {
			// We don't want Azurerm prefix in all kinds.
			kind := strings.TrimPrefix(strcase.ToCamel(name), "Azurerm")
			resource := resources[name]
			r := config.NewResource(version, kind, name)
			if err = r.OverrideConfig(config.Store.GetForResource(name)); err != nil {
				panic(errors.Wrap(err, "cannot override config"))
			}
			if err := crdGen.Generate(r, resource); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
			if err := tfGen.Generate(r); err != nil {
				panic(errors.Wrap(err, "cannot generate terraformed"))
			}
			ctrlPkgPath, err := ctrlGen.Generate(r, versionGen.Package().Path())
			if err != nil {
				panic(errors.Wrap(err, "cannot generate controller"))
			}
			controllerPkgList = append(controllerPkgList, ctrlPkgPath)
			count++
		}

		if err := versionGen.Generate(); err != nil {
			panic(errors.Wrap(err, "cannot generate version files"))
		}
		versionPkgList = append(versionPkgList, versionGen.Package().Path())
	}

	if err := pipeline.NewRegisterGenerator(wd, modulePath).Generate(versionPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate register file"))
	}
	if err := pipeline.NewSetupGenerator(wd, modulePath).Generate(controllerPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate setup file"))
	}
	if err := exec.Command("bash", "-c", "goimports -w $(find apis -iname 'zz_*')").Run(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for apis folder"))
	}
	if err := exec.Command("bash", "-c", "goimports -w $(find internal -iname 'zz_*')").Run(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for internal folder"))
	}
	fmt.Printf("\nGenerated %d resources!\n", count)
}

// delete API subdirs for a clean start
func deleteGenDirs(rootDir string, keepMap map[string]struct{}) {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		panic(errors.Wrapf(err, "cannot list files under %s", rootDir))
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		if _, ok := keepMap[f.Name()]; ok {
			continue
		}
		removeDir := filepath.Join(rootDir, f.Name())
		if err := os.RemoveAll(removeDir); err != nil {
			panic(errors.Wrapf(err, "cannot remove API dir: %s", removeDir))
		}
	}
}
