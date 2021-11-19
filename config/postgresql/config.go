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
	"context"
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

const (
	errFmtNoAttribute    = `"attribute not found: %s`
	errFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// Configure configures postgresql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_postgresql_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL", "servers", "name")
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_configuration", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"server_id": config.Reference{
				Type: "FlexibleServer",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_postgresql_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/databases/database1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"databases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_active_directory_administrator", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): this may have to be a reference to the server's resource group
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["login"] = name
			},
			OmittedFields:     []string{"login"},
			GetExternalNameFn: common.GetNameFromFullyQualifiedID,
			// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/administrators/activeDirectory
			GetIDFn: common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
				"servers", "server_name",
				"administrators", "login",
			),
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_database", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"server_id": config.Reference{
				Type: "FlexibleServer",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/databases/database1
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			serverID, ok := parameters["server_id"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "server_id")
			}
			serverIDStr, ok := serverID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "server_id")
			}
			return fmt.Sprintf("%s/databases/%s", serverIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_firewall_rule", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/firewallRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"firewallRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"server_id": config.Reference{
				Type: "FlexibleServer",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			serverID, ok := parameters["server_id"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "server_id")
			}
			serverIDStr, ok := serverID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "server_id")
			}
			return fmt.Sprintf("%s/firewallRules/%s", serverIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"delegated_subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"flexibleServers", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_virtual_network_rule", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"server_name": config.Reference{
				Type: "Server",
			},
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"subnet_id": config.Reference{
				Type: rconfig.APISPackagePath + "/network/v1alpha1.Subnet",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/virtualNetworkRules/vnetrulename
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"virtualNetworkRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_server_key", func(r *config.Resource) {
		r.References = config.References{
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"server_id": config.Reference{
				Type: "Server",
			},*/
		}
		r.UseAsync = true
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
		// NOTE(muvaf): There is no actual name in the arguments that we can use.
		// It's all calculated by Azure.
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_postgresql_configuration", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL", "servers", "server_name", "configurations", "name")
	})
}
