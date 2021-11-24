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
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	resourcegrouppolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azure/resourcegroup"
	kubernetescluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetesclusternodepool"
	account "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/account"
	cassandrakeyspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/mongodatabase"
	notebookworkspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/notebookworkspace"
	sqlcontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/sqldatabase"
	sqlfunction "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/sqlfunction"
	sqlstoredprocedure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/table"
	consumergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/consumergroup"
	dps "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/dps"
	dpscertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/dpscertificate"
	dpssharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/dpssharedaccesspolicy"
	endpointeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/endpointeventhub"
	endpointservicebusqueue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/endpointservicebusqueue"
	endpointservicebustopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/endpointservicebustopic"
	endpointstoragecontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/endpointstoragecontainer"
	enrichment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/enrichment"
	fallbackroute "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/fallbackroute"
	iothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/iothub"
	route "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/route"
	sharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/sharedaccesspolicy"
	workspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/loganalytics/workspace"
	server "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/server"
	loadbalancer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/loadbalancer"
	subnet "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	virtualnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualwan"
	activedirectoryadministrator "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/activedirectoryadministrator"
	configuration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/configuration"
	database "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/database"
	firewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/firewallrule"
	flexibleserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserver"
	flexibleserverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverfirewallrule"
	serverpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/server"
	virtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/virtualnetworkrule"
	providerconfig "github.com/crossplane-contrib/provider-jet-azure/internal/controller/providerconfig"
	resourcegrouptemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	serversql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/server"
	accountstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/account"
	blob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/blob"
	container "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/container"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		resourcegrouppolicyassignment.Setup,
		resourcegroup.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		account.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		notebookworkspace.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqlfunction.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		consumergroup.Setup,
		dps.Setup,
		dpscertificate.Setup,
		dpssharedaccesspolicy.Setup,
		endpointeventhub.Setup,
		endpointservicebusqueue.Setup,
		endpointservicebustopic.Setup,
		endpointstoragecontainer.Setup,
		enrichment.Setup,
		fallbackroute.Setup,
		iothub.Setup,
		route.Setup,
		sharedaccesspolicy.Setup,
		workspace.Setup,
		server.Setup,
		loadbalancer.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		activedirectoryadministrator.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		serverpostgresql.Setup,
		virtualnetworkrule.Setup,
		providerconfig.Setup,
		resourcegrouptemplatedeployment.Setup,
		serversql.Setup,
		accountstorage.Setup,
		blob.Setup,
		container.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
