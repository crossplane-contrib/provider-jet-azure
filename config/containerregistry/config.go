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

package containerregistry

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

// Configure configures containerregistry group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_container_registry_token", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"scope_map_id": config.Reference{
				Type:      "ScopeMap"
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.ContainerRegistry/registries/myregistry/tokens/mytok1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.ContainerRegistry",
			"registries", "container_registry_name",
			"tokens", "name",
		)
	})
}
