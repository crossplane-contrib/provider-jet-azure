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

package postgresql

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_postgresql_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_configuration", func(r *config.Resource) {
		r.References = config.References{
			"server_id": config.Reference{
				Type: "FlexibleServer",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "Server",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_active_directory_administrator", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): this may have to be a reference to the server's resource group
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "Server",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_database", func(r *config.Resource) {
		r.References = config.References{
			"server_id": config.Reference{
				Type: "FlexibleServer",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_firewall_rule", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "Server",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.References = config.References{
			"server_id": config.Reference{
				Type: "FlexibleServer",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"delegated_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_virtual_network_rule", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "Server",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_server_key", func(r *config.Resource) {
		r.References = config.References{
			"server_id": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_postgresql_configuration", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "Server",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})
}
