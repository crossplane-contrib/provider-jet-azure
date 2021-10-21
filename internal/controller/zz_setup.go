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

package controller

import (
	tjcontroller "github.com/crossplane-contrib/terrajet/pkg/controller"
	"k8s.io/apimachinery/pkg/runtime/schema"

	config "github.com/crossplane-contrib/provider-tf-azure/internal/controller/config"
	cosmosdbaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbaccount"
	cosmosdbcassandrakeyspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbcassandrakeyspace"
	cosmosdbcassandratable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbcassandratable"
	cosmosdbgremlindatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbgremlindatabase"
	cosmosdbgremlingraph "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbgremlingraph"
	cosmosdbmongocollection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbmongocollection"
	cosmosdbmongodatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbmongodatabase"
	cosmosdbnotebookworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbnotebookworkspace"
	cosmosdbsqlcontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlcontainer"
	cosmosdbsqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqldatabase"
	cosmosdbsqlfunction "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlfunction"
	cosmosdbsqlstoredprocedure "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlstoredprocedure"
	cosmosdbsqltrigger "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqltrigger"
	cosmosdbtable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cosmosdbtable"
	kubernetescluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/kubernetescluster"
	kubernetesclusternodepool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/kubernetesclusternodepool"
	loadbalancer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/loadbalancer"
	postgresqlactivedirectoryadministrator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlactivedirectoryadministrator"
	postgresqlconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlconfiguration"
	postgresqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqldatabase"
	postgresqlfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlfirewallrule"
	postgresqlflexibleserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserver"
	postgresqlflexibleserverconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverconfiguration"
	postgresqlflexibleserverdatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverdatabase"
	postgresqlflexibleserverfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverfirewallrule"
	postgresqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlserver"
	postgresqlserverkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlserverkey"
	postgresqlvirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlvirtualnetworkrule"
	resourcegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegroup"
	resourcegrouppolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegrouppolicyassignment"
	resourcegrouptemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegrouptemplatedeployment"
	sqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlserver"
	storageaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageaccount"
	storageblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageblob"
	storagecontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagecontainer"
	subnet "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnet"
	subnetnatgatewayassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetserviceendpointstoragepolicy"
	virtualdesktopapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopapplication"
	virtualdesktophostpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktophostpool"
	virtualdesktopworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopworkspace"
	virtualhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhub"
	virtualhubbgpconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubbgpconnection"
	virtualhubconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubconnection"
	virtualhubip "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubip"
	virtualhubroutetable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubroutetable"
	virtualhubsecuritypartnerprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubsecuritypartnerprovider"
	virtualnetwork "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetwork"
	virtualnetworkgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkpeering"
	virtualwan "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualwan"
)

