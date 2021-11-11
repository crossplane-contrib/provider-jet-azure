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

package cosmosdb

import (
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-tf-azure/config/common"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cosmosdb_sql_container", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "SqlDatabase",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/sqlDatabases/db1/containers/container1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			dbName, ok := parameters["database_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlDatabases/%s/containers/%s", subID, rg, account, dbName, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_collection", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "MongoDatabase",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1/collections/collection1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			dbName, ok := parameters["database_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/mongodbDatabases/%s/collections/%s", subID, rg, account, dbName, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_keyspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/cassandraKeyspaces/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_table", func(r *config.Resource) {
		r.References = config.References{
			"cassandra_keyspace_id": config.Reference{
				Type: "CassandraKeyspace",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1/tables/table1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			ksID, ok := parameters["cassandra_keyspace_id"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/cassandraKeyspaces/%s/tables/%s", subID, rg, account, ksID, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_graph", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "GremlinDatabase",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1/graphs/graphs1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			dbName, ok := parameters["database_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/gremlinDatabases/%s/graphs/%s", subID, rg, account, dbName, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_function", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				Type: "SqlContainer",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/userDefinedFunctions/userDefinedFunction1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			containerID, ok := parameters["container_id"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("%s/userDefinedFunctions/%s", containerID, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_stored_procedure", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "SqlDatabase",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"container_name": config.Reference{
				Type:      "SqlContainer",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1/containers/c1/storedProcedures/sp1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			dbName, ok := parameters["database_name"].(string)
			if !ok {
				return ""
			}
			containerName, ok := parameters["container_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlDatabases/%s/containers/%s/storedProcedures/%s", subID, rg, account, dbName, containerName, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_database", func(r *config.Resource) {

		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/gremlinDatabases/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/mongodbDatabases/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlDatabases/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_table", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/tables/table1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/tables/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_account", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("/Microsoft.DocumentDB/databaseAccounts")
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_notebook_workspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/notebookWorkspaces/notebookWorkspace1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			subID, ok := providerConfig["subscriptionId"].(string)
			if !ok {
				return ""
			}
			rg, ok := parameters["resource_group_name"].(string)
			if !ok {
				return ""
			}
			account, ok := parameters["account_name"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/notebookWorkspaces/%s", subID, rg, account, name)
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_trigger", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				Type: "SqlContainer",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/triggers/trigger1
		r.ExternalName.GetIDFn = func(name string, parameters map[string]interface{}, providerConfig map[string]interface{}) string {
			containerID, ok := parameters["container_id"].(string)
			if !ok {
				return ""
			}
			return fmt.Sprintf("%s/triggers/%s", containerID, name)
		}
	})
}
