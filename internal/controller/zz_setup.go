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
	cache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/cache"
	enterprisecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/enterprisecluster"
	enterprisedatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/enterprisedatabase"
	firewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/firewallrule"
	linkedserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/linkedserver"
	kubernetescluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetesclusternodepool"
	activedirectoryadministrator "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configuration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/configuration"
	database "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	server "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/server"
	virtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	consumergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/consumergroup"
	dps "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/dps"
	dpscertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/dpscertificate"
	dpssharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/dpssharedaccesspolicy"
	endpointeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/endpointeventhub"
	endpointservicebusqueue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/endpointservicebusqueue"
	endpointservicebustopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/endpointservicebustopic"
	endpointstoragecontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/endpointstoragecontainer"
	enrichment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/enrichment"
	fallbackroute "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/fallbackroute"
	iothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothub"
	route "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/route"
	sharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/sharedaccesspolicy"
	account "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/account"
	cassandrakeyspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/cassandrakeyspace"
	cassandratable "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/cassandratable"
	gremlindatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/gremlindatabase"
	gremlingraph "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/gremlingraph"
	mongocollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/mongocollection"
	mongodatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/mongodatabase"
	notebookworkspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/notebookworkspace"
	sqlcontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/sqlcontainer"
	sqldatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/sqldatabase"
	sqlfunction "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/sqlfunction"
	sqlstoredprocedure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/sqlstoredprocedure"
	sqltrigger "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/sqltrigger"
	table "github.com/crossplane-contrib/provider-jet-azure/internal/controller/documentdb/table"
	metricalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/metricalert"
	workspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/loganalytics/workspace"
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
		cache.Setup,
		enterprisecluster.Setup,
		enterprisedatabase.Setup,
		firewallrule.Setup,
		linkedserver.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		activedirectoryadministrator.Setup,
		configuration.Setup,
		database.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
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
		metricalert.Setup,
		workspace.Setup,
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
		providerconfig.Setup,
		resourcegrouptemplatedeployment.Setup,
		serversql.Setup,
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
