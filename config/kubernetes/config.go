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

package kubernetes

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

// Configure configures kubernetes group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kubernetes_cluster", func(r *config.Resource) {
		r.Kind = "KubernetesCluster"
		r.ShortGroup = "containerservice"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kubelet_identity"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"default_node_pool[*].pod_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},
			"default_node_pool[*].vnet_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},
			"addon_profile[*].ingress_application_gateway[*].subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_kubernetes_cluster_node_pool", func(r *config.Resource) {
		r.Kind = "KubernetesClusterNodePool"
		r.ShortGroup = "containerservice"
		r.References = config.References{
			"kubernetes_cluster_id": config.Reference{
				Type: "KubernetesCluster",
			},
			"pod_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},
			"vnet_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
	})
}
