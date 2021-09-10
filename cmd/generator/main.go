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
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
)

// Constants to use in generated artifacts.
const (
	modulePath  = "github.com/crossplane-contrib/provider-tf-azure"
	groupSuffix = ".azure.tf.crossplane.io"
)

var (
	// We expect a function called "ProviderConfigBuilder" of type
	// "ProviderConfigFn" (https://github.com/crossplane-contrib/terrajet/blob/4246657031f181fdaf0de83e83ceaa8735307180/pkg/terraform/controller.go#L33)
	// available at this path
	providerConfigBuilderPath = filepath.Join(modulePath, "internal", "clients")
)
var skipList = map[string]struct{}{
	"azurerm_mssql_server_extended_auditing_policy": {},
	// group prefix collision
	"azurerm_api_management_group":              {},
	"azurerm_api_management_product_group":      {},
	"azurerm_dedicated_host_group":              {},
	"azurerm_storage_sync_group":                {},
	"azurerm_virtual_desktop_application_group": {},
	// generated name too long
	"azurerm_network_interface_application_gateway_backend_address_pool_association": {},
	"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection":  {},
}

var includeList = map[string]struct{}{
	"azurerm_virtual_hub_connection":                                  {},
	"azurerm_virtual_machine_configuration_policy_assignment":         {},
	"azurerm_virtual_hub_bgp_connection":                              {},
	"azurerm_virtual_hub_security_partner_provider":                   {},
	"azurerm_virtual_machine_data_disk_attachment":                    {},
	"azurerm_virtual_wan":                                             {},
	"azurerm_virtual_hub_ip":                                          {},
	"azurerm_virtual_desktop_host_pool":                               {},
	"azurerm_virtual_desktop_application":                             {},
	"azurerm_virtual_machine_extension":                               {},
	"azurerm_virtual_machine":                                         {},
	"azurerm_virtual_network_gateway":                                 {},
	"azurerm_virtual_network_peering":                                 {},
	"azurerm_virtual_machine_scale_set_extension":                     {},
	"azurerm_virtual_desktop_application_group":                       {},
	"azurerm_virtual_hub":                                             {},
	"azurerm_virtual_network_dns_servers":                             {},
	"azurerm_virtual_network_gateway_connection":                      {},
	"azurerm_virtual_desktop_workspace_application_group_association": {},
	"azurerm_virtual_hub_route_table":                                 {},
	"azurerm_virtual_machine_scale_set":                               {},
	"azurerm_virtual_network":                                         {},
	"azurerm_virtual_desktop_workspace":                               {},
}

// "make prepare.azurerm" should be run before running this generator pipeline
func main() { // nolint:gocyclo
	// Cyclomatic complexity of this function is above our goal of 10,
	// and it establishes a Terrajet code generation pipeline that's very similar
	// to other Terrajet based providers.
	// delete API dirs
	deleteGenDirs("apis", map[string]struct{}{
		"v1alpha1": {},
	})
	// delete controller dirs
	deleteGenDirs("internal/controller", map[string]struct{}{
		"config": {},
	})

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
			if _, ok := includeList[name]; !ok {
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
		ctrlGen := pipeline.NewControllerGenerator(wd, modulePath, strings.ToLower(group)+groupSuffix, providerConfigBuilderPath)

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
			if err := crdGen.Generate(version, kind, resources[name]); err != nil {
				panic(errors.Wrap(err, "cannot generate crd"))
			}
			if err := tfGen.Generate(version, kind, name, "id"); err != nil {
				panic(errors.Wrap(err, "cannot generate terraformed"))
			}
			ctrlPkgPath, err := ctrlGen.Generate(versionGen.Package().Path(), kind)
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
