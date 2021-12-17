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

package iothub

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

// Configure configures iothub group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_iothub", func(r *config.Resource) {
		r.Kind = "IOTHub"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"endpoint.resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"endpoint.container_name": config.Reference{
				Type: rconfig.APISPackagePath + "/storage/v1alpha1.Container",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "IotHubs", "name")
	})

	p.AddResourceConfigurator("azurerm_iothub_consumer_group", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"iothub_name": config.Reference{
				Type: "IOTHub",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/eventHubEndpoints/events/ConsumerGroups/group1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "IotHubs", "iothub_name",
			"eventHubEndpoints", "eventhub_endpoint_name", "ConsumerGroups", "name")
	})

	p.AddResourceConfigurator("azurerm_iothub_dps", func(r *config.Resource) {
		r.Kind = "DPS"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "provisioningServices", "name")
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_certificate", func(r *config.Resource) {
		r.Kind = "DPSCertificate"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"iot_dps_name": config.Reference{
				Type: "DPS",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example/certificates/example
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "provisioningServices", "iot_dps_name",
			"certificates", "name")
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_shared_access_policy", func(r *config.Resource) {
		r.Kind = "DPSSharedAccessPolicy"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"iothub_dps_name": config.Reference{
				Type: "DPS",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/dps1/keys/shared_access_policy1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "provisioningServices", "iothub_dps_name",
			"keys", "name")
	})

	p.AddResourceConfigurator("azurerm_iothub_shared_access_policy", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"iothub_name": config.Reference{
				Type: "IOTHub",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/IotHubKeys/shared_access_policy1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Devices", "IotHubs", "iothub_name",
			"IotHubKeys", "name")
	})
}
