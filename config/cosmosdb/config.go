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
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cosmosdb_sql_container", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
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
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_collection", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
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
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_keyspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_table", func(r *config.Resource) {
		r.References = config.References{
			"cassandra_keyspace_id": config.Reference{
				Type: "CassandraKeyspace",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_graph", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
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
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_function", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				Type: "SqlContainer",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_stored_procedure", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
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
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_database", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_table", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_account", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_notebook_workspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "Account",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_trigger", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				Type: "SqlContainer",
			},
		}
		r.UseAsync = true
	})
}
