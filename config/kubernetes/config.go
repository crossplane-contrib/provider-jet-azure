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
	"fmt"

	"github.com/crossplane-contrib/provider-tf-azure/config/common"
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kubernetes_cluster", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kubelet_identity"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"default_node_pool[*].pod_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
			"default_node_pool[*].vnet_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
			"addon_profile[*].ingress_application_gateway[*].subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.ContainerService", "managedClusters", "name")
	})

	p.AddResourceConfigurator("azurerm_kubernetes_cluster_node_pool", func(r *config.Resource) {
		r.References = config.References{
			"kubernetes_cluster_id": config.Reference{
				Type: "Cluster",
			},
			"pod_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
			"vnet_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/virtual/v1alpha1.Subnet",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			clusterID, ok := parameters["kubernetes_cluster_id"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("%s/agentPools/%s", clusterID, name)
		}
	})
}
