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

package redis

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

// Configure configures redis group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_redis_cache", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"subnet_id": config.Reference{
				Type:      rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Cache",
			"Redis", "name",
		)
	})
	p.AddResourceConfigurator("azurerm_redis_firewall_rule", func(r *config.Resource) {
		r.References = config.References{
			"redis_cache_name": config.Reference{
				Type: "Cache",
			},
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1/firewallRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Cache",
			"Redis", "redis_cache_name",
			"firewallRules", "name",
		)
	})
	p.AddResourceConfigurator("azurerm_redis_linked_server", func(r *config.Resource) {
		r.References = config.References{
			"linked_redis_cache_id": config.Reference{
				Type:      "Cache",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"target_redis_cache_name": config.Reference{
				Type: "Cache",
			},
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.IdentifierFromProvider
	})
}
