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
	"os"
	"path/filepath"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"gopkg.in/alecthomas/kingpin.v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	tjcontroller "github.com/crossplane-contrib/terrajet/pkg/controller"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	"github.com/crossplane-contrib/provider-tf-azure/apis"
	genConfig "github.com/crossplane-contrib/provider-tf-azure/cmd/generator/config"
	"github.com/crossplane-contrib/provider-tf-azure/internal/clients"
	"github.com/crossplane-contrib/provider-tf-azure/internal/controller"
)

func main() {
	var (
		app              = kingpin.New(filepath.Base(os.Args[0]), "Azure support for Crossplane.").DefaultEnvars()
		debug            = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod       = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		leaderElection   = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()
		enabledAPIs      = app.Flag("enabled-apis", "Enabled APIs for the provider.").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-tf-azure"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String())

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-tf-azure",
		SyncPeriod:       syncPeriod,
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")

	genConfig.SetResourceConfigurations()
	ws := terraform.NewWorkspaceStore(log)
	rl := ratelimiter.NewGlobal(ratelimiter.DefaultGlobalRPS)
	var apiArr []string
	if enabledAPIs != nil && strings.TrimSpace(*enabledAPIs) != "" {
		apiArr = strings.Split(strings.TrimSpace(*enabledAPIs), ",")
	}

	// register API groups & start controllers for enabled APIs
	schemeBuilders := apis.GetRegisterMap()
	for gvk, setup := range controller.GetSetupMap() {
		enabled, err := tjcontroller.IsAPIEnabled(gvk, apiArr)
		kingpin.FatalIfError(err, "Cannot setup Azure controller for: %s", gvk.String())
		if !enabled {
			continue
		}
		builder, ok := schemeBuilders[gvk.GroupVersion()]
		if !ok {
			kingpin.Fatalf("Scheme builder not found for: %s", gvk.String())
		}
		kingpin.FatalIfError(builder.AddToScheme(mgr.GetScheme()), "Cannot add Azure API to scheme for: %s", gvk.GroupVersion().String())
		kingpin.FatalIfError(
			setup(mgr, log, rl, clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion), ws, 1),
			"Cannot setup controller for: ", gvk.String())
	}

	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
