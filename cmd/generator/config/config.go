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

package config

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func ExtractResourceName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		p, err := fieldpath.PaveObject(mr)
		if err != nil {
			return ""
		}

		name, err := p.GetString("spec.forProvider.name")
		if err != nil {
			return ""
		}
		return name
	}
}

func SetResourceConfigurations() {
	config.Store.SetForResource("azurerm_resource_group", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("azurerm_postgresql_server", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"SslEnforcement", "StorageProfile"},
		},
		UseAsync: true,
	})

	config.Store.SetForResource("azurerm_virtual_network", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup",
				Extractor: "github.com/crossplane-contrib/provider-tf-azure/cmd/generator/config.ExtractResourceName()",
			},
		},
	})

	config.Store.SetForResource("azurerm_kubernetes_cluster", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"KubeletIdentity"},
		},
		UseAsync: true,
	})
}
