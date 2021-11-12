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

package resource

import (
	"context"
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-tf-azure/config/common"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

const (
	errFmtNoAttribute    = `"attribute not found: %s`
	errFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_resource_group", func(r *config.Resource) {
		r.Kind = "ResourceGroup"
		r.ShortGroup = ""
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, _ map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			subID, ok := providerConfig["subscriptionId"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "subscriptionId")
			}
			subIDStr, ok := subID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "subscriptionId")
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_resource_group_template_deployment", func(r *config.Resource) {
		r.Kind = "ResourceGroupTemplateDeployment"
		r.ShortGroup = "resources"
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true

		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Resources/deployments/template1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Resources",
			"deployments", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_resource_group_policy_assignment", func(r *config.Resource) {
		r.Kind = "ResourceGroupPolicyAssignment"
		r.ShortGroup = "authorization"
		r.References = config.References{
			"resource_group_id": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Authorization",
			"policyAssignments", "name",
		)
	})
}
