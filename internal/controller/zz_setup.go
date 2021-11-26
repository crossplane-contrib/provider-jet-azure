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

	directorydomainservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/active/directorydomainservice"
	directorydomainservicereplicaset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/active/directorydomainservicereplicaset"
	threatprotection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/advanced/threatprotection"
	servicesserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/analysis/servicesserver"
	management "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/management"
	managementapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapi"
	managementapidiagnostic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapidiagnostic"
	managementapioperation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapioperation"
	managementapioperationpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapioperationpolicy"
	managementapioperationtag "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapioperationtag"
	managementapipolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapipolicy"
	managementapirelease "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapirelease"
	managementapischema "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapischema"
	managementapiversionset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementapiversionset"
	managementauthorizationserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementauthorizationserver"
	managementbackend "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementbackend"
	managementcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementcertificate"
	managementcustomdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementcustomdomain"
	managementdiagnostic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementdiagnostic"
	managementemailtemplate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementemailtemplate"
	managementgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementgateway"
	managementgatewayapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementgatewayapi"
	managementidentityprovideraad "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityprovideraad"
	managementidentityprovideraadb2c "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityprovideraadb2c"
	managementidentityproviderfacebook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityproviderfacebook"
	managementidentityprovidergoogle "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityprovidergoogle"
	managementidentityprovidermicrosoft "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityprovidermicrosoft"
	managementidentityprovidertwitter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementidentityprovidertwitter"
	managementlogger "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementlogger"
	managementnamedvalue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementnamedvalue"
	managementnotificationrecipientemail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementnotificationrecipientemail"
	managementopenidconnectprovider "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementopenidconnectprovider"
	managementpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementpolicy"
	managementproduct "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementproduct"
	managementproductapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementproductapi"
	managementproductpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementproductpolicy"
	managementproperty "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementproperty"
	managementrediscache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementrediscache"
	managementsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementsubscription"
	managementtag "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementtag"
	managementuser "github.com/crossplane-contrib/provider-jet-azure/internal/controller/api/managementuser"
	configuration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/configuration"
	configurationkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/configurationkey"
	service "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/service"
	serviceactiveslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceactiveslot"
	servicecertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicecertificate"
	servicecertificatebinding "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicecertificatebinding"
	servicecertificateorder "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicecertificateorder"
	servicecustomhostnamebinding "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicecustomhostnamebinding"
	serviceenvironment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceenvironment"
	serviceenvironmentv3 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceenvironmentv3"
	servicehybridconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicehybridconnection"
	servicemanagedcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicemanagedcertificate"
	serviceplan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceplan"
	serviceslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceslot"
	serviceslotvirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/serviceslotvirtualnetworkswiftconnection"
	servicesourcecontroltoken "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicesourcecontroltoken"
	servicevirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/servicevirtualnetworkswiftconnection"
	gateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/gateway"
	insights "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/insights"
	insightsanalyticsitem "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/insightsanalyticsitem"
	insightsapikey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/insightsapikey"
	insightssmartdetectionrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/insightssmartdetectionrule"
	insightswebtest "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/insightswebtest"
	securitygroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/application/securitygroup"
	provider "github.com/crossplane-contrib/provider-jet-azure/internal/controller/attestation/provider"
	resourcegrouppolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	account "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/account"
	certificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/certificate"
	connection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/connection"
	connectioncertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/connectioncertificate"
	connectionclassiccertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/connectionclassiccertificate"
	connectionserviceprincipal "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/connectionserviceprincipal"
	credential "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/credential"
	dscnodeconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/dscnodeconfiguration"
	jobschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/jobschedule"
	module "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/module"
	runbook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/runbook"
	schedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/schedule"
	variablebool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/variablebool"
	variabledatetime "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/variabledatetime"
	variableint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/variableint"
	variablestring "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/variablestring"
	set "github.com/crossplane-contrib/provider-jet-azure/internal/controller/availability/set"
	resourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azure/resourcegroup"
	dashboard "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/dashboard"
	eventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/eventhub"
	firewall "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/firewall"
	frontdoor "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/frontdoor"
	healthbot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/healthbot"
	image "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/image"
	snapshot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/snapshot"
	subscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurerm/subscription"
	containerstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/backup/containerstorageaccount"
	policyfileshare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/backup/policyfileshare"
	policyvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/backup/policyvm"
	protectedfileshare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/backup/protectedfileshare"
	protectedvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/backup/protectedvm"
	host "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bastion/host"
	accountbatch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/account"
	application "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/application"
	certificatebatch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/certificate"
	job "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/job"
	pool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/pool"
	assignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/blueprint/assignment"
	channelalexa "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelalexa"
	channeldirectline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channeldirectline"
	channeldirectlinespeech "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channeldirectlinespeech"
	channelemail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelemail"
	channelfacebook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelfacebook"
	channelline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelline"
	channelmsteams "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelmsteams"
	channelslack "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelslack"
	channelsms "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelsms"
	channelsregistration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelsregistration"
	channelwebchat "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/channelwebchat"
	connectionbot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/connection"
	webapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/bot/webapp"
	endpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/endpoint"
	endpointcustomdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/endpointcustomdomain"
	profile "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/profile"
	accountcognitive "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cognitive/account"
	accountcustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cognitive/accountcustomermanagedkey"
	servicecommunication "github.com/crossplane-contrib/provider-jet-azure/internal/controller/communication/service"
	budgetresourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/consumption/budgetresourcegroup"
	budgetsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/consumption/budgetsubscription"
	registry "github.com/crossplane-contrib/provider-jet-azure/internal/controller/container/registry"
	registryscopemap "github.com/crossplane-contrib/provider-jet-azure/internal/controller/container/registryscopemap"
	registrytoken "github.com/crossplane-contrib/provider-jet-azure/internal/controller/container/registrytoken"
	registrywebhook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/container/registrywebhook"
	kubernetescluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerservice/kubernetesclusternodepool"
	accountcosmosdb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cosmosdb/account"
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
	managementexportresourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cost/managementexportresourcegroup"
	providercustom "github.com/crossplane-contrib/provider-jet-azure/internal/controller/custom/provider"
	factory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factory"
	factorycustomdataset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorycustomdataset"
	factorydataflow "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydataflow"
	factorydatasetazureblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetazureblob"
	factorydatasetbinary "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetbinary"
	factorydatasetcosmosdbsqlapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetcosmosdbsqlapi"
	factorydatasetdelimitedtext "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetdelimitedtext"
	factorydatasethttp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasethttp"
	factorydatasetjson "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetjson"
	factorydatasetmysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetmysql"
	factorydatasetparquet "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetparquet"
	factorydatasetpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetpostgresql"
	factorydatasetsnowflake "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetsnowflake"
	factorydatasetsqlservertable "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorydatasetsqlservertable"
	factoryintegrationruntimeazure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factoryintegrationruntimeazure"
	factoryintegrationruntimeazuressis "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factoryintegrationruntimeazuressis"
	factoryintegrationruntimemanaged "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factoryintegrationruntimemanaged"
	factoryintegrationruntimeselfhosted "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factoryintegrationruntimeselfhosted"
	factorylinkedcustomservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedcustomservice"
	factorylinkedserviceazureblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazureblobstorage"
	factorylinkedserviceazuredatabricks "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazuredatabricks"
	factorylinkedserviceazurefilestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazurefilestorage"
	factorylinkedserviceazurefunction "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazurefunction"
	factorylinkedserviceazuresearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazuresearch"
	factorylinkedserviceazuresqldatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazuresqldatabase"
	factorylinkedserviceazuretablestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceazuretablestorage"
	factorylinkedservicecosmosdb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicecosmosdb"
	factorylinkedservicedatalakestoragegen2 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicedatalakestoragegen2"
	factorylinkedservicekeyvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicekeyvault"
	factorylinkedservicekusto "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicekusto"
	factorylinkedservicemysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicemysql"
	factorylinkedserviceodata "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceodata"
	factorylinkedservicepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicepostgresql"
	factorylinkedservicesftp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicesftp"
	factorylinkedservicesnowflake "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicesnowflake"
	factorylinkedservicesqlserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicesqlserver"
	factorylinkedservicesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicesynapse"
	factorylinkedserviceweb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedserviceweb"
	factorymanagedprivateendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorymanagedprivateendpoint"
	factorypipeline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorypipeline"
	factorytriggerblobevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorytriggerblobevent"
	factorytriggercustomevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorytriggercustomevent"
	factorytriggerschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorytriggerschedule"
	lakeanalyticsfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/lakeanalyticsfirewallrule"
	lakestore "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/lakestore"
	lakestorefile "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/lakestorefile"
	lakestorefirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/lakestorefirewallrule"
	lakestorevirtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/lakestorevirtualnetworkrule"
	protectionbackupinstanceblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackupinstanceblobstorage"
	protectionbackupinstancedisk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackupinstancedisk"
	protectionbackupinstancepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackupinstancepostgresql"
	protectionbackuppolicyblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackuppolicyblobstorage"
	protectionbackuppolicydisk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackuppolicydisk"
	protectionbackuppolicypostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackuppolicypostgresql"
	protectionbackupvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/protectionbackupvault"
	share "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/share"
	shareaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/shareaccount"
	sharedatasetblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/sharedatasetblobstorage"
	sharedatasetdatalakegen1 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/sharedatasetdatalakegen1"
	sharedatasetdatalakegen2 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/sharedatasetdatalakegen2"
	sharedatasetkustocluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/sharedatasetkustocluster"
	sharedatasetkustodatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/sharedatasetkustodatabase"
	migrationproject "github.com/crossplane-contrib/provider-jet-azure/internal/controller/database/migrationproject"
	migrationservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/database/migrationservice"
	edgedevice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databox/edgedevice"
	edgeorder "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databox/edgeorder"
	workspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databricks/workspace"
	workspacecustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databricks/workspacecustomermanagedkey"
	hardwaresecuritymodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dedicated/hardwaresecuritymodule"
	hostdedicated "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dedicated/host"
	testglobalvmshutdownschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testglobalvmshutdownschedule"
	testlab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testlab"
	testlinuxvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testlinuxvirtualmachine"
	testpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testpolicy"
	testschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testschedule"
	testvirtualnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testvirtualnetwork"
	testwindowsvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dev/testwindowsvirtualmachine"
	controller "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devspace/controller"
	twinsendpointeventgrid "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digital/twinsendpointeventgrid"
	twinsendpointeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digital/twinsendpointeventhub"
	twinsendpointservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digital/twinsendpointservicebus"
	twinsinstance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digital/twinsinstance"
	access "github.com/crossplane-contrib/provider-jet-azure/internal/controller/disk/access"
	encryptionset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/disk/encryptionset"
	aaaarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/aaaarecord"
	arecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/arecord"
	caarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/caarecord"
	cnamerecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/cnamerecord"
	mxrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/mxrecord"
	nsrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/nsrecord"
	ptrrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/ptrrecord"
	srvrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/srvrecord"
	txtrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/txtrecord"
	zone "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dns/zone"
	domain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/domain"
	domaintopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/domaintopic"
	eventsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/eventsubscription"
	systemtopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/systemtopic"
	systemtopiceventsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/systemtopiceventsubscription"
	topic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/topic"
	authorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/authorizationrule"
	cluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/cluster"
	consumergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/consumergroup"
	namespace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespace"
	namespaceauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespaceauthorizationrule"
	namespacecustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespacecustomermanagedkey"
	namespacedisasterrecoveryconfig "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespacedisasterrecoveryconfig"
	routecircuit "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routecircuit"
	routecircuitauthorization "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routecircuitauthorization"
	routecircuitconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routecircuitconnection"
	routecircuitpeering "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routecircuitpeering"
	routeconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routeconnection"
	routegateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routegateway"
	routeport "github.com/crossplane-contrib/provider-jet-azure/internal/controller/express/routeport"
	applicationrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/firewall/applicationrulecollection"
	natrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/firewall/natrulecollection"
	networkrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/firewall/networkrulecollection"
	policy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/firewall/policy"
	policyrulecollectiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/firewall/policyrulecollectiongroup"
	customhttpsconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/frontdoor/customhttpsconfiguration"
	firewallpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/frontdoor/firewallpolicy"
	rulesengine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/frontdoor/rulesengine"
	app "github.com/crossplane-contrib/provider-jet-azure/internal/controller/function/app"
	appslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/function/appslot"
	hadoopcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/hadoopcluster"
	hbasecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/hbasecluster"
	interactivequerycluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/interactivequerycluster"
	kafkacluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/kafkacluster"
	mlservicescluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/mlservicescluster"
	rservercluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/rservercluster"
	sparkcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/sparkcluster"
	stormcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/stormcluster"
	servicehealthcare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/healthcare/service"
	cache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hpc/cache"
	cacheaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hpc/cacheaccesspolicy"
	cacheblobnfstarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hpc/cacheblobnfstarget"
	cacheblobtarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hpc/cacheblobtarget"
	cachenfstarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hpc/cachenfstarget"
	serviceenvironmentintegration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/integration/serviceenvironment"
	securitydevicegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/securitydevicegroup"
	securitysolution "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/securitysolution"
	timeseriesinsightsaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightsaccesspolicy"
	timeseriesinsightseventsourceiothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightseventsourceiothub"
	timeseriesinsightsgen2environment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightsgen2environment"
	timeseriesinsightsreferencedataset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightsreferencedataset"
	timeseriesinsightsstandardenvironment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightsstandardenvironment"
	applicationiotcentral "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iotcentral/application"
	consumergroupiothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iothub/consumergroup"
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
	vault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vault"
	vaultaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultaccesspolicy"
	vaultcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultcertificate"
	vaultcertificateissuer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultcertificateissuer"
	vaultkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultkey"
	vaultmanagedhardwaresecuritymodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultmanagedhardwaresecuritymodule"
	vaultmanagedstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultmanagedstorageaccount"
	vaultmanagedstorageaccountsastokendefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultmanagedstorageaccountsastokendefinition"
	vaultsecret "github.com/crossplane-contrib/provider-jet-azure/internal/controller/key/vaultsecret"
	attacheddatabaseconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/attacheddatabaseconfiguration"
	clusterkusto "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/cluster"
	clustercustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/clustercustomermanagedkey"
	clusterprincipalassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/clusterprincipalassignment"
	database "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/database"
	databaseprincipal "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/databaseprincipal"
	databaseprincipalassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/iothubdataconnection"
	backendaddresspool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/backendaddresspool"
	backendaddresspooladdress "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/backendaddresspooladdress"
	natpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/natpool"
	natrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/natrule"
	outboundrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/outboundrule"
	probe "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/probe"
	rule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lb/rule"
	assignmentlighthouse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lighthouse/assignment"
	definition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/lighthouse/definition"
	virtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/linux/virtualmachine"
	virtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/linux/virtualmachinescaleset"
	networkgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/local/networkgateway"
	analyticscluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticscluster"
	analyticsclustercustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticsclustercustomermanagedkey"
	analyticsdataexportrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticsdataexportrule"
	analyticsdatasourcewindowsevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticsdatasourcewindowsevent"
	analyticsdatasourcewindowsperformancecounter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticsdatasourcewindowsperformancecounter"
	analyticslinkedservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticslinkedservice"
	analyticslinkedstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticslinkedstorageaccount"
	analyticssavedsearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticssavedsearch"
	analyticssolution "github.com/crossplane-contrib/provider-jet-azure/internal/controller/log/analyticssolution"
	workspaceloganalytics "github.com/crossplane-contrib/provider-jet-azure/internal/controller/loganalytics/workspace"
	appactioncustom "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appactioncustom"
	appactionhttp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appactionhttp"
	appintegrationaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccount"
	appintegrationaccountagreement "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountagreement"
	appintegrationaccountassembly "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountassembly"
	appintegrationaccountbatchconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountmap "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountmap"
	appintegrationaccountpartner "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appintegrationaccountsession"
	apptriggercustom "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggercustom"
	apptriggerhttprequest "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggerrecurrence"
	appworkflow "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appworkflow"
	learningcomputecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machine/learningcomputecluster"
	learningcomputeinstance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machine/learningcomputeinstance"
	learningsynapsespark "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machine/learningsynapsespark"
	learningworkspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machine/learningworkspace"
	assignmentdedicatedhost "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/assignmentdedicatedhost"
	assignmentvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/assignmentvirtualmachine"
	assignmentvirtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/assignmentvirtualmachinescaleset"
	configurationmaintenance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/configuration"
	applicationmanaged "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managed/application"
	applicationdefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managed/applicationdefinition"
	disk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managed/disk"
	grouppolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/grouppolicyassignment"
	groupsubscriptionassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/groupsubscriptionassociation"
	grouptemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/grouptemplatedeployment"
	lock "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/lock"
	managementgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/managementgroup"
	accountmaps "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maps/account"
	configurationmariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mariadb/configuration"
	databasemariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mariadb/database"
	firewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mariadb/firewallrule"
	server "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mariadb/server"
	virtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mariadb/virtualnetworkrule"
	agreement "github.com/crossplane-contrib/provider-jet-azure/internal/controller/marketplace/agreement"
	asset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/asset"
	assetfilter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/assetfilter"
	contentkeypolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/contentkeypolicy"
	jobmedia "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/job"
	liveevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/liveevent"
	liveeventoutput "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/liveeventoutput"
	servicesaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/servicesaccount"
	streamingendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/streamingendpoint"
	streaminglocator "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/streaminglocator"
	streamingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/streamingpolicy"
	transform "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/transform"
	aaddiagnosticsetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/aaddiagnosticsetting"
	actiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/actiongroup"
	actionruleactiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/actionruleactiongroup"
	actionrulesuppression "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/actionrulesuppression"
	activitylogalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/activitylogalert"
	autoscalesetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/autoscalesetting"
	diagnosticsetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/diagnosticsetting"
	metricalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/metricalert"
	scheduledqueryrulesalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/scheduledqueryrulesalert"
	scheduledqueryruleslog "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/scheduledqueryruleslog"
	smartdetectoralertrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/monitor/smartdetectoralertrule"
	databasemssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/database"
	databaseextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/databaseextendedauditingpolicy"
	databasevulnerabilityassessmentrulebaseline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/databasevulnerabilityassessmentrulebaseline"
	elasticpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/elasticpool"
	failovergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/failovergroup"
	firewallrulemssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/firewallrule"
	jobagent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/jobagent"
	jobcredential "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/jobcredential"
	servermssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/server"
	serversecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/serversecurityalertpolicy"
	servertransparentdataencryption "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/servertransparentdataencryption"
	servervulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/servervulnerabilityassessment"
	virtualmachinemssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/virtualmachine"
	virtualnetworkrulemssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mssql/virtualnetworkrule"
	activedirectoryadministrator "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/activedirectoryadministrator"
	configurationmysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/configuration"
	databasemysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/database"
	firewallrulemysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/firewallrule"
	servermysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/server"
	serverkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/serverkey"
	virtualnetworkrulemysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/virtualnetworkrule"
	gatewaynat "github.com/crossplane-contrib/provider-jet-azure/internal/controller/nat/gateway"
	gatewaypublicipassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/nat/gatewaypublicipassociation"
	gatewaypublicipprefixassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/nat/gatewaypublicipprefixassociation"
	accountnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/account"
	poolnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/pool"
	snapshotnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/snapshot"
	volume "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/volume"
	connectionmonitor "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/connectionmonitor"
	ddosprotectionplan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/ddosprotectionplan"
	interfaceapplicationsecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfaceapplicationsecuritygroupassociation"
	interfacebackendaddresspoolassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacebackendaddresspoolassociation"
	interfacenatruleassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacenatruleassociation"
	interfacesecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacesecuritygroupassociation"
	ipgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/ipgroup"
	loadbalancer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/loadbalancer"
	networkinterface "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/networkinterface"
	packetcapture "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/packetcapture"
	profilenetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/profile"
	securitygroupnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/securitygroup"
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
	watcher "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/watcher"
	watcherflowlog "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/watcherflowlog"
	hub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notification/hub"
	hubauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notification/hubauthorizationrule"
	hubnamespace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notification/hubnamespace"
	virtualmachinescalesetorchestrated "github.com/crossplane-contrib/provider-jet-azure/internal/controller/orchestrated/virtualmachinescaleset"
	capture "github.com/crossplane-contrib/provider-jet-azure/internal/controller/packet/capture"
	tositevpngateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/point/tositevpngateway"
	assignmentpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policy/assignment"
	definitionpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policy/definition"
	remediation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policy/remediation"
	setdefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policy/setdefinition"
	virtualmachineconfigurationassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policy/virtualmachineconfigurationassignment"
	tenantconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/portal/tenantconfiguration"
	activedirectoryadministratorpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/activedirectoryadministrator"
	configurationpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/configuration"
	databasepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/database"
	firewallrulepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/firewallrule"
	flexibleserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserver"
	flexibleserverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/flexibleserverfirewallrule"
	serverpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/server"
	virtualnetworkrulepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/postgresql/virtualnetworkrule"
	embedded "github.com/crossplane-contrib/provider-jet-azure/internal/controller/powerbi/embedded"
	dnsaaaarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnsaaaarecord"
	dnsarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnsarecord"
	dnscnamerecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnscnamerecord"
	dnsmxrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnsmxrecord"
	dnsptrrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnsptrrecord"
	dnssrvrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnssrvrecord"
	dnstxtrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnstxtrecord"
	dnszone "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnszone"
	dnszonevirtualnetworklink "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/dnszonevirtualnetworklink"
	endpointprivate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/endpoint"
	linkservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/private/linkservice"
	providerconfig "github.com/crossplane-contrib/provider-jet-azure/internal/controller/providerconfig"
	placementgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/proximity/placementgroup"
	ip "github.com/crossplane-contrib/provider-jet-azure/internal/controller/public/ip"
	ipprefix "github.com/crossplane-contrib/provider-jet-azure/internal/controller/public/ipprefix"
	accountpurview "github.com/crossplane-contrib/provider-jet-azure/internal/controller/purview/account"
	servicesvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recovery/servicesvault"
	cacheredis "github.com/crossplane-contrib/provider-jet-azure/internal/controller/redis/cache"
	enterprisecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/redis/enterprisecluster"
	enterprisedatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/redis/enterprisedatabase"
	firewallruleredis "github.com/crossplane-contrib/provider-jet-azure/internal/controller/redis/firewallrule"
	linkedserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/redis/linkedserver"
	hybridconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/hybridconnectionauthorizationrule"
	namespacerelay "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/namespace"
	namespaceauthorizationrulerelay "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/namespaceauthorizationrule"
	policyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resource/policyassignment"
	providerregistration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resource/providerregistration"
	resourcegrouptemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	assignmentrole "github.com/crossplane-contrib/provider-jet-azure/internal/controller/role/assignment"
	definitionrole "github.com/crossplane-contrib/provider-jet-azure/internal/controller/role/definition"
	servicesearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/search/service"
	centerassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerassessment"
	centerassessmentmetadata "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerassessmentmetadata"
	centerassessmentpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerassessmentpolicy"
	centerautoprovisioning "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerautoprovisioning"
	centercontact "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centercontact"
	centerservervulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerservervulnerabilityassessment"
	centersetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centersetting"
	centersubscriptionpricing "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centersubscriptionpricing"
	centerworkspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/centerworkspace"
	alertrulefusion "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/alertrulefusion"
	alertrulemachinelearningbehavioranalytics "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/alertrulemachinelearningbehavioranalytics"
	alertrulemssecurityincident "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/alertrulemssecurityincident"
	alertrulescheduled "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/alertrulescheduled"
	dataconnectorawscloudtrail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectorawscloudtrail"
	dataconnectorazureactivedirectory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectorazureactivedirectory"
	dataconnectorazureadvancedthreatprotection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectorazureadvancedthreatprotection"
	dataconnectorazuresecuritycenter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectorazuresecuritycenter"
	dataconnectormicrosoftcloudappsecurity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectormicrosoftcloudappsecurity"
	dataconnectoroffice365 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectoroffice365"
	dataconnectorthreatintelligence "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/dataconnectorthreatintelligence"
	fabriccluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/service/fabriccluster"
	fabricmeshapplication "github.com/crossplane-contrib/provider-jet-azure/internal/controller/service/fabricmeshapplication"
	fabricmeshlocalnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/service/fabricmeshlocalnetwork"
	fabricmeshsecret "github.com/crossplane-contrib/provider-jet-azure/internal/controller/service/fabricmeshsecret"
	fabricmeshsecretvalue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/service/fabricmeshsecretvalue"
	namespaceservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/namespace"
	namespaceauthorizationruleservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfigservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/namespacedisasterrecoveryconfig"
	namespacenetworkruleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/namespacenetworkruleset"
	queue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/queue"
	queueauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/queueauthorizationrule"
	subscriptionservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/subscription"
	subscriptionrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/subscriptionrule"
	topicservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/topic"
	topicauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicebus/topicauthorizationrule"
	imageshared "github.com/crossplane-contrib/provider-jet-azure/internal/controller/shared/image"
	imagegallery "github.com/crossplane-contrib/provider-jet-azure/internal/controller/shared/imagegallery"
	imageversion "github.com/crossplane-contrib/provider-jet-azure/internal/controller/shared/imageversion"
	servicesignalr "github.com/crossplane-contrib/provider-jet-azure/internal/controller/signalr/service"
	servicenetworkacl "github.com/crossplane-contrib/provider-jet-azure/internal/controller/signalr/servicenetworkacl"
	recoveryfabric "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoveryfabric"
	recoverynetworkmapping "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoverynetworkmapping"
	recoveryprotectioncontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoveryprotectioncontainer"
	recoveryprotectioncontainermapping "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoveryprotectioncontainermapping"
	recoveryreplicatedvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoveryreplicatedvm"
	recoveryreplicationpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/site/recoveryreplicationpolicy"
	anchorsaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spatial/anchorsaccount"
	cloudactivedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudactivedeployment"
	cloudapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudapp"
	cloudappcosmosdbassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudappcosmosdbassociation"
	cloudappmysqlassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudappmysqlassociation"
	cloudappredisassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudappredisassociation"
	cloudcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudcertificate"
	cloudcustomdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudcustomdomain"
	cloudjavadeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudjavadeployment"
	cloudservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/spring/cloudservice"
	activedirectoryadministratorsql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/activedirectoryadministrator"
	databasesql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/database"
	elasticpoolsql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/elasticpool"
	firewallrulesql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/firewallrule"
	manageddatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/manageddatabase"
	managedinstance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/managedinstance"
	serversql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/server"
	publickey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/ssh/publickey"
	hcicluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stack/hcicluster"
	site "github.com/crossplane-contrib/provider-jet-azure/internal/controller/static/site"
	accountstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/account"
	accountcustomermanagedkeystorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/accountcustomermanagedkey"
	accountnetworkrules "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/datalakegen2filesystem"
	datalakegen2path "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/datalakegen2path"
	encryptionscope "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/encryptionscope"
	managementpolicystorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/objectreplication"
	queuestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/queue"
	sharestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/share"
	sharedirectory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/sharedirectory"
	sync "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/sync"
	synccloudendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/synccloudendpoint"
	tablestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/table"
	tableentity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/tableentity"
	analyticsfunctionjavascriptudf "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsfunctionjavascriptudf"
	analyticsjob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsjob"
	analyticsoutputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputblob"
	analyticsoutputeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputeventhub"
	analyticsoutputmssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputmssql"
	analyticsoutputservicebusqueue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputservicebusqueue"
	analyticsoutputservicebustopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputservicebustopic"
	analyticsreferenceinputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsreferenceinputblob"
	analyticsstreaminputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsstreaminputblob"
	analyticsstreaminputeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsstreaminputeventhub"
	analyticsstreaminputiothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsstreaminputiothub"
	policyassignmentsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/subscription/policyassignment"
	templatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/subscription/templatedeployment"
	firewallrulesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/firewallrule"
	integrationruntimeazure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/integrationruntimeazure"
	integrationruntimeselfhosted "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/linkedservice"
	managedprivateendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/managedprivateendpoint"
	privatelinkhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/privatelinkhub"
	roleassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/roleassignment"
	sparkpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sparkpool"
	sqlpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolvulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolvulnerabilityassessment"
	workspacesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspace"
	workspaceextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspaceextendedauditingpolicy"
	workspacekey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacekey"
	workspacesecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacesecurityalertpolicy"
	workspacevulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacevulnerabilityassessment"
	deployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/template/deployment"
	templatedeploymenttenant "github.com/crossplane-contrib/provider-jet-azure/internal/controller/tenant/templatedeployment"
	managerendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/traffic/managerendpoint"
	managerprofile "github.com/crossplane-contrib/provider-jet-azure/internal/controller/traffic/managerprofile"
	assignedidentity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/user/assignedidentity"
	analyzer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/video/analyzer"
	analyzeredgemodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/video/analyzeredgemodule"
	clustervmware "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vmware/cluster"
	expressrouteauthorization "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vmware/expressrouteauthorization"
	privatecloud "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vmware/privatecloud"
	gatewayvpn "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vpn/gateway"
	gatewayconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vpn/gatewayconnection"
	serverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vpn/serverconfiguration"
	sitevpn "github.com/crossplane-contrib/provider-jet-azure/internal/controller/vpn/site"
	applicationfirewallpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/applicationfirewallpolicy"
	virtualmachinewindows "github.com/crossplane-contrib/provider-jet-azure/internal/controller/windows/virtualmachine"
	virtualmachinescalesetwindows "github.com/crossplane-contrib/provider-jet-azure/internal/controller/windows/virtualmachinescaleset"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		directorydomainservice.Setup,
		directorydomainservicereplicaset.Setup,
		threatprotection.Setup,
		servicesserver.Setup,
		management.Setup,
		managementapi.Setup,
		managementapidiagnostic.Setup,
		managementapioperation.Setup,
		managementapioperationpolicy.Setup,
		managementapioperationtag.Setup,
		managementapipolicy.Setup,
		managementapirelease.Setup,
		managementapischema.Setup,
		managementapiversionset.Setup,
		managementauthorizationserver.Setup,
		managementbackend.Setup,
		managementcertificate.Setup,
		managementcustomdomain.Setup,
		managementdiagnostic.Setup,
		managementemailtemplate.Setup,
		managementgateway.Setup,
		managementgatewayapi.Setup,
		managementidentityprovideraad.Setup,
		managementidentityprovideraadb2c.Setup,
		managementidentityproviderfacebook.Setup,
		managementidentityprovidergoogle.Setup,
		managementidentityprovidermicrosoft.Setup,
		managementidentityprovidertwitter.Setup,
		managementlogger.Setup,
		managementnamedvalue.Setup,
		managementnotificationrecipientemail.Setup,
		managementopenidconnectprovider.Setup,
		managementpolicy.Setup,
		managementproduct.Setup,
		managementproductapi.Setup,
		managementproductpolicy.Setup,
		managementproperty.Setup,
		managementrediscache.Setup,
		managementsubscription.Setup,
		managementtag.Setup,
		managementuser.Setup,
		configuration.Setup,
		configurationkey.Setup,
		service.Setup,
		serviceactiveslot.Setup,
		servicecertificate.Setup,
		servicecertificatebinding.Setup,
		servicecertificateorder.Setup,
		servicecustomhostnamebinding.Setup,
		serviceenvironment.Setup,
		serviceenvironmentv3.Setup,
		servicehybridconnection.Setup,
		servicemanagedcertificate.Setup,
		serviceplan.Setup,
		serviceslot.Setup,
		serviceslotvirtualnetworkswiftconnection.Setup,
		servicesourcecontroltoken.Setup,
		servicevirtualnetworkswiftconnection.Setup,
		gateway.Setup,
		insights.Setup,
		insightsanalyticsitem.Setup,
		insightsapikey.Setup,
		insightssmartdetectionrule.Setup,
		insightswebtest.Setup,
		securitygroup.Setup,
		provider.Setup,
		resourcegrouppolicyassignment.Setup,
		account.Setup,
		certificate.Setup,
		connection.Setup,
		connectioncertificate.Setup,
		connectionclassiccertificate.Setup,
		connectionserviceprincipal.Setup,
		credential.Setup,
		dscnodeconfiguration.Setup,
		jobschedule.Setup,
		module.Setup,
		runbook.Setup,
		schedule.Setup,
		variablebool.Setup,
		variabledatetime.Setup,
		variableint.Setup,
		variablestring.Setup,
		set.Setup,
		resourcegroup.Setup,
		dashboard.Setup,
		eventhub.Setup,
		firewall.Setup,
		frontdoor.Setup,
		healthbot.Setup,
		image.Setup,
		snapshot.Setup,
		subscription.Setup,
		containerstorageaccount.Setup,
		policyfileshare.Setup,
		policyvm.Setup,
		protectedfileshare.Setup,
		protectedvm.Setup,
		host.Setup,
		accountbatch.Setup,
		application.Setup,
		certificatebatch.Setup,
		job.Setup,
		pool.Setup,
		assignment.Setup,
		channelalexa.Setup,
		channeldirectline.Setup,
		channeldirectlinespeech.Setup,
		channelemail.Setup,
		channelfacebook.Setup,
		channelline.Setup,
		channelmsteams.Setup,
		channelslack.Setup,
		channelsms.Setup,
		channelsregistration.Setup,
		channelwebchat.Setup,
		connectionbot.Setup,
		webapp.Setup,
		endpoint.Setup,
		endpointcustomdomain.Setup,
		profile.Setup,
		accountcognitive.Setup,
		accountcustomermanagedkey.Setup,
		servicecommunication.Setup,
		budgetresourcegroup.Setup,
		budgetsubscription.Setup,
		registry.Setup,
		registryscopemap.Setup,
		registrytoken.Setup,
		registrywebhook.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		accountcosmosdb.Setup,
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
		managementexportresourcegroup.Setup,
		providercustom.Setup,
		factory.Setup,
		factorycustomdataset.Setup,
		factorydataflow.Setup,
		factorydatasetazureblob.Setup,
		factorydatasetbinary.Setup,
		factorydatasetcosmosdbsqlapi.Setup,
		factorydatasetdelimitedtext.Setup,
		factorydatasethttp.Setup,
		factorydatasetjson.Setup,
		factorydatasetmysql.Setup,
		factorydatasetparquet.Setup,
		factorydatasetpostgresql.Setup,
		factorydatasetsnowflake.Setup,
		factorydatasetsqlservertable.Setup,
		factoryintegrationruntimeazure.Setup,
		factoryintegrationruntimeazuressis.Setup,
		factoryintegrationruntimemanaged.Setup,
		factoryintegrationruntimeselfhosted.Setup,
		factorylinkedcustomservice.Setup,
		factorylinkedserviceazureblobstorage.Setup,
		factorylinkedserviceazuredatabricks.Setup,
		factorylinkedserviceazurefilestorage.Setup,
		factorylinkedserviceazurefunction.Setup,
		factorylinkedserviceazuresearch.Setup,
		factorylinkedserviceazuresqldatabase.Setup,
		factorylinkedserviceazuretablestorage.Setup,
		factorylinkedservicecosmosdb.Setup,
		factorylinkedservicedatalakestoragegen2.Setup,
		factorylinkedservicekeyvault.Setup,
		factorylinkedservicekusto.Setup,
		factorylinkedservicemysql.Setup,
		factorylinkedserviceodata.Setup,
		factorylinkedservicepostgresql.Setup,
		factorylinkedservicesftp.Setup,
		factorylinkedservicesnowflake.Setup,
		factorylinkedservicesqlserver.Setup,
		factorylinkedservicesynapse.Setup,
		factorylinkedserviceweb.Setup,
		factorymanagedprivateendpoint.Setup,
		factorypipeline.Setup,
		factorytriggerblobevent.Setup,
		factorytriggercustomevent.Setup,
		factorytriggerschedule.Setup,
		lakeanalyticsfirewallrule.Setup,
		lakestore.Setup,
		lakestorefile.Setup,
		lakestorefirewallrule.Setup,
		lakestorevirtualnetworkrule.Setup,
		protectionbackupinstanceblobstorage.Setup,
		protectionbackupinstancedisk.Setup,
		protectionbackupinstancepostgresql.Setup,
		protectionbackuppolicyblobstorage.Setup,
		protectionbackuppolicydisk.Setup,
		protectionbackuppolicypostgresql.Setup,
		protectionbackupvault.Setup,
		share.Setup,
		shareaccount.Setup,
		sharedatasetblobstorage.Setup,
		sharedatasetdatalakegen1.Setup,
		sharedatasetdatalakegen2.Setup,
		sharedatasetkustocluster.Setup,
		sharedatasetkustodatabase.Setup,
		migrationproject.Setup,
		migrationservice.Setup,
		edgedevice.Setup,
		edgeorder.Setup,
		workspace.Setup,
		workspacecustomermanagedkey.Setup,
		hardwaresecuritymodule.Setup,
		hostdedicated.Setup,
		testglobalvmshutdownschedule.Setup,
		testlab.Setup,
		testlinuxvirtualmachine.Setup,
		testpolicy.Setup,
		testschedule.Setup,
		testvirtualnetwork.Setup,
		testwindowsvirtualmachine.Setup,
		controller.Setup,
		twinsendpointeventgrid.Setup,
		twinsendpointeventhub.Setup,
		twinsendpointservicebus.Setup,
		twinsinstance.Setup,
		access.Setup,
		encryptionset.Setup,
		aaaarecord.Setup,
		arecord.Setup,
		caarecord.Setup,
		cnamerecord.Setup,
		mxrecord.Setup,
		nsrecord.Setup,
		ptrrecord.Setup,
		srvrecord.Setup,
		txtrecord.Setup,
		zone.Setup,
		domain.Setup,
		domaintopic.Setup,
		eventsubscription.Setup,
		systemtopic.Setup,
		systemtopiceventsubscription.Setup,
		topic.Setup,
		authorizationrule.Setup,
		cluster.Setup,
		consumergroup.Setup,
		namespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacecustomermanagedkey.Setup,
		namespacedisasterrecoveryconfig.Setup,
		routecircuit.Setup,
		routecircuitauthorization.Setup,
		routecircuitconnection.Setup,
		routecircuitpeering.Setup,
		routeconnection.Setup,
		routegateway.Setup,
		routeport.Setup,
		applicationrulecollection.Setup,
		natrulecollection.Setup,
		networkrulecollection.Setup,
		policy.Setup,
		policyrulecollectiongroup.Setup,
		customhttpsconfiguration.Setup,
		firewallpolicy.Setup,
		rulesengine.Setup,
		app.Setup,
		appslot.Setup,
		hadoopcluster.Setup,
		hbasecluster.Setup,
		interactivequerycluster.Setup,
		kafkacluster.Setup,
		mlservicescluster.Setup,
		rservercluster.Setup,
		sparkcluster.Setup,
		stormcluster.Setup,
		servicehealthcare.Setup,
		cache.Setup,
		cacheaccesspolicy.Setup,
		cacheblobnfstarget.Setup,
		cacheblobtarget.Setup,
		cachenfstarget.Setup,
		serviceenvironmentintegration.Setup,
		securitydevicegroup.Setup,
		securitysolution.Setup,
		timeseriesinsightsaccesspolicy.Setup,
		timeseriesinsightseventsourceiothub.Setup,
		timeseriesinsightsgen2environment.Setup,
		timeseriesinsightsreferencedataset.Setup,
		timeseriesinsightsstandardenvironment.Setup,
		applicationiotcentral.Setup,
		consumergroupiothub.Setup,
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
		vault.Setup,
		vaultaccesspolicy.Setup,
		vaultcertificate.Setup,
		vaultcertificateissuer.Setup,
		vaultkey.Setup,
		vaultmanagedhardwaresecuritymodule.Setup,
		vaultmanagedstorageaccount.Setup,
		vaultmanagedstorageaccountsastokendefinition.Setup,
		vaultsecret.Setup,
		attacheddatabaseconfiguration.Setup,
		clusterkusto.Setup,
		clustercustomermanagedkey.Setup,
		clusterprincipalassignment.Setup,
		database.Setup,
		databaseprincipal.Setup,
		databaseprincipalassignment.Setup,
		eventgriddataconnection.Setup,
		eventhubdataconnection.Setup,
		iothubdataconnection.Setup,
		backendaddresspool.Setup,
		backendaddresspooladdress.Setup,
		natpool.Setup,
		natrule.Setup,
		outboundrule.Setup,
		probe.Setup,
		rule.Setup,
		assignmentlighthouse.Setup,
		definition.Setup,
		virtualmachine.Setup,
		virtualmachinescaleset.Setup,
		networkgateway.Setup,
		analyticscluster.Setup,
		analyticsclustercustomermanagedkey.Setup,
		analyticsdataexportrule.Setup,
		analyticsdatasourcewindowsevent.Setup,
		analyticsdatasourcewindowsperformancecounter.Setup,
		analyticslinkedservice.Setup,
		analyticslinkedstorageaccount.Setup,
		analyticssavedsearch.Setup,
		analyticssolution.Setup,
		workspaceloganalytics.Setup,
		appactioncustom.Setup,
		appactionhttp.Setup,
		appintegrationaccount.Setup,
		appintegrationaccountagreement.Setup,
		appintegrationaccountassembly.Setup,
		appintegrationaccountbatchconfiguration.Setup,
		appintegrationaccountmap.Setup,
		appintegrationaccountpartner.Setup,
		appintegrationaccountschema.Setup,
		appintegrationaccountsession.Setup,
		apptriggercustom.Setup,
		apptriggerhttprequest.Setup,
		apptriggerrecurrence.Setup,
		appworkflow.Setup,
		learningcomputecluster.Setup,
		learningcomputeinstance.Setup,
		learningsynapsespark.Setup,
		learningworkspace.Setup,
		assignmentdedicatedhost.Setup,
		assignmentvirtualmachine.Setup,
		assignmentvirtualmachinescaleset.Setup,
		configurationmaintenance.Setup,
		applicationmanaged.Setup,
		applicationdefinition.Setup,
		disk.Setup,
		grouppolicyassignment.Setup,
		groupsubscriptionassociation.Setup,
		grouptemplatedeployment.Setup,
		lock.Setup,
		managementgroup.Setup,
		accountmaps.Setup,
		configurationmariadb.Setup,
		databasemariadb.Setup,
		firewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
		agreement.Setup,
		asset.Setup,
		assetfilter.Setup,
		contentkeypolicy.Setup,
		jobmedia.Setup,
		liveevent.Setup,
		liveeventoutput.Setup,
		servicesaccount.Setup,
		streamingendpoint.Setup,
		streaminglocator.Setup,
		streamingpolicy.Setup,
		transform.Setup,
		aaddiagnosticsetting.Setup,
		actiongroup.Setup,
		actionruleactiongroup.Setup,
		actionrulesuppression.Setup,
		activitylogalert.Setup,
		autoscalesetting.Setup,
		diagnosticsetting.Setup,
		metricalert.Setup,
		scheduledqueryrulesalert.Setup,
		scheduledqueryruleslog.Setup,
		smartdetectoralertrule.Setup,
		databasemssql.Setup,
		databaseextendedauditingpolicy.Setup,
		databasevulnerabilityassessmentrulebaseline.Setup,
		elasticpool.Setup,
		failovergroup.Setup,
		firewallrulemssql.Setup,
		jobagent.Setup,
		jobcredential.Setup,
		servermssql.Setup,
		serversecurityalertpolicy.Setup,
		servertransparentdataencryption.Setup,
		servervulnerabilityassessment.Setup,
		virtualmachinemssql.Setup,
		virtualnetworkrulemssql.Setup,
		activedirectoryadministrator.Setup,
		configurationmysql.Setup,
		databasemysql.Setup,
		firewallrulemysql.Setup,
		servermysql.Setup,
		serverkey.Setup,
		virtualnetworkrulemysql.Setup,
		gatewaynat.Setup,
		gatewaypublicipassociation.Setup,
		gatewaypublicipprefixassociation.Setup,
		accountnetapp.Setup,
		poolnetapp.Setup,
		snapshotnetapp.Setup,
		volume.Setup,
		connectionmonitor.Setup,
		ddosprotectionplan.Setup,
		interfaceapplicationsecuritygroupassociation.Setup,
		interfacebackendaddresspoolassociation.Setup,
		interfacenatruleassociation.Setup,
		interfacesecuritygroupassociation.Setup,
		ipgroup.Setup,
		loadbalancer.Setup,
		networkinterface.Setup,
		packetcapture.Setup,
		profilenetwork.Setup,
		securitygroupnetwork.Setup,
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
		watcher.Setup,
		watcherflowlog.Setup,
		hub.Setup,
		hubauthorizationrule.Setup,
		hubnamespace.Setup,
		virtualmachinescalesetorchestrated.Setup,
		capture.Setup,
		tositevpngateway.Setup,
		assignmentpolicy.Setup,
		definitionpolicy.Setup,
		remediation.Setup,
		setdefinition.Setup,
		virtualmachineconfigurationassignment.Setup,
		tenantconfiguration.Setup,
		activedirectoryadministratorpostgresql.Setup,
		configurationpostgresql.Setup,
		databasepostgresql.Setup,
		firewallrulepostgresql.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		serverpostgresql.Setup,
		virtualnetworkrulepostgresql.Setup,
		embedded.Setup,
		dnsaaaarecord.Setup,
		dnsarecord.Setup,
		dnscnamerecord.Setup,
		dnsmxrecord.Setup,
		dnsptrrecord.Setup,
		dnssrvrecord.Setup,
		dnstxtrecord.Setup,
		dnszone.Setup,
		dnszonevirtualnetworklink.Setup,
		endpointprivate.Setup,
		linkservice.Setup,
		providerconfig.Setup,
		placementgroup.Setup,
		ip.Setup,
		ipprefix.Setup,
		accountpurview.Setup,
		servicesvault.Setup,
		cacheredis.Setup,
		enterprisecluster.Setup,
		enterprisedatabase.Setup,
		firewallruleredis.Setup,
		linkedserver.Setup,
		hybridconnection.Setup,
		hybridconnectionauthorizationrule.Setup,
		namespacerelay.Setup,
		namespaceauthorizationrulerelay.Setup,
		policyassignment.Setup,
		providerregistration.Setup,
		resourcegrouptemplatedeployment.Setup,
		assignmentrole.Setup,
		definitionrole.Setup,
		servicesearch.Setup,
		centerassessment.Setup,
		centerassessmentmetadata.Setup,
		centerassessmentpolicy.Setup,
		centerautoprovisioning.Setup,
		centercontact.Setup,
		centerservervulnerabilityassessment.Setup,
		centersetting.Setup,
		centersubscriptionpricing.Setup,
		centerworkspace.Setup,
		alertrulefusion.Setup,
		alertrulemachinelearningbehavioranalytics.Setup,
		alertrulemssecurityincident.Setup,
		alertrulescheduled.Setup,
		dataconnectorawscloudtrail.Setup,
		dataconnectorazureactivedirectory.Setup,
		dataconnectorazureadvancedthreatprotection.Setup,
		dataconnectorazuresecuritycenter.Setup,
		dataconnectormicrosoftcloudappsecurity.Setup,
		dataconnectoroffice365.Setup,
		dataconnectorthreatintelligence.Setup,
		fabriccluster.Setup,
		fabricmeshapplication.Setup,
		fabricmeshlocalnetwork.Setup,
		fabricmeshsecret.Setup,
		fabricmeshsecretvalue.Setup,
		namespaceservicebus.Setup,
		namespaceauthorizationruleservicebus.Setup,
		namespacedisasterrecoveryconfigservicebus.Setup,
		namespacenetworkruleset.Setup,
		queue.Setup,
		queueauthorizationrule.Setup,
		subscriptionservicebus.Setup,
		subscriptionrule.Setup,
		topicservicebus.Setup,
		topicauthorizationrule.Setup,
		imageshared.Setup,
		imagegallery.Setup,
		imageversion.Setup,
		servicesignalr.Setup,
		servicenetworkacl.Setup,
		recoveryfabric.Setup,
		recoverynetworkmapping.Setup,
		recoveryprotectioncontainer.Setup,
		recoveryprotectioncontainermapping.Setup,
		recoveryreplicatedvm.Setup,
		recoveryreplicationpolicy.Setup,
		anchorsaccount.Setup,
		cloudactivedeployment.Setup,
		cloudapp.Setup,
		cloudappcosmosdbassociation.Setup,
		cloudappmysqlassociation.Setup,
		cloudappredisassociation.Setup,
		cloudcertificate.Setup,
		cloudcustomdomain.Setup,
		cloudjavadeployment.Setup,
		cloudservice.Setup,
		activedirectoryadministratorsql.Setup,
		databasesql.Setup,
		elasticpoolsql.Setup,
		firewallrulesql.Setup,
		manageddatabase.Setup,
		managedinstance.Setup,
		serversql.Setup,
		publickey.Setup,
		hcicluster.Setup,
		site.Setup,
		accountstorage.Setup,
		accountcustomermanagedkeystorage.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		datalakegen2filesystem.Setup,
		datalakegen2path.Setup,
		encryptionscope.Setup,
		managementpolicystorage.Setup,
		objectreplication.Setup,
		queuestorage.Setup,
		sharestorage.Setup,
		sharedirectory.Setup,
		sync.Setup,
		synccloudendpoint.Setup,
		tablestorage.Setup,
		tableentity.Setup,
		analyticsfunctionjavascriptudf.Setup,
		analyticsjob.Setup,
		analyticsoutputblob.Setup,
		analyticsoutputeventhub.Setup,
		analyticsoutputmssql.Setup,
		analyticsoutputservicebusqueue.Setup,
		analyticsoutputservicebustopic.Setup,
		analyticsreferenceinputblob.Setup,
		analyticsstreaminputblob.Setup,
		analyticsstreaminputeventhub.Setup,
		analyticsstreaminputiothub.Setup,
		policyassignmentsubscription.Setup,
		templatedeployment.Setup,
		firewallrulesynapse.Setup,
		integrationruntimeazure.Setup,
		integrationruntimeselfhosted.Setup,
		linkedservice.Setup,
		managedprivateendpoint.Setup,
		privatelinkhub.Setup,
		roleassignment.Setup,
		sparkpool.Setup,
		sqlpool.Setup,
		sqlpoolextendedauditingpolicy.Setup,
		sqlpoolsecurityalertpolicy.Setup,
		sqlpoolvulnerabilityassessment.Setup,
		workspacesynapse.Setup,
		workspaceextendedauditingpolicy.Setup,
		workspacekey.Setup,
		workspacesecurityalertpolicy.Setup,
		workspacevulnerabilityassessment.Setup,
		deployment.Setup,
		templatedeploymenttenant.Setup,
		managerendpoint.Setup,
		managerprofile.Setup,
		assignedidentity.Setup,
		analyzer.Setup,
		analyzeredgemodule.Setup,
		clustervmware.Setup,
		expressrouteauthorization.Setup,
		privatecloud.Setup,
		gatewayvpn.Setup,
		gatewayconnection.Setup,
		serverconfiguration.Setup,
		sitevpn.Setup,
		applicationfirewallpolicy.Setup,
		virtualmachinewindows.Setup,
		virtualmachinescalesetwindows.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
