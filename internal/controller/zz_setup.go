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
	virtualdesktopapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopapplication"
	virtualdesktophostpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktophostpool"
	virtualdesktopworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopworkspace"
	virtualdesktopworkspaceapplicationgroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopworkspaceapplicationgroupassociation"
	virtualhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhub"
	virtualhubbgpconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubbgpconnection"
	virtualhubconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubconnection"
	virtualhubip "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubip"
	virtualhubroutetable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubroutetable"
	virtualhubsecuritypartnerprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubsecuritypartnerprovider"
	virtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachine"
	virtualmachineconfigurationpolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachineconfigurationpolicyassignment"
	virtualmachinedatadiskattachment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachineextension"
	virtualmachinescaleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachinescaleset"
	virtualmachinescalesetextension "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualmachinescalesetextension"
	virtualnetwork "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetwork"
	virtualnetworkdnsservers "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkdnsservers"
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
		virtualdesktopapplication.Setup,
		virtualdesktophostpool.Setup,
		virtualdesktopworkspace.Setup,
		virtualdesktopworkspaceapplicationgroupassociation.Setup,
		virtualhub.Setup,
		virtualhubbgpconnection.Setup,
		virtualhubconnection.Setup,
		virtualhubip.Setup,
		virtualhubroutetable.Setup,
		virtualhubsecuritypartnerprovider.Setup,
		virtualmachine.Setup,
		virtualmachineconfigurationpolicyassignment.Setup,
		virtualmachinedatadiskattachment.Setup,
		virtualmachineextension.Setup,
		virtualmachinescaleset.Setup,
		virtualmachinescalesetextension.Setup,
		virtualnetwork.Setup,
		virtualnetworkdnsservers.Setup,
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
