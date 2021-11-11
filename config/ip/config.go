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

package ip

import (
	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-tf-azure/config/common"
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ip_group", func(r *config.Resource) {
		r.Kind = "IPGroup"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/ipGroups/myIpGroup
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("/Microsoft.Network/ipGroups")
	})
}
