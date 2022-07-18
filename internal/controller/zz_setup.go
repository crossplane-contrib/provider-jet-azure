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
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	activedirectorydomainservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/aad/activedirectorydomainservice"
	activedirectorydomainservicereplicaset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/aad/activedirectorydomainservicereplicaset"
	monitoraaddiagnosticsetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/aadiam/monitoraaddiagnosticsetting"
	monitoractionruleactiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/crossplane-contrib/provider-jet-azure/internal/controller/alertsmanagement/monitoractionrulesuppression"
	monitorsmartdetectoralertrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/alertsmanagement/monitorsmartdetectoralertrule"
	server "github.com/crossplane-contrib/provider-jet-azure/internal/controller/analysisservices/server"
	api "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/api"
	apidiagnostic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apidiagnostic"
	apioperation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apioperation"
	apioperationpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apioperationpolicy"
	apioperationtag "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apioperationtag"
	apipolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apipolicy"
	apirelease "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apirelease"
	apischema "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apischema"
	apiversionset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/apiversionset"
	authorizationserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/authorizationserver"
	backend "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/backend"
	certificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/certificate"
	customdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/customdomain"
	diagnostic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/diagnostic"
	emailtemplate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/emailtemplate"
	gateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/gateway"
	gatewayapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/gatewayapi"
	identityprovideraad "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityprovideraad"
	identityprovideraadb2c "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityprovideraadb2c"
	identityproviderfacebook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/identityprovidertwitter"
	logger "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/logger"
	management "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/management"
	namedvalue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/namedvalue"
	notificationrecipientemail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/notificationrecipientemail"
	openidconnectprovider "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/openidconnectprovider"
	policy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/policy"
	product "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/product"
	productapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/productapi"
	productpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/productpolicy"
	property "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/property"
	rediscache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/rediscache"
	subscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/subscription"
	tag "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/tag"
	user "github.com/crossplane-contrib/provider-jet-azure/internal/controller/apimanagement/user"
	configurationfeature "github.com/crossplane-contrib/provider-jet-azure/internal/controller/app/configurationfeature"
	configuration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appconfiguration/configuration"
	key "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appconfiguration/key"
	springcloudactivedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudactivedeployment"
	springcloudapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudappredisassociation"
	springcloudcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudcertificate"
	springcloudcustomdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudcustomdomain"
	springcloudjavadeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/appplatform/springcloudservice"
	provider "github.com/crossplane-contrib/provider-jet-azure/internal/controller/attestation/provider"
	managementgrouppolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/managementgrouppolicyassignment"
	managementlock "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/managementlock"
	policyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/policyassignment"
	policydefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/policydefinition"
	policysetdefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcepolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/resourcepolicyassignment"
	roleassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/roleassignment"
	roledefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/authorization/subscriptionpolicyassignment"
	account "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/account"
	certificateautomation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/automation/certificate"
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
	vmwarecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/avs/vmwarecluster"
	vmwareexpressrouteauthorization "github.com/crossplane-contrib/provider-jet-azure/internal/controller/avs/vmwareexpressrouteauthorization"
	vmwareprivatecloud "github.com/crossplane-contrib/provider-jet-azure/internal/controller/avs/vmwareprivatecloud"
	resourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azure/resourceproviderregistration"
	subscriptionazure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azure/subscription"
	cluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/azurestackhci/cluster"
	accountbatch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/account"
	application "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/application"
	certificatebatch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/certificate"
	job "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/job"
	pool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/batch/pool"
	assignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/blueprint/assignment"
	botchannelalexa "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelalexa"
	botchanneldirectline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchanneldirectline"
	botchanneldirectlinespeech "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchanneldirectlinespeech"
	botchannelemail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelemail"
	botchannelfacebook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelfacebook"
	botchannelline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelline"
	botchannelmsteams "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelmsteams"
	botchannelslack "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelslack"
	botchannelsms "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelsms"
	botchannelsregistration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelsregistration"
	botchannelwebchat "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botchannelwebchat"
	botconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botconnection"
	botwebapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/botservice/botwebapp"
	rediscachecache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cache/redislinkedserver"
	endpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/endpoint"
	endpointcustomdomain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/endpointcustomdomain"
	profile "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cdn/profile"
	appservicecertificateorder "github.com/crossplane-contrib/provider-jet-azure/internal/controller/certificateregistration/appservicecertificateorder"
	accountcognitiveservices "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cognitiveservices/account"
	accountcustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/cognitiveservices/accountcustomermanagedkey"
	service "github.com/crossplane-contrib/provider-jet-azure/internal/controller/communication/service"
	availabilityset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/availabilityset"
	dedicatedhost "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/dedicatedhost"
	diskaccess "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/diskaccess"
	diskencryptionset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/diskencryptionset"
	image "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/image"
	linuxvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/manageddisk"
	orchestratedvirtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/orchestratedvirtualmachinescaleset"
	policyvirtualmachineconfigurationassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/policyvirtualmachineconfigurationassignment"
	proximityplacementgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/proximityplacementgroup"
	sharedimage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/sharedimage"
	sharedimagegallery "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/sharedimagegallery"
	sharedimageversion "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/sharedimageversion"
	snapshot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/snapshot"
	sshpublickey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/sshpublickey"
	windowsvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/compute/windowsvirtualmachinescaleset"
	budgetresourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/consumption/budgetresourcegroup"
	budgetsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/consumption/budgetsubscription"
	registry "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerregistry/registry"
	scopemap "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerregistry/scopemap"
	token "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerregistry/token"
	webhook "github.com/crossplane-contrib/provider-jet-azure/internal/controller/containerregistry/webhook"
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
	exportresourcegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/costmanagement/exportresourcegroup"
	customprovider "github.com/crossplane-contrib/provider-jet-azure/internal/controller/customproviders/customprovider"
	factorylinkedservicecosmosdbmongoapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/data/factorylinkedservicecosmosdbmongoapi"
	device "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databoxedge/device"
	order "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databoxedge/order"
	workspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databricks/workspace"
	workspacecustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/databricks/workspacecustomermanagedkey"
	customdataset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/customdataset"
	dataflow "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/dataflow"
	datasetazureblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetazureblob"
	datasetbinary "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetbinary"
	datasetcosmosdbsqlapi "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetcosmosdbsqlapi"
	datasetdelimitedtext "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetdelimitedtext"
	datasethttp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasethttp"
	datasetjson "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetjson"
	datasetmysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetmysql"
	datasetparquet "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetparquet"
	datasetpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetpostgresql"
	datasetsnowflake "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetsnowflake"
	datasetsqlservertable "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/datasetsqlservertable"
	factory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/factory"
	integrationruntimeazure "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/integrationruntimeazuressis"
	integrationruntimemanaged "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/integrationruntimemanaged"
	integrationruntimeselfhosted "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/integrationruntimeselfhosted"
	linkedcustomservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedcustomservice"
	linkedserviceazureblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazureblobstorage"
	linkedserviceazuredatabricks "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazuredatabricks"
	linkedserviceazurefilestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazurefilestorage"
	linkedserviceazurefunction "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazurefunction"
	linkedserviceazuresearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazuresearch"
	linkedserviceazuresqldatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazuresqldatabase"
	linkedserviceazuretablestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceazuretablestorage"
	linkedservicecosmosdb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicecosmosdb"
	linkedservicedatalakestoragegen2 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicedatalakestoragegen2"
	linkedservicekeyvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicekeyvault"
	linkedservicekusto "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicekusto"
	linkedservicemysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicemysql"
	linkedserviceodata "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceodata"
	linkedservicepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicepostgresql"
	linkedservicesftp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicesftp"
	linkedservicesnowflake "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicesnowflake"
	linkedservicesqlserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicesqlserver"
	linkedservicesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedservicesynapse"
	linkedserviceweb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/linkedserviceweb"
	managedprivateendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/managedprivateendpoint"
	pipeline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/pipeline"
	triggerblobevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/triggerblobevent"
	triggercustomevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/triggercustomevent"
	triggerschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datafactory/triggerschedule"
	firewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datalakeanalytics/firewallrule"
	file "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datalakestore/file"
	firewallruledatalakestore "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datalakestore/firewallrule"
	store "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datalakestore/store"
	virtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datalakestore/virtualnetworkrule"
	databasemigrationproject "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backupinstancedisk"
	backupinstancepostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backuppolicydisk"
	backuppolicypostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backuppolicypostgresql"
	backupvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dataprotection/backupvault"
	accountdatashare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/account"
	datasetblobstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datasetblobstorage"
	datasetdatalakegen1 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datasetdatalakegen1"
	datasetdatalakegen2 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datasetkustodatabase"
	datashare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/datashare/datashare"
	configurationdbformariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformariadb/configuration"
	database "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformariadb/database"
	firewallruledbformariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformariadb/firewallrule"
	serverdbformariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformariadb/server"
	virtualnetworkruledbformariadb "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformariadb/virtualnetworkrule"
	activedirectoryadministrator "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/activedirectoryadministrator"
	configurationdbformysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/configuration"
	databasedbformysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/database"
	firewallruledbformysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/firewallrule"
	serverdbformysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/server"
	serverkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/serverkey"
	virtualnetworkruledbformysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministratordbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/configuration"
	databasedbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/server"
	serverkeydbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothub"
	iothubconsumergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubconsumergroup"
	iothubdps "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubdps"
	iothubdpscertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubenrichment"
	iothubfallbackroute "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubfallbackroute"
	iothubroute "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devices/iothubsharedaccesspolicy"
	devspacecontroller "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devspaces/devspacecontroller"
	globalvmshutdownschedule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/globalvmshutdownschedule"
	lab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/lab"
	linuxvirtualmachinedevtestlab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/linuxvirtualmachine"
	policydevtestlab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/policy"
	scheduledevtestlab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/schedule"
	virtualnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/virtualnetwork"
	windowsvirtualmachinedevtestlab "github.com/crossplane-contrib/provider-jet-azure/internal/controller/devtestlab/windowsvirtualmachine"
	endpointeventgrid "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digitaltwins/endpointeventgrid"
	endpointeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digitaltwins/endpointeventhub"
	endpointservicebus "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digitaltwins/endpointservicebus"
	instance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/digitaltwins/instance"
	domain "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/domain"
	domaintopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/domaintopic"
	eventsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/eventsubscription"
	systemtopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/systemtopic"
	systemtopiceventsubscription "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/systemtopiceventsubscription"
	topic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventgrid/topic"
	authorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/authorizationrule"
	clustereventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/cluster"
	consumergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespaceauthorizationrule"
	namespacecustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespacecustomermanagedkey"
	namespacedisasterrecoveryconfig "github.com/crossplane-contrib/provider-jet-azure/internal/controller/eventhub/namespacedisasterrecoveryconfig"
	dedicatedhardwaresecuritymodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hardwaresecuritymodules/dedicatedhardwaresecuritymodule"
	hadoopcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/hadoopcluster"
	hbasecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/hbasecluster"
	interactivequerycluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/interactivequerycluster"
	kafkacluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/kafkacluster"
	mlservicescluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/mlservicescluster"
	rservercluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/rservercluster"
	sparkcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/sparkcluster"
	stormcluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/hdinsight/stormcluster"
	healthbot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/healthbot/healthbot"
	healthcareservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/healthcareapis/healthcareservice"
	applicationinsights "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/applicationinsightssmartdetectionrule"
	applicationinsightswebtest "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/applicationinsightswebtest"
	monitoractiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitorautoscalesetting"
	monitordiagnosticsetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitormetricalert"
	monitorscheduledqueryrulesalert "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryruleslog "github.com/crossplane-contrib/provider-jet-azure/internal/controller/insights/monitorscheduledqueryruleslog"
	timeseriesinsightseventsourceeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iot/timeseriesinsightseventsourceeventhub"
	applicationiotcentral "github.com/crossplane-contrib/provider-jet-azure/internal/controller/iotcentral/application"
	accesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/accesspolicy"
	certificatekeyvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/certificateissuer"
	keykeyvault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/secret"
	vault "github.com/crossplane-contrib/provider-jet-azure/internal/controller/keyvault/vault"
	attacheddatabaseconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/attacheddatabaseconfiguration"
	clusterkusto "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/cluster"
	clustercustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/clustercustomermanagedkey"
	clusterprincipalassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/clusterprincipalassignment"
	databasekusto "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/database"
	databaseprincipal "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/databaseprincipal"
	databaseprincipalassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/iothubdataconnection"
	script "github.com/crossplane-contrib/provider-jet-azure/internal/controller/kusto/script"
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
	appstandard "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appstandard"
	apptriggercustom "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggercustom"
	apptriggerhttprequest "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/apptriggerrecurrence"
	appworkflow "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/appworkflow"
	integrationserviceenvironment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/logic/integrationserviceenvironment"
	computecluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machinelearningservices/computecluster"
	computeinstance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machinelearningservices/computeinstance"
	synapsespark "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machinelearningservices/synapsespark"
	workspacemachinelearningservices "github.com/crossplane-contrib/provider-jet-azure/internal/controller/machinelearningservices/workspace"
	maintenanceassignmentdedicatedhost "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceassignmentvirtualmachinescaleset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachinescaleset"
	maintenanceconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maintenance/maintenanceconfiguration"
	userassignedidentity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managedidentity/userassignedidentity"
	lighthouseassignment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managedservices/lighthouseassignment"
	lighthousedefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/managedservices/lighthousedefinition"
	managementgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/management/managementgroup"
	accountmaps "github.com/crossplane-contrib/provider-jet-azure/internal/controller/maps/account"
	marketplaceagreement "github.com/crossplane-contrib/provider-jet-azure/internal/controller/marketplaceordering/marketplaceagreement"
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
	videoanalyzer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/videoanalyzer"
	videoanalyzeredgemodule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/media/videoanalyzeredgemodule"
	spatialanchorsaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mixedreality/spatialanchorsaccount"
	flexibleservermysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/flexibleserver"
	flexibleserverconfigurationmysql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/mysql/flexibleserverconfiguration"
	accountnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/account"
	poolnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/pool"
	snapshotnetapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/snapshot"
	volume "github.com/crossplane-contrib/provider-jet-azure/internal/controller/netapp/volume"
	applicationgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/applicationgateway"
	applicationsecuritygroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/applicationsecuritygroup"
	bastionhost "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/bastionhost"
	connectionmonitor "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/connectionmonitor"
	ddosprotectionplan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/ddosprotectionplan"
	dnsaaaarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnsaaaarecord"
	dnsarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnsarecord"
	dnscaarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnscaarecord"
	dnscnamerecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnscnamerecord"
	dnsmxrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnsmxrecord"
	dnsnsrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnsnsrecord"
	dnsptrrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnsptrrecord"
	dnssrvrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnssrvrecord"
	dnstxtrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnstxtrecord"
	dnszone "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/dnszone"
	expressroutecircuit "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressrouteconnection"
	expressroutegateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressroutegateway"
	expressrouteport "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/expressrouteport"
	firewall "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewall"
	firewallapplicationrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/frontdoorrulesengine"
	interfaceapplicationsecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfaceapplicationsecuritygroupassociation"
	interfacebackendaddresspoolassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacebackendaddresspoolassociation"
	interfacenatruleassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacenatruleassociation"
	interfacesecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/interfacesecuritygroupassociation"
	ipgroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/ipgroup"
	lbbackendaddresspool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbbackendaddresspool"
	lbbackendaddresspooladdress "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbbackendaddresspooladdress"
	lbnatpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbnatpool"
	lbnatrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbnatrule"
	lboutboundrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lboutboundrule"
	lbprobe "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbprobe"
	lbrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/lbrule"
	loadbalancer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/loadbalancer"
	localnetworkgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/localnetworkgateway"
	natgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/natgateway"
	natgatewaypublicipassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/networkinterface"
	networkpacketcapture "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/networkpacketcapture"
	packetcapture "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/packetcapture"
	pointtositevpngateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednsaaaarecord"
	privatednsarecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednsarecord"
	privatednscnamerecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednscnamerecord"
	privatednsmxrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednsmxrecord"
	privatednsptrrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednsptrrecord"
	privatednssrvrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednssrvrecord"
	privatednstxtrecord "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednstxtrecord"
	privatednszone "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privateendpoint"
	privatelinkservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/privatelinkservice"
	profilenetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/profile"
	publicip "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/publicip"
	publicipprefix "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/publicipprefix"
	securitygroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/securitygroup"
	subnet "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	trafficmanagerendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/trafficmanagerendpoint"
	trafficmanagerprofile "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/trafficmanagerprofile"
	virtualnetworknetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/virtualwan"
	vpngateway "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/vpngateway"
	vpngatewayconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/vpnserverconfiguration"
	vpnsite "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/vpnsite"
	watcher "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/watcher"
	watcherflowlog "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/network/webapplicationfirewallpolicy"
	authorizationrulenotificationhubs "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notificationhubs/authorizationrule"
	namespace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notificationhubs/namespace"
	notificationhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/notificationhubs/notificationhub"
	loganalyticscluster "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticscluster"
	loganalyticsclustercustomermanagedkey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticsclustercustomermanagedkey"
	loganalyticsdataexportrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticssavedsearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticssavedsearch"
	loganalyticssolution "github.com/crossplane-contrib/provider-jet-azure/internal/controller/operationalinsights/loganalyticssolution"
	policyremediation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/policyinsights/policyremediation"
	dashboard "github.com/crossplane-contrib/provider-jet-azure/internal/controller/portal/dashboard"
	tenantconfiguration "github.com/crossplane-contrib/provider-jet-azure/internal/controller/portal/tenantconfiguration"
	powerbiembedded "github.com/crossplane-contrib/provider-jet-azure/internal/controller/powerbidedicated/powerbiembedded"
	providerconfig "github.com/crossplane-contrib/provider-jet-azure/internal/controller/providerconfig"
	accountpurview "github.com/crossplane-contrib/provider-jet-azure/internal/controller/purview/account"
	backupcontainerstorageaccount "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/backuppolicyvm"
	backupprotectedfileshare "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicatedvm "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoveryreplicatedvm"
	siterecoveryreplicationpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/siterecoveryreplicationpolicy"
	vaultrecoveryservices "github.com/crossplane-contrib/provider-jet-azure/internal/controller/recoveryservices/vault"
	hybridconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/hybridconnectionauthorizationrule"
	namespacerelay "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/namespace"
	namespaceauthorizationrulerelay "github.com/crossplane-contrib/provider-jet-azure/internal/controller/relay/namespaceauthorizationrule"
	managementgroupsubscriptionassociation "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/managementgroupsubscriptionassociation"
	managementgrouptemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/managementgrouptemplatedeployment"
	resourcegrouptemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/subscriptiontemplatedeployment"
	templatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/templatedeployment"
	tenanttemplatedeployment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/resources/tenanttemplatedeployment"
	servicesearch "github.com/crossplane-contrib/provider-jet-azure/internal/controller/search/service"
	advancedthreatprotection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/iotsecuritysolution"
	securitycenterassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterassessment"
	securitycenterassessmentmetadata "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterassessmentmetadata"
	securitycenterassessmentpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterservervulnerabilityassessment"
	securitycentersetting "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/crossplane-contrib/provider-jet-azure/internal/controller/security/securitycenterworkspace"
	sentinelalertrulefusion "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentinelalertrulemssecurityincident"
	sentinelalertrulescheduled "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentinelalertrulescheduled"
	sentineldataconnectorawscloudtrail "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectorawscloudtrail"
	sentineldataconnectorazureactivedirectory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectorazureactivedirectory"
	sentineldataconnectorazureadvancedthreatprotection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectorazureadvancedthreatprotection"
	sentineldataconnectorazuresecuritycenter "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectorazuresecuritycenter"
	sentineldataconnectormicrosoftcloudappsecurity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectormicrosoftcloudappsecurity"
	sentineldataconnectoroffice365 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectoroffice365"
	sentineldataconnectorthreatintelligence "github.com/crossplane-contrib/provider-jet-azure/internal/controller/securityinsights/sentineldataconnectorthreatintelligence"
	automationrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sentinel/automationrule"
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
	clusterservicefabric "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicefabric/cluster"
	applicationservicefabricmesh "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicefabricmesh/application"
	localnetwork "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicefabricmesh/localnetwork"
	secretservicefabricmesh "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicefabricmesh/secret"
	secretvalue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/servicefabricmesh/secretvalue"
	networkacl "github.com/crossplane-contrib/provider-jet-azure/internal/controller/signalrservice/networkacl"
	servicesignalrservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/signalrservice/service"
	managedapplication "github.com/crossplane-contrib/provider-jet-azure/internal/controller/solutions/managedapplication"
	managedapplicationdefinition "github.com/crossplane-contrib/provider-jet-azure/internal/controller/solutions/managedapplicationdefinition"
	activedirectoryadministratorsql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/activedirectoryadministrator"
	databasesql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/database"
	elasticpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/elasticpool"
	firewallrulesql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/firewallrule"
	manageddatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/manageddatabase"
	managedinstance "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/managedinstance"
	mssqldatabase "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlfirewallrule"
	mssqljobagent "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqljobagent"
	mssqljobcredential "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqljobcredential"
	mssqlserver "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlserver"
	mssqlserversecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/mssqlvirtualnetworkrule"
	serversql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sql/server"
	mssqlvirtualmachine "github.com/crossplane-contrib/provider-jet-azure/internal/controller/sqlvirtualmachine/mssqlvirtualmachine"
	accountstorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/account"
	accountcustomermanagedkeystorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/accountcustomermanagedkey"
	accountnetworkrules "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/datalakegen2filesystem"
	datalakegen2path "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/datalakegen2path"
	encryptionscope "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/encryptionscope"
	managementpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/objectreplication"
	queuestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/queue"
	share "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/share"
	sharedirectory "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/sharedirectory"
	tablestorage "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/table"
	tableentity "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storage/tableentity"
	hpccache "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagecache/hpccachenfstarget"
	cloudendpoint "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagesync/cloudendpoint"
	storagesync "github.com/crossplane-contrib/provider-jet-azure/internal/controller/storagesync/storagesync"
	analyticsoutputtable "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsoutputtable"
	analyticsreferenceinputmssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/stream/analyticsreferenceinputmssql"
	functionjavascriptudf "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/functionjavascriptudf"
	jobstreamanalytics "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/job"
	outputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/outputblob"
	outputeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/outputeventhub"
	outputmssql "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/outputmssql"
	outputservicebusqueue "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/outputservicebustopic"
	referenceinputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/referenceinputblob"
	streaminputblob "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/streamanalytics/streaminputiothub"
	firewallrulesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/managedprivateendpoint"
	privatelinkhub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/privatelinkhub"
	roleassignmentsynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/roleassignment"
	sparkpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sparkpool"
	sqlpool "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolvulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolvulnerabilityassessment"
	sqlpoolvulnerabilityassessmentbaseline "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/sqlpoolvulnerabilityassessmentbaseline"
	workspacesynapse "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspace"
	workspaceextendedauditingpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspaceextendedauditingpolicy"
	workspacekey "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacekey"
	workspacesecurityalertpolicy "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacesecurityalertpolicy"
	workspacevulnerabilityassessment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/synapse/workspacevulnerabilityassessment"
	accesspolicytimeseriesinsights "github.com/crossplane-contrib/provider-jet-azure/internal/controller/timeseriesinsights/accesspolicy"
	eventsourceiothub "github.com/crossplane-contrib/provider-jet-azure/internal/controller/timeseriesinsights/eventsourceiothub"
	gen2environment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/timeseriesinsights/gen2environment"
	referencedataset "github.com/crossplane-contrib/provider-jet-azure/internal/controller/timeseriesinsights/referencedataset"
	standardenvironment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/timeseriesinsights/standardenvironment"
	appservice "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservice"
	appserviceactiveslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceactiveslot"
	appservicecertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicecertificate"
	appservicecertificatebinding "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicecertificatebinding"
	appservicecustomhostnamebinding "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicecustomhostnamebinding"
	appserviceenvironment "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceenvironment"
	appserviceenvironmentv3 "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceenvironmentv3"
	appservicehybridconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicehybridconnection"
	appservicemanagedcertificate "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicemanagedcertificate"
	appserviceplan "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceplan"
	appserviceslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceslot"
	appserviceslotvirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appserviceslotvirtualnetworkswiftconnection"
	appservicesourcecontroltoken "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicesourcecontroltoken"
	appservicevirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/appservicevirtualnetworkswiftconnection"
	functionapp "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/functionapp"
	functionappslot "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/functionappslot"
	staticsite "github.com/crossplane-contrib/provider-jet-azure/internal/controller/web/staticsite"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectorydomainservice.Setup,
		activedirectorydomainservicereplicaset.Setup,
		monitoraaddiagnosticsetting.Setup,
		monitoractionruleactiongroup.Setup,
		monitoractionrulesuppression.Setup,
		monitorsmartdetectoralertrule.Setup,
		server.Setup,
		api.Setup,
		apidiagnostic.Setup,
		apioperation.Setup,
		apioperationpolicy.Setup,
		apioperationtag.Setup,
		apipolicy.Setup,
		apirelease.Setup,
		apischema.Setup,
		apiversionset.Setup,
		authorizationserver.Setup,
		backend.Setup,
		certificate.Setup,
		customdomain.Setup,
		diagnostic.Setup,
		emailtemplate.Setup,
		gateway.Setup,
		gatewayapi.Setup,
		identityprovideraad.Setup,
		identityprovideraadb2c.Setup,
		identityproviderfacebook.Setup,
		identityprovidergoogle.Setup,
		identityprovidermicrosoft.Setup,
		identityprovidertwitter.Setup,
		logger.Setup,
		management.Setup,
		namedvalue.Setup,
		notificationrecipientemail.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		product.Setup,
		productapi.Setup,
		productpolicy.Setup,
		property.Setup,
		rediscache.Setup,
		subscription.Setup,
		tag.Setup,
		user.Setup,
		configurationfeature.Setup,
		configuration.Setup,
		key.Setup,
		springcloudactivedeployment.Setup,
		springcloudapp.Setup,
		springcloudappcosmosdbassociation.Setup,
		springcloudappmysqlassociation.Setup,
		springcloudappredisassociation.Setup,
		springcloudcertificate.Setup,
		springcloudcustomdomain.Setup,
		springcloudjavadeployment.Setup,
		springcloudservice.Setup,
		provider.Setup,
		managementgrouppolicyassignment.Setup,
		managementlock.Setup,
		policyassignment.Setup,
		policydefinition.Setup,
		policysetdefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcepolicyassignment.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		subscriptionpolicyassignment.Setup,
		account.Setup,
		certificateautomation.Setup,
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
		vmwarecluster.Setup,
		vmwareexpressrouteauthorization.Setup,
		vmwareprivatecloud.Setup,
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscriptionazure.Setup,
		cluster.Setup,
		accountbatch.Setup,
		application.Setup,
		certificatebatch.Setup,
		job.Setup,
		pool.Setup,
		assignment.Setup,
		botchannelalexa.Setup,
		botchanneldirectline.Setup,
		botchanneldirectlinespeech.Setup,
		botchannelemail.Setup,
		botchannelfacebook.Setup,
		botchannelline.Setup,
		botchannelmsteams.Setup,
		botchannelslack.Setup,
		botchannelsms.Setup,
		botchannelsregistration.Setup,
		botchannelwebchat.Setup,
		botconnection.Setup,
		botwebapp.Setup,
		rediscachecache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		endpoint.Setup,
		endpointcustomdomain.Setup,
		profile.Setup,
		appservicecertificateorder.Setup,
		accountcognitiveservices.Setup,
		accountcustomermanagedkey.Setup,
		service.Setup,
		availabilityset.Setup,
		dedicatedhost.Setup,
		diskaccess.Setup,
		diskencryptionset.Setup,
		image.Setup,
		linuxvirtualmachine.Setup,
		linuxvirtualmachinescaleset.Setup,
		manageddisk.Setup,
		orchestratedvirtualmachinescaleset.Setup,
		policyvirtualmachineconfigurationassignment.Setup,
		proximityplacementgroup.Setup,
		sharedimage.Setup,
		sharedimagegallery.Setup,
		sharedimageversion.Setup,
		snapshot.Setup,
		sshpublickey.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
		budgetresourcegroup.Setup,
		budgetsubscription.Setup,
		registry.Setup,
		scopemap.Setup,
		token.Setup,
		webhook.Setup,
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
		exportresourcegroup.Setup,
		customprovider.Setup,
		factorylinkedservicecosmosdbmongoapi.Setup,
		device.Setup,
		order.Setup,
		workspace.Setup,
		workspacecustomermanagedkey.Setup,
		customdataset.Setup,
		dataflow.Setup,
		datasetazureblob.Setup,
		datasetbinary.Setup,
		datasetcosmosdbsqlapi.Setup,
		datasetdelimitedtext.Setup,
		datasethttp.Setup,
		datasetjson.Setup,
		datasetmysql.Setup,
		datasetparquet.Setup,
		datasetpostgresql.Setup,
		datasetsnowflake.Setup,
		datasetsqlservertable.Setup,
		factory.Setup,
		integrationruntimeazure.Setup,
		integrationruntimeazuressis.Setup,
		integrationruntimemanaged.Setup,
		integrationruntimeselfhosted.Setup,
		linkedcustomservice.Setup,
		linkedserviceazureblobstorage.Setup,
		linkedserviceazuredatabricks.Setup,
		linkedserviceazurefilestorage.Setup,
		linkedserviceazurefunction.Setup,
		linkedserviceazuresearch.Setup,
		linkedserviceazuresqldatabase.Setup,
		linkedserviceazuretablestorage.Setup,
		linkedservicecosmosdb.Setup,
		linkedservicedatalakestoragegen2.Setup,
		linkedservicekeyvault.Setup,
		linkedservicekusto.Setup,
		linkedservicemysql.Setup,
		linkedserviceodata.Setup,
		linkedservicepostgresql.Setup,
		linkedservicesftp.Setup,
		linkedservicesnowflake.Setup,
		linkedservicesqlserver.Setup,
		linkedservicesynapse.Setup,
		linkedserviceweb.Setup,
		managedprivateendpoint.Setup,
		pipeline.Setup,
		triggerblobevent.Setup,
		triggercustomevent.Setup,
		triggerschedule.Setup,
		firewallrule.Setup,
		file.Setup,
		firewallruledatalakestore.Setup,
		store.Setup,
		virtualnetworkrule.Setup,
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backupinstancepostgresql.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backuppolicypostgresql.Setup,
		backupvault.Setup,
		accountdatashare.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen1.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
		configurationdbformariadb.Setup,
		database.Setup,
		firewallruledbformariadb.Setup,
		serverdbformariadb.Setup,
		virtualnetworkruledbformariadb.Setup,
		activedirectoryadministrator.Setup,
		configurationdbformysql.Setup,
		databasedbformysql.Setup,
		firewallruledbformysql.Setup,
		serverdbformysql.Setup,
		serverkey.Setup,
		virtualnetworkruledbformysql.Setup,
		activedirectoryadministratordbforpostgresql.Setup,
		configurationdbforpostgresql.Setup,
		databasedbforpostgresql.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		serverdbforpostgresql.Setup,
		serverkeydbforpostgresql.Setup,
		virtualnetworkruledbforpostgresql.Setup,
		iothub.Setup,
		iothubconsumergroup.Setup,
		iothubdps.Setup,
		iothubdpscertificate.Setup,
		iothubdpssharedaccesspolicy.Setup,
		iothubendpointeventhub.Setup,
		iothubendpointservicebusqueue.Setup,
		iothubendpointservicebustopic.Setup,
		iothubendpointstoragecontainer.Setup,
		iothubenrichment.Setup,
		iothubfallbackroute.Setup,
		iothubroute.Setup,
		iothubsharedaccesspolicy.Setup,
		devspacecontroller.Setup,
		globalvmshutdownschedule.Setup,
		lab.Setup,
		linuxvirtualmachinedevtestlab.Setup,
		policydevtestlab.Setup,
		scheduledevtestlab.Setup,
		virtualnetwork.Setup,
		windowsvirtualmachinedevtestlab.Setup,
		endpointeventgrid.Setup,
		endpointeventhub.Setup,
		endpointservicebus.Setup,
		instance.Setup,
		domain.Setup,
		domaintopic.Setup,
		eventsubscription.Setup,
		systemtopic.Setup,
		systemtopiceventsubscription.Setup,
		topic.Setup,
		authorizationrule.Setup,
		clustereventhub.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacecustomermanagedkey.Setup,
		namespacedisasterrecoveryconfig.Setup,
		dedicatedhardwaresecuritymodule.Setup,
		hadoopcluster.Setup,
		hbasecluster.Setup,
		interactivequerycluster.Setup,
		kafkacluster.Setup,
		mlservicescluster.Setup,
		rservercluster.Setup,
		sparkcluster.Setup,
		stormcluster.Setup,
		healthbot.Setup,
		healthcareservice.Setup,
		applicationinsights.Setup,
		applicationinsightsanalyticsitem.Setup,
		applicationinsightsapikey.Setup,
		applicationinsightssmartdetectionrule.Setup,
		applicationinsightswebtest.Setup,
		monitoractiongroup.Setup,
		monitoractivitylogalert.Setup,
		monitorautoscalesetting.Setup,
		monitordiagnosticsetting.Setup,
		monitormetricalert.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryruleslog.Setup,
		timeseriesinsightseventsourceeventhub.Setup,
		applicationiotcentral.Setup,
		accesspolicy.Setup,
		certificatekeyvault.Setup,
		certificateissuer.Setup,
		keykeyvault.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		attacheddatabaseconfiguration.Setup,
		clusterkusto.Setup,
		clustercustomermanagedkey.Setup,
		clusterprincipalassignment.Setup,
		databasekusto.Setup,
		databaseprincipal.Setup,
		databaseprincipalassignment.Setup,
		eventgriddataconnection.Setup,
		eventhubdataconnection.Setup,
		iothubdataconnection.Setup,
		script.Setup,
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
		appstandard.Setup,
		apptriggercustom.Setup,
		apptriggerhttprequest.Setup,
		apptriggerrecurrence.Setup,
		appworkflow.Setup,
		integrationserviceenvironment.Setup,
		computecluster.Setup,
		computeinstance.Setup,
		synapsespark.Setup,
		workspacemachinelearningservices.Setup,
		maintenanceassignmentdedicatedhost.Setup,
		maintenanceassignmentvirtualmachine.Setup,
		maintenanceassignmentvirtualmachinescaleset.Setup,
		maintenanceconfiguration.Setup,
		userassignedidentity.Setup,
		lighthouseassignment.Setup,
		lighthousedefinition.Setup,
		managementgroup.Setup,
		accountmaps.Setup,
		marketplaceagreement.Setup,
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
		videoanalyzer.Setup,
		videoanalyzeredgemodule.Setup,
		spatialanchorsaccount.Setup,
		flexibleservermysql.Setup,
		flexibleserverconfigurationmysql.Setup,
		accountnetapp.Setup,
		poolnetapp.Setup,
		snapshotnetapp.Setup,
		volume.Setup,
		applicationgateway.Setup,
		applicationsecuritygroup.Setup,
		bastionhost.Setup,
		connectionmonitor.Setup,
		ddosprotectionplan.Setup,
		dnsaaaarecord.Setup,
		dnsarecord.Setup,
		dnscaarecord.Setup,
		dnscnamerecord.Setup,
		dnsmxrecord.Setup,
		dnsnsrecord.Setup,
		dnsptrrecord.Setup,
		dnssrvrecord.Setup,
		dnstxtrecord.Setup,
		dnszone.Setup,
		expressroutecircuit.Setup,
		expressroutecircuitauthorization.Setup,
		expressroutecircuitconnection.Setup,
		expressroutecircuitpeering.Setup,
		expressrouteconnection.Setup,
		expressroutegateway.Setup,
		expressrouteport.Setup,
		firewall.Setup,
		firewallapplicationrulecollection.Setup,
		firewallnatrulecollection.Setup,
		firewallnetworkrulecollection.Setup,
		firewallpolicy.Setup,
		firewallpolicyrulecollectiongroup.Setup,
		frontdoor.Setup,
		frontdoorcustomhttpsconfiguration.Setup,
		frontdoorfirewallpolicy.Setup,
		frontdoorrulesengine.Setup,
		interfaceapplicationsecuritygroupassociation.Setup,
		interfacebackendaddresspoolassociation.Setup,
		interfacenatruleassociation.Setup,
		interfacesecuritygroupassociation.Setup,
		ipgroup.Setup,
		lbbackendaddresspool.Setup,
		lbbackendaddresspooladdress.Setup,
		lbnatpool.Setup,
		lbnatrule.Setup,
		lboutboundrule.Setup,
		lbprobe.Setup,
		lbrule.Setup,
		loadbalancer.Setup,
		localnetworkgateway.Setup,
		natgateway.Setup,
		natgatewaypublicipassociation.Setup,
		natgatewaypublicipprefixassociation.Setup,
		networkinterface.Setup,
		networkpacketcapture.Setup,
		packetcapture.Setup,
		pointtositevpngateway.Setup,
		privatednsaaaarecord.Setup,
		privatednsarecord.Setup,
		privatednscnamerecord.Setup,
		privatednsmxrecord.Setup,
		privatednsptrrecord.Setup,
		privatednssrvrecord.Setup,
		privatednstxtrecord.Setup,
		privatednszone.Setup,
		privatednszonevirtualnetworklink.Setup,
		privateendpoint.Setup,
		privatelinkservice.Setup,
		profilenetwork.Setup,
		publicip.Setup,
		publicipprefix.Setup,
		securitygroup.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		trafficmanagerendpoint.Setup,
		trafficmanagerprofile.Setup,
		virtualnetworknetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		vpngateway.Setup,
		vpngatewayconnection.Setup,
		vpnserverconfiguration.Setup,
		vpnsite.Setup,
		watcher.Setup,
		watcherflowlog.Setup,
		webapplicationfirewallpolicy.Setup,
		authorizationrulenotificationhubs.Setup,
		namespace.Setup,
		notificationhub.Setup,
		loganalyticscluster.Setup,
		loganalyticsclustercustomermanagedkey.Setup,
		loganalyticsdataexportrule.Setup,
		loganalyticsdatasourcewindowsevent.Setup,
		loganalyticsdatasourcewindowsperformancecounter.Setup,
		loganalyticslinkedservice.Setup,
		loganalyticslinkedstorageaccount.Setup,
		loganalyticssavedsearch.Setup,
		loganalyticssolution.Setup,
		policyremediation.Setup,
		dashboard.Setup,
		tenantconfiguration.Setup,
		powerbiembedded.Setup,
		providerconfig.Setup,
		accountpurview.Setup,
		backupcontainerstorageaccount.Setup,
		backuppolicyfileshare.Setup,
		backuppolicyvm.Setup,
		backupprotectedfileshare.Setup,
		backupprotectedvm.Setup,
		siterecoveryfabric.Setup,
		siterecoverynetworkmapping.Setup,
		siterecoveryprotectioncontainer.Setup,
		siterecoveryprotectioncontainermapping.Setup,
		siterecoveryreplicatedvm.Setup,
		siterecoveryreplicationpolicy.Setup,
		vaultrecoveryservices.Setup,
		hybridconnection.Setup,
		hybridconnectionauthorizationrule.Setup,
		namespacerelay.Setup,
		namespaceauthorizationrulerelay.Setup,
		managementgroupsubscriptionassociation.Setup,
		managementgrouptemplatedeployment.Setup,
		resourcegrouptemplatedeployment.Setup,
		subscriptiontemplatedeployment.Setup,
		templatedeployment.Setup,
		tenanttemplatedeployment.Setup,
		servicesearch.Setup,
		advancedthreatprotection.Setup,
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		securitycenterassessment.Setup,
		securitycenterassessmentmetadata.Setup,
		securitycenterassessmentpolicy.Setup,
		securitycenterautoprovisioning.Setup,
		securitycentercontact.Setup,
		securitycenterservervulnerabilityassessment.Setup,
		securitycentersetting.Setup,
		securitycentersubscriptionpricing.Setup,
		securitycenterworkspace.Setup,
		sentinelalertrulefusion.Setup,
		sentinelalertrulemachinelearningbehavioranalytics.Setup,
		sentinelalertrulemssecurityincident.Setup,
		sentinelalertrulescheduled.Setup,
		sentineldataconnectorawscloudtrail.Setup,
		sentineldataconnectorazureactivedirectory.Setup,
		sentineldataconnectorazureadvancedthreatprotection.Setup,
		sentineldataconnectorazuresecuritycenter.Setup,
		sentineldataconnectormicrosoftcloudappsecurity.Setup,
		sentineldataconnectoroffice365.Setup,
		sentineldataconnectorthreatintelligence.Setup,
		automationrule.Setup,
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
		clusterservicefabric.Setup,
		applicationservicefabricmesh.Setup,
		localnetwork.Setup,
		secretservicefabricmesh.Setup,
		secretvalue.Setup,
		networkacl.Setup,
		servicesignalrservice.Setup,
		managedapplication.Setup,
		managedapplicationdefinition.Setup,
		activedirectoryadministratorsql.Setup,
		databasesql.Setup,
		elasticpool.Setup,
		firewallrulesql.Setup,
		manageddatabase.Setup,
		managedinstance.Setup,
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqljobagent.Setup,
		mssqljobcredential.Setup,
		mssqlserver.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		mssqlvirtualnetworkrule.Setup,
		serversql.Setup,
		mssqlvirtualmachine.Setup,
		accountstorage.Setup,
		accountcustomermanagedkeystorage.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		datalakegen2filesystem.Setup,
		datalakegen2path.Setup,
		encryptionscope.Setup,
		managementpolicy.Setup,
		objectreplication.Setup,
		queuestorage.Setup,
		share.Setup,
		sharedirectory.Setup,
		tablestorage.Setup,
		tableentity.Setup,
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
		cloudendpoint.Setup,
		storagesync.Setup,
		analyticsoutputtable.Setup,
		analyticsreferenceinputmssql.Setup,
		functionjavascriptudf.Setup,
		jobstreamanalytics.Setup,
		outputblob.Setup,
		outputeventhub.Setup,
		outputmssql.Setup,
		outputservicebusqueue.Setup,
		outputservicebustopic.Setup,
		referenceinputblob.Setup,
		streaminputblob.Setup,
		streaminputeventhub.Setup,
		streaminputiothub.Setup,
		firewallrulesynapse.Setup,
		integrationruntimeazuresynapse.Setup,
		integrationruntimeselfhostedsynapse.Setup,
		linkedservice.Setup,
		managedprivateendpointsynapse.Setup,
		privatelinkhub.Setup,
		roleassignmentsynapse.Setup,
		sparkpool.Setup,
		sqlpool.Setup,
		sqlpoolextendedauditingpolicy.Setup,
		sqlpoolsecurityalertpolicy.Setup,
		sqlpoolvulnerabilityassessment.Setup,
		sqlpoolvulnerabilityassessmentbaseline.Setup,
		workspacesynapse.Setup,
		workspaceextendedauditingpolicy.Setup,
		workspacekey.Setup,
		workspacesecurityalertpolicy.Setup,
		workspacevulnerabilityassessment.Setup,
		accesspolicytimeseriesinsights.Setup,
		eventsourceiothub.Setup,
		gen2environment.Setup,
		referencedataset.Setup,
		standardenvironment.Setup,
		appservice.Setup,
		appserviceactiveslot.Setup,
		appservicecertificate.Setup,
		appservicecertificatebinding.Setup,
		appservicecustomhostnamebinding.Setup,
		appserviceenvironment.Setup,
		appserviceenvironmentv3.Setup,
		appservicehybridconnection.Setup,
		appservicemanagedcertificate.Setup,
		appserviceplan.Setup,
		appserviceslot.Setup,
		appserviceslotvirtualnetworkswiftconnection.Setup,
		appservicesourcecontroltoken.Setup,
		appservicevirtualnetworkswiftconnection.Setup,
		functionapp.Setup,
		functionappslot.Setup,
		staticsite.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