// GetSetupMap returns a map from GVKs to
// the respective controller setup functions.
func GetSetupMap() map[schema.GroupVersionKind]tjcontroller.SetupFunc {
	m := map[schema.GroupVersionKind]tjcontroller.SetupFunc{
		schema.GroupVersionKind{Group: "azure.tf.crossplane.io", Version: "v1alpha1", Kind: "ProviderConfig"}:                                    config.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbAccount"}:                          cosmosdbaccount.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbCassandraKeyspace"}:                cosmosdbcassandrakeyspace.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbCassandraTable"}:                   cosmosdbcassandratable.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbGremlinDatabase"}:                  cosmosdbgremlindatabase.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbGremlinGraph"}:                     cosmosdbgremlingraph.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbMongoCollection"}:                  cosmosdbmongocollection.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbMongoDatabase"}:                    cosmosdbmongodatabase.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbNotebookWorkspace"}:                cosmosdbnotebookworkspace.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbSqlContainer"}:                     cosmosdbsqlcontainer.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbSqlDatabase"}:                      cosmosdbsqldatabase.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbSqlFunction"}:                      cosmosdbsqlfunction.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbSqlStoredProcedure"}:               cosmosdbsqlstoredprocedure.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbSqlTrigger"}:                       cosmosdbsqltrigger.Setup,
		schema.GroupVersionKind{Group: "cosmosdb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "CosmosdbTable"}:                            cosmosdbtable.Setup,
		schema.GroupVersionKind{Group: "kubernetes.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "KubernetesCluster"}:                      kubernetescluster.Setup,
		schema.GroupVersionKind{Group: "kubernetes.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "KubernetesClusterNodePool"}:              kubernetesclusternodepool.Setup,
		schema.GroupVersionKind{Group: "lb.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "LoadBalancer"}:                                   loadbalancer.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlActiveDirectoryAdministrator"}: postgresqlactivedirectoryadministrator.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlConfiguration"}:                postgresqlconfiguration.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlDatabase"}:                     postgresqldatabase.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlFirewallRule"}:                 postgresqlfirewallrule.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlFlexibleServer"}:               postgresqlflexibleserver.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlFlexibleServerConfiguration"}:  postgresqlflexibleserverconfiguration.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlFlexibleServerDatabase"}:       postgresqlflexibleserverdatabase.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlFlexibleServerFirewallRule"}:   postgresqlflexibleserverfirewallrule.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlServer"}:                       postgresqlserver.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlServerKey"}:                    postgresqlserverkey.Setup,
		schema.GroupVersionKind{Group: "postgresql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "PostgresqlVirtualNetworkRule"}:           postgresqlvirtualnetworkrule.Setup,
		schema.GroupVersionKind{Group: "resource.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "ResourceGroup"}:                            resourcegroup.Setup,
		schema.GroupVersionKind{Group: "resource.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "ResourceGroupPolicyAssignment"}:            resourcegrouppolicyassignment.Setup,
		schema.GroupVersionKind{Group: "resource.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "ResourceGroupTemplateDeployment"}:          resourcegrouptemplatedeployment.Setup,
		schema.GroupVersionKind{Group: "sql.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "SqlServer"}:                                     sqlserver.Setup,
		schema.GroupVersionKind{Group: "storage.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "StorageAccount"}:                            storageaccount.Setup,
		schema.GroupVersionKind{Group: "storage.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "StorageBlob"}:                               storageblob.Setup,
		schema.GroupVersionKind{Group: "storage.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "StorageContainer"}:                          storagecontainer.Setup,
		schema.GroupVersionKind{Group: "subnet.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "Subnet"}:                                     subnet.Setup,
		schema.GroupVersionKind{Group: "subnet.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "SubnetNatGatewayAssociation"}:                subnetnatgatewayassociation.Setup,
		schema.GroupVersionKind{Group: "subnet.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "SubnetNetworkSecurityGroupAssociation"}:      subnetnetworksecuritygroupassociation.Setup,
		schema.GroupVersionKind{Group: "subnet.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "SubnetRouteTableAssociation"}:                subnetroutetableassociation.Setup,
		schema.GroupVersionKind{Group: "subnet.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "SubnetServiceEndpointStoragePolicy"}:         subnetserviceendpointstoragepolicy.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualDesktopApplication"}:                 virtualdesktopapplication.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualDesktopHostPool"}:                    virtualdesktophostpool.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualDesktopWorkspace"}:                   virtualdesktopworkspace.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHub"}:                                virtualhub.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHubBgpConnection"}:                   virtualhubbgpconnection.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHubConnection"}:                      virtualhubconnection.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHubIp"}:                              virtualhubip.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHubRouteTable"}:                      virtualhubroutetable.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualHubSecurityPartnerProvider"}:         virtualhubsecuritypartnerprovider.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualNetwork"}:                            virtualnetwork.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualNetworkGateway"}:                     virtualnetworkgateway.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualNetworkGatewayConnection"}:           virtualnetworkgatewayconnection.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualNetworkPeering"}:                     virtualnetworkpeering.Setup,
		schema.GroupVersionKind{Group: "virtual.azure.tf.crossplane.io", Version: "v1alpha1", Kind: "VirtualWan"}:                                virtualwan.Setup,
	}
	return m
}
