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

	accountcosmosdb "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/account"
	cassandrakeyspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/mongodatabase"
	notebookworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/notebookworkspace"
	sqlcontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/sqldatabase"
	sqlfunction "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/sqlfunction"
	sqlstoredprocedure "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cosmosdb/table"
	cluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/cluster"
	clusternodepool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/clusternodepool"
	loadbalancer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/loadbalancer"
	activedirectoryadministrator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/activedirectoryadministrator"
	configuration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/configuration"
	database "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/database"
	firewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/firewallrule"
	flexibleserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/flexibleserver"
	flexibleserverconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/flexibleserverfirewallrule"
	serverpostgresql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/server"
	virtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/virtualnetworkrule"
	providerconfig "github.com/crossplane-contrib/provider-tf-azure/internal/controller/providerconfig"
	grouppolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/grouppolicyassignment"
	grouptemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/grouptemplatedeployment"
	resourcegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegroup"
	server "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/server"
	account "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/account"
	blob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/blob"
	container "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/container"
	natgatewayassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/natgatewayassociation"
	networksecuritygroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/networksecuritygroupassociation"
	routetableassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/routetableassociation"
	serviceendpointstoragepolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/serviceendpointstoragepolicy"
	desktopapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/desktopapplication"
	desktophostpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/desktophostpool"
	network "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/network"
	networkgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/networkgateway"
	networkgatewayconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/networkgatewayconnection"
	networkpeering "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/networkpeering"
	subnet "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/subnet"
	wan "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/wan"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		account.Setup,
		accountcosmosdb.Setup,
		activedirectoryadministrator.Setup,
		blob.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		cluster.Setup,
		clusternodepool.Setup,
		configuration.Setup,
		container.Setup,
		database.Setup,
		desktopapplication.Setup,
		desktophostpool.Setup,
		firewallrule.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		grouppolicyassignment.Setup,
		grouptemplatedeployment.Setup,
		loadbalancer.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		natgatewayassociation.Setup,
		network.Setup,
		networkgateway.Setup,
		networkgatewayconnection.Setup,
		networkpeering.Setup,
		networksecuritygroupassociation.Setup,
		notebookworkspace.Setup,
		providerconfig.Setup,
		resourcegroup.Setup,
		routetableassociation.Setup,
		server.Setup,
		serverpostgresql.Setup,
		serviceendpointstoragepolicy.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqlfunction.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		subnet.Setup,
		table.Setup,
		virtualnetworkrule.Setup,
		wan.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
