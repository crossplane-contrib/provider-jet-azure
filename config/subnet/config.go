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

package subnet

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

// Configure configures subnet group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_subnet_nat_gateway_association", func(r *config.Resource) {
		r.References = config.References{
			"subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_network_security_group_association", func(r *config.Resource) {
		r.References = config.References{
			"subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_service_endpoint_storage_policy", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_route_table_association", func(r *config.Resource) {
		r.References = config.References{
			"subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})
}
