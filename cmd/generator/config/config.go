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
)

// APISPackagePath is the package path for generated APIs root package
const APISPackagePath = "github.com/crossplane-contrib/provider-tf-azure/apis"

// SetResourceConfigurations sets the Terraformed resource configurations
// for the generated resources
func SetResourceConfigurations() {
	config.Store.SetForResource("azurerm_resource_group", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("azurerm_kubernetes_cluster", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"KubeletIdentity"},
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_kubernetes_cluster_node_pool", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"kubernetes_cluster_id": config.Reference{
				Type: "KubernetesCluster",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_network", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
	})
	config.Store.SetForResource("azurerm_postgresql_server", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"SslEnforcement", "StorageProfile"},
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_sql_container", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "CosmosdbSqlDatabase",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub_connection", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"virtual_hub_id": config.Reference{
				Type: "VirtualHub",
			},
			"remote_virtual_network_id": config.Reference{
				Type: "VirtualNetwork",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_network_gateway", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_mongo_collection", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "CosmosdbMongoDatabase",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_network_peering", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"virtual_network_name": config.Reference{
				Type:      "VirtualNetwork",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"remote_virtual_network_id": config.Reference{
				Type: "VirtualNetwork",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_desktop_application", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_cassandra_keyspace", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_desktop_host_pool", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub_bgp_connection", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"virtual_hub_id": config.Reference{
				Type: "VirtualHub",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_flexible_server_configuration", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"server_id": config.Reference{
				Type: "PostgresqlFlexibleServer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_database", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "PostgresqlServer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_cassandra_table", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"cassandra_keyspace_id": config.Reference{
				Type: "CosmosdbCassandraKeyspace",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_gremlin_graph", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "CosmosdbGremlinDatabase",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_network_gateway_connection", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"virtual_network_gateway_id": config.Reference{
				Type: "VirtualNetworkGateway",
			},
			"peer_virtual_network_gateway_id": config.Reference{
				Type: "VirtualNetworkGateway",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_sql_function", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"container_id": config.Reference{
				Type: "CosmosdbSqlContainer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_active_directory_administrator", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			// TODO(aru): this may have to be a reference to the server's resource group
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "PostgresqlServer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_flexible_server_database", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"server_id": config.Reference{
				Type: "PostgresqlFlexibleServer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_sql_stored_procedure", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"database_name": config.Reference{
				Type:      "CosmosdbSqlDatabase",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"container_name": config.Reference{
				Type:      "CosmosdbSqlContainer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_gremlin_database", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_mongo_database", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_firewall_rule", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "PostgresqlServer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub_security_partner_provider", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"virtual_hub_id": config.Reference{
				Type: "VirtualHub",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_sql_database", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_flexible_server_firewall_rule", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"server_id": config.Reference{
				Type: "PostgresqlFlexibleServer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_flexible_server", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"SslEnforcement", "StorageProfile"},
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_virtual_network_rule", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "PostgresqlServer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_table", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_account", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_server_key", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"server_id": config.Reference{
				Type: "PostgresqlServer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub_route_table", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"virtual_hub_id": config.Reference{
				Type: "VirtualHub",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_desktop_workspace", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_resource_group_template_deployment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      "ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_resource_group_policy_assignment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_id": config.Reference{
				Type: "ResourceGroup",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_postgresql_configuration", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"server_name": config.Reference{
				Type:      "PostgresqlServer",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub_ip", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"virtual_hub_id": config.Reference{
				Type: "VirtualHub",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_notebook_workspace", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"account_name": config.Reference{
				Type:      "CosmosdbAccount",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_cosmosdb_sql_trigger", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"container_id": config.Reference{
				Type: "CosmosdbSqlContainer",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_wan", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
		},
		UseAsync: true,
	})
	config.Store.SetForResource("azurerm_virtual_hub", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"resource_group_name": config.Reference{
				Type:      APISPackagePath + "/resource/v1alpha1.ResourceGroup",
				Extractor: APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			"virtual_wan_id": config.Reference{
				Type: "VirtualWan",
			},
		},
		UseAsync: true,
	})
}
