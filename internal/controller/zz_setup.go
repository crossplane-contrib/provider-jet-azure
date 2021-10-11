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

	"github.com/crossplane-contrib/terrajet/pkg/terraform"

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

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, int) error{
		config.Setup,
		cosmosdbaccount.Setup,
		cosmosdbcassandrakeyspace.Setup,
		cosmosdbcassandratable.Setup,
		cosmosdbgremlindatabase.Setup,
		cosmosdbgremlingraph.Setup,
		cosmosdbmongocollection.Setup,
		cosmosdbmongodatabase.Setup,
		cosmosdbnotebookworkspace.Setup,
		cosmosdbsqlcontainer.Setup,
		cosmosdbsqldatabase.Setup,
		cosmosdbsqlfunction.Setup,
		cosmosdbsqlstoredprocedure.Setup,
		cosmosdbsqltrigger.Setup,
		cosmosdbtable.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		postgresqlactivedirectoryadministrator.Setup,
		postgresqlconfiguration.Setup,
		postgresqldatabase.Setup,
		postgresqlfirewallrule.Setup,
		postgresqlflexibleserver.Setup,
		postgresqlflexibleserverconfiguration.Setup,
		postgresqlflexibleserverdatabase.Setup,
		postgresqlflexibleserverfirewallrule.Setup,
		postgresqlserver.Setup,
		postgresqlserverkey.Setup,
		postgresqlvirtualnetworkrule.Setup,
		resourcegroup.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcegrouptemplatedeployment.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		virtualdesktopapplication.Setup,
		virtualdesktophostpool.Setup,
		virtualdesktopworkspace.Setup,
		virtualhub.Setup,
		virtualhubbgpconnection.Setup,
		virtualhubconnection.Setup,
		virtualhubip.Setup,
		virtualhubroutetable.Setup,
		virtualhubsecuritypartnerprovider.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, concurrency); err != nil {
			return err
		}
	}
	return nil
}
