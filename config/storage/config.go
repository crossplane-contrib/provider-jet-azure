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

package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/terrajet/pkg/config"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

// Configure configures storage group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_storage_account", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + common.ResourceGroupReferencePath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Storage",
			"storageAccounts", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_storage_blob", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + common.ResourceGroupReferencePath,
			},
			"storage_account_name": config.Reference{
				Type: "Account",
			},
			"storage_container_name": config.Reference{
				Type: "Container",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		// https://example.blob.core.windows.net/container/blob.vhd
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"]
			if !ok {
				return "", errors.Errorf(common.ErrFmtNoAttribute, "id")
			}
			idStr, ok := id.(string)
			if !ok {
				return "", errors.Errorf(common.ErrFmtUnexpectedType, "id")
			}
			idStr = strings.TrimSuffix(idStr, ".blob.core.windows.net/container/blob.vhd")
			return strings.TrimPrefix(idStr, "https://"), nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			return fmt.Sprintf("https://%s.blob.core.windows.net/container/blob.vhd", name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_storage_container", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"storage_account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		// https://example.blob.core.windows.net/container
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"]
			if !ok {
				return "", errors.Errorf(common.ErrFmtNoAttribute, "id")
			}
			idStr, ok := id.(string)
			if !ok {
				return "", errors.Errorf(common.ErrFmtUnexpectedType, "id")
			}
			idStr = strings.TrimSuffix(idStr, ".blob.core.windows.net/container")
			return strings.TrimPrefix(idStr, "https://"), nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			return fmt.Sprintf("https://%s.blob.core.windows.net/container", name), nil
		}
	})
}
