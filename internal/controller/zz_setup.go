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

	activedirectorydomainservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/active/activedirectorydomainservice"
	activedirectorydomainservicereplicaset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/active/activedirectorydomainservicereplicaset"
	advancedthreatprotection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/advanced/advancedthreatprotection"
	analysisservicesserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/analysis/analysisservicesserver"
	apimanagement "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagement"
	apimanagementapi "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapi"
	apimanagementapidiagnostic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapidiagnostic"
	apimanagementapioperation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapioperation"
	apimanagementapioperationpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapioperationpolicy"
	apimanagementapioperationtag "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapioperationtag"
	apimanagementapipolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapipolicy"
	apimanagementapirelease "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapirelease"
	apimanagementapischema "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapischema"
	apimanagementapiversionset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementapiversionset"
	apimanagementauthorizationserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementauthorizationserver"
	apimanagementbackend "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementbackend"
	apimanagementcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementcertificate"
	apimanagementcustomdomain "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementcustomdomain"
	apimanagementdiagnostic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementdiagnostic"
	apimanagementemailtemplate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementemailtemplate"
	apimanagementgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementgateway"
	apimanagementgatewayapi "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementgatewayapi"
	apimanagementgroupuser "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementgroupuser"
	apimanagementidentityprovideraad "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityprovideraad"
	apimanagementidentityprovideraadb2c "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityprovideraadb2c"
	apimanagementidentityproviderfacebook "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityproviderfacebook"
	apimanagementidentityprovidergoogle "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityprovidergoogle"
	apimanagementidentityprovidermicrosoft "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityprovidermicrosoft"
	apimanagementidentityprovidertwitter "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementidentityprovidertwitter"
	apimanagementlogger "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementlogger"
	apimanagementnamedvalue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementnamedvalue"
	apimanagementnotificationrecipientemail "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementnotificationrecipientemail"
	apimanagementopenidconnectprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementopenidconnectprovider"
	apimanagementpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementpolicy"
	apimanagementproduct "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementproduct"
	apimanagementproductapi "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementproductapi"
	apimanagementproductpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementproductpolicy"
	apimanagementproperty "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementproperty"
	apimanagementrediscache "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementrediscache"
	apimanagementsubscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementsubscription"
	apimanagementtag "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementtag"
	apimanagementuser "github.com/crossplane-contrib/provider-tf-azure/internal/controller/api/apimanagementuser"
	appconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appconfiguration"
	appconfigurationkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appconfigurationkey"
	appservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservice"
	appserviceactiveslot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceactiveslot"
	appservicecertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicecertificate"
	appservicecertificatebinding "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicecertificatebinding"
	appservicecertificateorder "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicecertificateorder"
	appservicecustomhostnamebinding "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicecustomhostnamebinding"
	appserviceenvironment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceenvironment"
	appserviceenvironmentv3 "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceenvironmentv3"
	appservicehybridconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicehybridconnection"
	appservicemanagedcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicemanagedcertificate"
	appserviceplan "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceplan"
	appserviceslot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceslot"
	appserviceslotvirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appserviceslotvirtualnetworkswiftconnection"
	appservicesourcecontroltoken "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicesourcecontroltoken"
	appservicevirtualnetworkswiftconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/app/appservicevirtualnetworkswiftconnection"
	applicationgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationgateway"
	applicationinsights "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationinsights"
	applicationinsightsanalyticsitem "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationinsightssmartdetectionrule"
	applicationinsightswebtest "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationinsightswebtest"
	applicationsecuritygroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/application/applicationsecuritygroup"
	attestationprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/attestation/attestationprovider"
	automationaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationaccount"
	automationcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationcertificate"
	automationconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationconnection"
	automationconnectioncertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationconnectioncertificate"
	automationconnectionclassiccertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationconnectionclassiccertificate"
	automationconnectionserviceprincipal "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationconnectionserviceprincipal"
	automationcredential "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationcredential"
	automationdscnodeconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationdscnodeconfiguration"
	automationjobschedule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationjobschedule"
	automationmodule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationmodule"
	automationrunbook "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationrunbook"
	automationschedule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationschedule"
	automationvariablebool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationvariablebool"
	automationvariabledatetime "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationvariabledatetime"
	automationvariableint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationvariableint"
	automationvariablestring "github.com/crossplane-contrib/provider-tf-azure/internal/controller/automation/automationvariablestring"
	availabilityset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/availability/availabilityset"
	backupcontainerstorageaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/backup/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/crossplane-contrib/provider-tf-azure/internal/controller/backup/backuppolicyfileshare"
	backuppolicyvm "github.com/crossplane-contrib/provider-tf-azure/internal/controller/backup/backuppolicyvm"
	backupprotectedfileshare "github.com/crossplane-contrib/provider-tf-azure/internal/controller/backup/backupprotectedfileshare"
	backupprotectedvm "github.com/crossplane-contrib/provider-tf-azure/internal/controller/backup/backupprotectedvm"
	bastionhost "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bastion/bastionhost"
	batchaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/batch/batchaccount"
	batchapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/batch/batchapplication"
	batchcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/batch/batchcertificate"
	batchjob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/batch/batchjob"
	batchpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/batch/batchpool"
	blueprintassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/blueprint/blueprintassignment"
	botchannelalexa "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelalexa"
	botchanneldirectline "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchanneldirectline"
	botchanneldirectlinespeech "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchanneldirectlinespeech"
	botchannelemail "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelemail"
	botchannelfacebook "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelfacebook"
	botchannelline "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelline"
	botchannelmsteams "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelmsteams"
	botchannelslack "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelslack"
	botchannelsms "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelsms"
	botchannelsregistration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelsregistration"
	botchannelwebchat "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botchannelwebchat"
	botconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botconnection"
	botwebapp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/bot/botwebapp"
	cdnendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cdn/cdnendpoint"
	cdnendpointcustomdomain "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cdn/cdnendpointcustomdomain"
	cdnprofile "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cdn/cdnprofile"
	cognitiveaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cognitive/cognitiveaccount"
	cognitiveaccountcustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cognitive/cognitiveaccountcustomermanagedkey"
	communicationservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/communication/communicationservice"
	config "github.com/crossplane-contrib/provider-tf-azure/internal/controller/config"
	consumptionbudgetresourcegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/consumption/consumptionbudgetresourcegroup"
	consumptionbudgetsubscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/consumption/consumptionbudgetsubscription"
	containerregistry "github.com/crossplane-contrib/provider-tf-azure/internal/controller/container/containerregistry"
	containerregistryscopemap "github.com/crossplane-contrib/provider-tf-azure/internal/controller/container/containerregistryscopemap"
	containerregistrytoken "github.com/crossplane-contrib/provider-tf-azure/internal/controller/container/containerregistrytoken"
	containerregistrywebhook "github.com/crossplane-contrib/provider-tf-azure/internal/controller/container/containerregistrywebhook"
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
	costmanagementexportresourcegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/cost/costmanagementexportresourcegroup"
	customprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/custom/customprovider"
	dashboard "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dashboard/dashboard"
	datafactory "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactory"
	datafactorycustomdataset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorycustomdataset"
	datafactorydataflow "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydataflow"
	datafactorydatasetazureblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetazureblob"
	datafactorydatasetbinary "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetbinary"
	datafactorydatasetcosmosdbsqlapi "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetcosmosdbsqlapi"
	datafactorydatasetdelimitedtext "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetdelimitedtext"
	datafactorydatasethttp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasethttp"
	datafactorydatasetjson "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetjson"
	datafactorydatasetmysql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetmysql"
	datafactorydatasetparquet "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetparquet"
	datafactorydatasetpostgresql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetpostgresql"
	datafactorydatasetsnowflake "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetsnowflake"
	datafactorydatasetsqlservertable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorydatasetsqlservertable"
	datafactoryintegrationruntimeazure "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeazure"
	datafactoryintegrationruntimeazuressis "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeazuressis"
	datafactoryintegrationruntimemanaged "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimemanaged"
	datafactoryintegrationruntimeselfhosted "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeselfhosted"
	datafactorylinkedcustomservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedcustomservice"
	datafactorylinkedserviceazureblobstorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazureblobstorage"
	datafactorylinkedserviceazuredatabricks "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuredatabricks"
	datafactorylinkedserviceazurefilestorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazurefilestorage"
	datafactorylinkedserviceazurefunction "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazurefunction"
	datafactorylinkedserviceazuresearch "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuresearch"
	datafactorylinkedserviceazuresqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuresqldatabase"
	datafactorylinkedserviceazuretablestorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuretablestorage"
	datafactorylinkedservicecosmosdb "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicecosmosdb"
	datafactorylinkedservicedatalakestoragegen2 "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicedatalakestoragegen2"
	datafactorylinkedservicekeyvault "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicekeyvault"
	datafactorylinkedservicekusto "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicekusto"
	datafactorylinkedservicemysql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicemysql"
	datafactorylinkedserviceodata "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceodata"
	datafactorylinkedservicepostgresql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicepostgresql"
	datafactorylinkedservicesftp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicesftp"
	datafactorylinkedservicesnowflake "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicesnowflake"
	datafactorylinkedservicesqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicesqlserver"
	datafactorylinkedservicesynapse "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedservicesynapse"
	datafactorylinkedserviceweb "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorylinkedserviceweb"
	datafactorymanagedprivateendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorymanagedprivateendpoint"
	datafactorypipeline "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorypipeline"
	datafactorytriggerblobevent "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorytriggerblobevent"
	datafactorytriggercustomevent "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorytriggercustomevent"
	datafactorytriggerschedule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datafactorytriggerschedule"
	datalakeanalyticsfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datalakeanalyticsfirewallrule"
	datalakestore "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datalakestore"
	datalakestorefile "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datalakestorefile"
	datalakestorefirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datalakestorefirewallrule"
	datalakestorevirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datalakestorevirtualnetworkrule"
	dataprotectionbackupinstanceblobstorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackupinstanceblobstorage"
	dataprotectionbackupinstancedisk "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackupinstancedisk"
	dataprotectionbackupinstancepostgresql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackupinstancepostgresql"
	dataprotectionbackuppolicyblobstorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicyblobstorage"
	dataprotectionbackuppolicydisk "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicydisk"
	dataprotectionbackuppolicypostgresql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicypostgresql"
	dataprotectionbackupvault "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/dataprotectionbackupvault"
	datashare "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datashare"
	datashareaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datashareaccount"
	datasharedatasetblobstorage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datasharedatasetblobstorage"
	datasharedatasetdatalakegen1 "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datasharedatasetdatalakegen1"
	datasharedatasetdatalakegen2 "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datasharedatasetdatalakegen2"
	datasharedatasetkustocluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datasharedatasetkustocluster"
	datasharedatasetkustodatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/data/datasharedatasetkustodatabase"
	databasemigrationproject "github.com/crossplane-contrib/provider-tf-azure/internal/controller/database/databasemigrationproject"
	databasemigrationservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/database/databasemigrationservice"
	databoxedgedevice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/databox/databoxedgedevice"
	databoxedgeorder "github.com/crossplane-contrib/provider-tf-azure/internal/controller/databox/databoxedgeorder"
	databricksworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/databricks/databricksworkspace"
	databricksworkspacecustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/databricks/databricksworkspacecustomermanagedkey"
	dedicatedhardwaresecuritymodule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dedicated/dedicatedhardwaresecuritymodule"
	dedicatedhost "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dedicated/dedicatedhost"
	devtestglobalvmshutdownschedule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestglobalvmshutdownschedule"
	devtestlab "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestlab"
	devtestlinuxvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestlinuxvirtualmachine"
	devtestpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestpolicy"
	devtestschedule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestschedule"
	devtestvirtualnetwork "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestvirtualnetwork"
	devtestwindowsvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dev/devtestwindowsvirtualmachine"
	devspacecontroller "github.com/crossplane-contrib/provider-tf-azure/internal/controller/devspace/devspacecontroller"
	digitaltwinsendpointeventgrid "github.com/crossplane-contrib/provider-tf-azure/internal/controller/digital/digitaltwinsendpointeventgrid"
	digitaltwinsendpointeventhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/digital/digitaltwinsendpointeventhub"
	digitaltwinsendpointservicebus "github.com/crossplane-contrib/provider-tf-azure/internal/controller/digital/digitaltwinsendpointservicebus"
	digitaltwinsinstance "github.com/crossplane-contrib/provider-tf-azure/internal/controller/digital/digitaltwinsinstance"
	diskaccess "github.com/crossplane-contrib/provider-tf-azure/internal/controller/disk/diskaccess"
	diskencryptionset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/disk/diskencryptionset"
	dnsaaaarecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnsaaaarecord"
	dnsarecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnsarecord"
	dnscaarecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnscaarecord"
	dnscnamerecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnscnamerecord"
	dnsmxrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnsmxrecord"
	dnsnsrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnsnsrecord"
	dnsptrrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnsptrrecord"
	dnssrvrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnssrvrecord"
	dnstxtrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnstxtrecord"
	dnszone "github.com/crossplane-contrib/provider-tf-azure/internal/controller/dns/dnszone"
	eventgriddomain "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgriddomain"
	eventgriddomaintopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgriddomaintopic"
	eventgrideventsubscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgrideventsubscription"
	eventgridsystemtopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgridsystemtopic"
	eventgridsystemtopiceventsubscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgridsystemtopiceventsubscription"
	eventgridtopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventgrid/eventgridtopic"
	eventhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhub"
	eventhubauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubauthorizationrule"
	eventhubcluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubcluster"
	eventhubconsumergroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubconsumergroup"
	eventhubnamespace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubnamespace"
	eventhubnamespaceauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubnamespaceauthorizationrule"
	eventhubnamespacecustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubnamespacecustomermanagedkey"
	eventhubnamespacedisasterrecoveryconfig "github.com/crossplane-contrib/provider-tf-azure/internal/controller/eventhub/eventhubnamespacedisasterrecoveryconfig"
	expressroutecircuit "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressroutecircuit"
	expressroutecircuitauthorization "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressroutecircuitpeering"
	expressrouteconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressrouteconnection"
	expressroutegateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressroutegateway"
	expressrouteport "github.com/crossplane-contrib/provider-tf-azure/internal/controller/express/expressrouteport"
	firewall "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewall"
	firewallapplicationrulecollection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewallnetworkrulecollection"
	firewallpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/firewall/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/crossplane-contrib/provider-tf-azure/internal/controller/frontdoor/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/frontdoor/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/frontdoor/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/frontdoor/frontdoorrulesengine"
	functionapp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/function/functionapp"
	functionappslot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/function/functionappslot"
	hdinsighthadoopcluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsighthadoopcluster"
	hdinsighthbasecluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsighthbasecluster"
	hdinsightinteractivequerycluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightinteractivequerycluster"
	hdinsightkafkacluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightkafkacluster"
	hdinsightmlservicescluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightmlservicescluster"
	hdinsightrservercluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightrservercluster"
	hdinsightsparkcluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightsparkcluster"
	hdinsightstormcluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hdinsight/hdinsightstormcluster"
	healthbot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/healthbot/healthbot"
	healthcareservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/healthcare/healthcareservice"
	hpccache "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hpc/hpccache"
	hpccacheaccesspolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hpc/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hpc/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hpc/hpccacheblobtarget"
	hpccachenfstarget "github.com/crossplane-contrib/provider-tf-azure/internal/controller/hpc/hpccachenfstarget"
	image "github.com/crossplane-contrib/provider-tf-azure/internal/controller/image/image"
	integrationserviceenvironment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/integration/integrationserviceenvironment"
	iotsecuritydevicegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iotsecuritysolution"
	iottimeseriesinsightsaccesspolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsaccesspolicy"
	iottimeseriesinsightseventsourceiothub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iottimeseriesinsightseventsourceiothub"
	iottimeseriesinsightsgen2environment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsgen2environment"
	iottimeseriesinsightsreferencedataset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsreferencedataset"
	iottimeseriesinsightsstandardenvironment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsstandardenvironment"
	iotcentralapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iotcentral/iotcentralapplication"
	iothub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothub"
	iothubconsumergroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubconsumergroup"
	iothubdps "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubdps"
	iothubdpscertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubendpointstoragecontainer"
	iothubenrichment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubenrichment"
	iothubfallbackroute "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubfallbackroute"
	iothubroute "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubroute"
	iothubsharedaccesspolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/iothub/iothubsharedaccesspolicy"
	ipgroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/ip/ipgroup"
	keyvault "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvault"
	keyvaultaccesspolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultaccesspolicy"
	keyvaultcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultcertificate"
	keyvaultcertificateissuer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultcertificateissuer"
	keyvaultkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultkey"
	keyvaultmanagedhardwaresecuritymodule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultmanagedhardwaresecuritymodule"
	keyvaultmanagedstorageaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultmanagedstorageaccount"
	keyvaultmanagedstorageaccountsastokendefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultmanagedstorageaccountsastokendefinition"
	keyvaultsecret "github.com/crossplane-contrib/provider-tf-azure/internal/controller/key/keyvaultsecret"
	kubernetescluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/kubernetescluster"
	kubernetesclusternodepool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kubernetes/kubernetesclusternodepool"
	kustoattacheddatabaseconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoattacheddatabaseconfiguration"
	kustocluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustocluster"
	kustoclustercustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoclustercustomermanagedkey"
	kustoclusterprincipalassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoclusterprincipalassignment"
	kustodatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustodatabase"
	kustodatabaseprincipal "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustodatabaseprincipal"
	kustodatabaseprincipalassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustodatabaseprincipalassignment"
	kustoeventgriddataconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoeventgriddataconnection"
	kustoeventhubdataconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoeventhubdataconnection"
	kustoiothubdataconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/kusto/kustoiothubdataconnection"
	lbbackendaddresspool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbbackendaddresspool"
	lbbackendaddresspooladdress "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbbackendaddresspooladdress"
	lbnatpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbnatpool"
	lbnatrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbnatrule"
	lboutboundrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lboutboundrule"
	lbprobe "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbprobe"
	lbrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/lbrule"
	loadbalancer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lb/loadbalancer"
	lighthouseassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lighthouse/lighthouseassignment"
	lighthousedefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/lighthouse/lighthousedefinition"
	linuxvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/linux/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/linux/linuxvirtualmachinescaleset"
	localnetworkgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/local/localnetworkgateway"
	loganalyticscluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticscluster"
	loganalyticsclustercustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticsclustercustomermanagedkey"
	loganalyticsdataexportrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticslinkedstorageaccount"
	loganalyticssavedsearch "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticssavedsearch"
	loganalyticssolution "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticssolution"
	loganalyticsworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/log/loganalyticsworkspace"
	logicappactioncustom "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappactioncustom"
	logicappactionhttp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappactionhttp"
	logicappintegrationaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccount"
	logicappintegrationaccountagreement "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountagreement"
	logicappintegrationaccountassembly "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountassembly"
	logicappintegrationaccountbatchconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountbatchconfiguration"
	logicappintegrationaccountmap "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountmap"
	logicappintegrationaccountpartner "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountpartner"
	logicappintegrationaccountschema "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountschema"
	logicappintegrationaccountsession "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappintegrationaccountsession"
	logicapptriggercustom "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicapptriggercustom"
	logicapptriggerhttprequest "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicapptriggerhttprequest"
	logicapptriggerrecurrence "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicapptriggerrecurrence"
	logicappworkflow "github.com/crossplane-contrib/provider-tf-azure/internal/controller/logic/logicappworkflow"
	machinelearningcomputecluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/machine/machinelearningcomputecluster"
	machinelearningcomputeinstance "github.com/crossplane-contrib/provider-tf-azure/internal/controller/machine/machinelearningcomputeinstance"
	machinelearningsynapsespark "github.com/crossplane-contrib/provider-tf-azure/internal/controller/machine/machinelearningsynapsespark"
	machinelearningworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/machine/machinelearningworkspace"
	maintenanceassignmentdedicatedhost "github.com/crossplane-contrib/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceassignmentvirtualmachinescaleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachinescaleset"
	maintenanceconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/maintenance/maintenanceconfiguration"
	managedapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/managed/managedapplication"
	managedapplicationdefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/managed/managedapplicationdefinition"
	manageddisk "github.com/crossplane-contrib/provider-tf-azure/internal/controller/managed/manageddisk"
	managementgroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/management/managementgroup"
	managementgrouppolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/management/managementgrouppolicyassignment"
	managementgroupsubscriptionassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/management/managementgroupsubscriptionassociation"
	managementgrouptemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/management/managementgrouptemplatedeployment"
	managementlock "github.com/crossplane-contrib/provider-tf-azure/internal/controller/management/managementlock"
	mapsaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/maps/mapsaccount"
	mariadbconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mariadb/mariadbconfiguration"
	mariadbdatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mariadb/mariadbdatabase"
	mariadbfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mariadb/mariadbfirewallrule"
	mariadbserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mariadb/mariadbserver"
	mariadbvirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mariadb/mariadbvirtualnetworkrule"
	marketplaceagreement "github.com/crossplane-contrib/provider-tf-azure/internal/controller/marketplace/marketplaceagreement"
	mediaasset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediaasset"
	mediaassetfilter "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediaassetfilter"
	mediacontentkeypolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediacontentkeypolicy"
	mediajob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediajob"
	medialiveevent "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/medialiveevent"
	medialiveeventoutput "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/medialiveeventoutput"
	mediaservicesaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediaservicesaccount"
	mediastreamingendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediastreamingendpoint"
	mediastreaminglocator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediastreaminglocator"
	mediastreamingpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediastreamingpolicy"
	mediatransform "github.com/crossplane-contrib/provider-tf-azure/internal/controller/media/mediatransform"
	monitoraaddiagnosticsetting "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitoraaddiagnosticsetting"
	monitoractiongroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitoractiongroup"
	monitoractionruleactiongroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitoractionrulesuppression"
	monitoractivitylogalert "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitoractivitylogalert"
	monitorautoscalesetting "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitorautoscalesetting"
	monitordiagnosticsetting "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitordiagnosticsetting"
	monitormetricalert "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitormetricalert"
	monitorscheduledqueryrulesalert "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitorscheduledqueryrulesalert"
	monitorscheduledqueryruleslog "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitorscheduledqueryruleslog"
	monitorsmartdetectoralertrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/monitor/monitorsmartdetectoralertrule"
	mssqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlelasticpool"
	mssqlfailovergroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlfirewallrule"
	mssqljobagent "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqljobagent"
	mssqljobcredential "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqljobcredential"
	mssqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlserver"
	mssqlserversecurityalertpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlservervulnerabilityassessment"
	mssqlvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlvirtualmachine"
	mssqlvirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mssql/mssqlvirtualnetworkrule"
	mysqlactivedirectoryadministrator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlactivedirectoryadministrator"
	mysqlconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlconfiguration"
	mysqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqldatabase"
	mysqlfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlfirewallrule"
	mysqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlserver"
	mysqlserverkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlserverkey"
	mysqlvirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/mysql/mysqlvirtualnetworkrule"
	natgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/nat/natgateway"
	natgatewaypublicipassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/nat/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/nat/natgatewaypublicipprefixassociation"
	netappaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/netapp/netappaccount"
	netapppool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/netapp/netapppool"
	netappsnapshot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/netapp/netappsnapshot"
	netappvolume "github.com/crossplane-contrib/provider-tf-azure/internal/controller/netapp/netappvolume"
	networkconnectionmonitor "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkconnectionmonitor"
	networkddosprotectionplan "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkddosprotectionplan"
	networkinterface "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkinterfacesecuritygroupassociation"
	networkpacketcapture "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkpacketcapture"
	networkprofile "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkprofile"
	networksecuritygroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networksecuritygroup"
	networksecurityrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networksecurityrule"
	networkwatcher "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkwatcher"
	networkwatcherflowlog "github.com/crossplane-contrib/provider-tf-azure/internal/controller/network/networkwatcherflowlog"
	notificationhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/notification/notificationhub"
	notificationhubauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/notification/notificationhubauthorizationrule"
	notificationhubnamespace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/notification/notificationhubnamespace"
	orchestratedvirtualmachinescaleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/orchestrated/orchestratedvirtualmachinescaleset"
	packetcapture "github.com/crossplane-contrib/provider-tf-azure/internal/controller/packet/packetcapture"
	pointtositevpngateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/point/pointtositevpngateway"
	policyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/policy/policyassignment"
	policydefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/policy/policydefinition"
	policyremediation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/policy/policyremediation"
	policysetdefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/policy/policysetdefinition"
	policyvirtualmachineconfigurationassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/policy/policyvirtualmachineconfigurationassignment"
	portaltenantconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/portal/portaltenantconfiguration"
	postgresqlactivedirectoryadministrator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlactivedirectoryadministrator"
	postgresqlconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlconfiguration"
	postgresqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqldatabase"
	postgresqlfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlfirewallrule"
	postgresqlflexibleserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserver"
	postgresqlflexibleserverconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverconfiguration"
	postgresqlflexibleserverdatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverdatabase"
	postgresqlflexibleserverfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverfirewallrule"
	postgresqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlserver"
	postgresqlvirtualnetworkrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/postgresql/postgresqlvirtualnetworkrule"
	powerbiembedded "github.com/crossplane-contrib/provider-tf-azure/internal/controller/powerbi/powerbiembedded"
	privatednsaaaarecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednsaaaarecord"
	privatednsarecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednsarecord"
	privatednscnamerecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednscnamerecord"
	privatednsmxrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednsmxrecord"
	privatednsptrrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednsptrrecord"
	privatednssrvrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednssrvrecord"
	privatednstxtrecord "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednstxtrecord"
	privatednszone "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednszone"
	privatednszonevirtualnetworklink "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatednszonevirtualnetworklink"
	privateendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privateendpoint"
	privatelinkservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/private/privatelinkservice"
	proximityplacementgroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/proximity/proximityplacementgroup"
	publicip "github.com/crossplane-contrib/provider-tf-azure/internal/controller/public/publicip"
	publicipprefix "github.com/crossplane-contrib/provider-tf-azure/internal/controller/public/publicipprefix"
	purviewaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/purview/purviewaccount"
	recoveryservicesvault "github.com/crossplane-contrib/provider-tf-azure/internal/controller/recovery/recoveryservicesvault"
	rediscache "github.com/crossplane-contrib/provider-tf-azure/internal/controller/redis/rediscache"
	redisenterprisecluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/redis/redisenterprisecluster"
	redisenterprisedatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/redis/redisenterprisedatabase"
	redisfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/redis/redisfirewallrule"
	redislinkedserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/redis/redislinkedserver"
	relayhybridconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/relay/relayhybridconnection"
	relayhybridconnectionauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/relay/relayhybridconnectionauthorizationrule"
	relaynamespace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/relay/relaynamespace"
	relaynamespaceauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/relay/relaynamespaceauthorizationrule"
	resourcegroup "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegroup"
	resourcegrouppolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegrouppolicyassignment"
	resourcegrouptemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcegrouptemplatedeployment"
	resourcepolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourcepolicyassignment"
	resourceproviderregistration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/resource/resourceproviderregistration"
	roleassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/role/roleassignment"
	roledefinition "github.com/crossplane-contrib/provider-tf-azure/internal/controller/role/roledefinition"
	route "github.com/crossplane-contrib/provider-tf-azure/internal/controller/route/route"
	routefilter "github.com/crossplane-contrib/provider-tf-azure/internal/controller/route/routefilter"
	routetable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/route/routetable"
	searchservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/search/searchservice"
	securitycenterassessment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterassessment"
	securitycenterassessmentmetadata "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterassessmentmetadata"
	securitycenterassessmentpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterservervulnerabilityassessment"
	securitycentersetting "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/security/securitycenterworkspace"
	sentinelalertrulefusion "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentinelalertrulemssecurityincident"
	sentinelalertrulescheduled "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentinelalertrulescheduled"
	sentineldataconnectorawscloudtrail "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorawscloudtrail"
	sentineldataconnectorazureactivedirectory "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazureactivedirectory"
	sentineldataconnectorazureadvancedthreatprotection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazureadvancedthreatprotection"
	sentineldataconnectorazuresecuritycenter "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazuresecuritycenter"
	sentineldataconnectormicrosoftcloudappsecurity "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectormicrosoftcloudappsecurity"
	sentineldataconnectoroffice365 "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectoroffice365"
	sentineldataconnectorthreatintelligence "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorthreatintelligence"
	servicefabriccluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/service/servicefabriccluster"
	servicefabricmeshapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/service/servicefabricmeshapplication"
	servicefabricmeshlocalnetwork "github.com/crossplane-contrib/provider-tf-azure/internal/controller/service/servicefabricmeshlocalnetwork"
	servicefabricmeshsecret "github.com/crossplane-contrib/provider-tf-azure/internal/controller/service/servicefabricmeshsecret"
	servicefabricmeshsecretvalue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/service/servicefabricmeshsecretvalue"
	servicebusnamespace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusnamespace"
	servicebusnamespaceauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusnamespaceauthorizationrule"
	servicebusnamespacedisasterrecoveryconfig "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusnamespacedisasterrecoveryconfig"
	servicebusnamespacenetworkruleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusnamespacenetworkruleset"
	servicebusqueue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusqueue"
	servicebusqueueauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebusqueueauthorizationrule"
	servicebussubscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebussubscription"
	servicebussubscriptionrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebussubscriptionrule"
	servicebustopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebustopic"
	servicebustopicauthorizationrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/servicebus/servicebustopicauthorizationrule"
	sharedimage "github.com/crossplane-contrib/provider-tf-azure/internal/controller/shared/sharedimage"
	sharedimagegallery "github.com/crossplane-contrib/provider-tf-azure/internal/controller/shared/sharedimagegallery"
	sharedimageversion "github.com/crossplane-contrib/provider-tf-azure/internal/controller/shared/sharedimageversion"
	signalrservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/signalr/signalrservice"
	signalrservicenetworkacl "github.com/crossplane-contrib/provider-tf-azure/internal/controller/signalr/signalrservicenetworkacl"
	siterecoveryfabric "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoveryprotectioncontainermapping"
	siterecoveryreplicatedvm "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoveryreplicatedvm"
	siterecoveryreplicationpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/site/siterecoveryreplicationpolicy"
	snapshot "github.com/crossplane-contrib/provider-tf-azure/internal/controller/snapshot/snapshot"
	spatialanchorsaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spatial/spatialanchorsaccount"
	springcloudactivedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudactivedeployment"
	springcloudapp "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudapp"
	springcloudappcosmosdbassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudappredisassociation"
	springcloudcertificate "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudcertificate"
	springcloudcustomdomain "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudcustomdomain"
	springcloudjavadeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudjavadeployment"
	springcloudservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/spring/springcloudservice"
	sqlactivedirectoryadministrator "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlactivedirectoryadministrator"
	sqldatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqldatabase"
	sqlelasticpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlelasticpool"
	sqlfirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlfirewallrule"
	sqlmanageddatabase "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlmanageddatabase"
	sqlmanagedinstance "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlmanagedinstance"
	sqlserver "github.com/crossplane-contrib/provider-tf-azure/internal/controller/sql/sqlserver"
	sshpublickey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/ssh/sshpublickey"
	stackhcicluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stack/stackhcicluster"
	staticsite "github.com/crossplane-contrib/provider-tf-azure/internal/controller/static/staticsite"
	storageaccount "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageaccount"
	storageaccountcustomermanagedkey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageaccountcustomermanagedkey"
	storageaccountnetworkrules "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageaccountnetworkrules"
	storageblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageblob"
	storageblobinventorypolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageblobinventorypolicy"
	storagecontainer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagecontainer"
	storagedatalakegen2filesystem "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagedatalakegen2filesystem"
	storagedatalakegen2path "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagedatalakegen2path"
	storageencryptionscope "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageencryptionscope"
	storagemanagementpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagemanagementpolicy"
	storageobjectreplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageobjectreplication"
	storagequeue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagequeue"
	storageshare "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storageshare"
	storagesharedirectory "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagesharedirectory"
	storagesync "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagesync"
	storagesynccloudendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagesynccloudendpoint"
	storagetable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagetable"
	storagetableentity "github.com/crossplane-contrib/provider-tf-azure/internal/controller/storage/storagetableentity"
	streamanalyticsfunctionjavascriptudf "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsfunctionjavascriptudf"
	streamanalyticsjob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsjob"
	streamanalyticsoutputblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsoutputblob"
	streamanalyticsoutputeventhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsoutputeventhub"
	streamanalyticsoutputmssql "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsoutputmssql"
	streamanalyticsoutputservicebusqueue "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsoutputservicebusqueue"
	streamanalyticsoutputservicebustopic "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsoutputservicebustopic"
	streamanalyticsreferenceinputblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsreferenceinputblob"
	streamanalyticsstreaminputblob "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputblob"
	streamanalyticsstreaminputeventhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputeventhub"
	streamanalyticsstreaminputiothub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputiothub"
	subnet "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnet"
	subnetnatgatewayassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subnet/subnetserviceendpointstoragepolicy"
	subscription "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subscription/subscription"
	subscriptionpolicyassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subscription/subscriptionpolicyassignment"
	subscriptiontemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/subscription/subscriptiontemplatedeployment"
	synapsefirewallrule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsefirewallrule"
	synapseintegrationruntimeazure "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseintegrationruntimeazure"
	synapseintegrationruntimeselfhosted "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseintegrationruntimeselfhosted"
	synapselinkedservice "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapselinkedservice"
	synapsemanagedprivateendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsemanagedprivateendpoint"
	synapseprivatelinkhub "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseprivatelinkhub"
	synapseroleassignment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseroleassignment"
	synapsesparkpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsesparkpool"
	synapsesqlpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsesqlpool"
	synapsesqlpoolextendedauditingpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsesqlpoolextendedauditingpolicy"
	synapsesqlpoolsecurityalertpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsesqlpoolsecurityalertpolicy"
	synapsesqlpoolvulnerabilityassessment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapsesqlpoolvulnerabilityassessment"
	synapseworkspace "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseworkspace"
	synapseworkspaceextendedauditingpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseworkspaceextendedauditingpolicy"
	synapseworkspacekey "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseworkspacekey"
	synapseworkspacesecurityalertpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseworkspacesecurityalertpolicy"
	synapseworkspacevulnerabilityassessment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/synapse/synapseworkspacevulnerabilityassessment"
	templatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/template/templatedeployment"
	tenanttemplatedeployment "github.com/crossplane-contrib/provider-tf-azure/internal/controller/tenant/tenanttemplatedeployment"
	trafficmanagerendpoint "github.com/crossplane-contrib/provider-tf-azure/internal/controller/traffic/trafficmanagerendpoint"
	trafficmanagerprofile "github.com/crossplane-contrib/provider-tf-azure/internal/controller/traffic/trafficmanagerprofile"
	userassignedidentity "github.com/crossplane-contrib/provider-tf-azure/internal/controller/user/userassignedidentity"
	videoanalyzer "github.com/crossplane-contrib/provider-tf-azure/internal/controller/video/videoanalyzer"
	videoanalyzeredgemodule "github.com/crossplane-contrib/provider-tf-azure/internal/controller/video/videoanalyzeredgemodule"
	virtualdesktopapplication "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktopapplication"
	virtualdesktophostpool "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualdesktophostpool"
	virtualhubconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubconnection"
	virtualhubip "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubip"
	virtualhubroutetable "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubroutetable"
	virtualhubsecuritypartnerprovider "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualhubsecuritypartnerprovider"
	virtualnetwork "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetwork"
	virtualnetworkgateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualnetworkpeering"
	virtualwan "github.com/crossplane-contrib/provider-tf-azure/internal/controller/virtual/virtualwan"
	vmwarecluster "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vmware/vmwarecluster"
	vmwareexpressrouteauthorization "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vmware/vmwareexpressrouteauthorization"
	vmwareprivatecloud "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vmware/vmwareprivatecloud"
	vpngateway "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vpn/vpngateway"
	vpngatewayconnection "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vpn/vpngatewayconnection"
	vpnserverconfiguration "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vpn/vpnserverconfiguration"
	vpnsite "github.com/crossplane-contrib/provider-tf-azure/internal/controller/vpn/vpnsite"
	webapplicationfirewallpolicy "github.com/crossplane-contrib/provider-tf-azure/internal/controller/web/webapplicationfirewallpolicy"
	windowsvirtualmachine "github.com/crossplane-contrib/provider-tf-azure/internal/controller/windows/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/crossplane-contrib/provider-tf-azure/internal/controller/windows/windowsvirtualmachinescaleset"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, int) error{
		activedirectorydomainservice.Setup,
		activedirectorydomainservicereplicaset.Setup,
		advancedthreatprotection.Setup,
		analysisservicesserver.Setup,
		apimanagement.Setup,
		apimanagementapi.Setup,
		apimanagementapidiagnostic.Setup,
		apimanagementapioperation.Setup,
		apimanagementapioperationpolicy.Setup,
		apimanagementapioperationtag.Setup,
		apimanagementapipolicy.Setup,
		apimanagementapirelease.Setup,
		apimanagementapischema.Setup,
		apimanagementapiversionset.Setup,
		apimanagementauthorizationserver.Setup,
		apimanagementbackend.Setup,
		apimanagementcertificate.Setup,
		apimanagementcustomdomain.Setup,
		apimanagementdiagnostic.Setup,
		apimanagementemailtemplate.Setup,
		apimanagementgateway.Setup,
		apimanagementgatewayapi.Setup,
		apimanagementgroupuser.Setup,
		apimanagementidentityprovideraad.Setup,
		apimanagementidentityprovideraadb2c.Setup,
		apimanagementidentityproviderfacebook.Setup,
		apimanagementidentityprovidergoogle.Setup,
		apimanagementidentityprovidermicrosoft.Setup,
		apimanagementidentityprovidertwitter.Setup,
		apimanagementlogger.Setup,
		apimanagementnamedvalue.Setup,
		apimanagementnotificationrecipientemail.Setup,
		apimanagementopenidconnectprovider.Setup,
		apimanagementpolicy.Setup,
		apimanagementproduct.Setup,
		apimanagementproductapi.Setup,
		apimanagementproductpolicy.Setup,
		apimanagementproperty.Setup,
		apimanagementrediscache.Setup,
		apimanagementsubscription.Setup,
		apimanagementtag.Setup,
		apimanagementuser.Setup,
		appconfiguration.Setup,
		appconfigurationkey.Setup,
		applicationgateway.Setup,
		applicationinsights.Setup,
		applicationinsightsanalyticsitem.Setup,
		applicationinsightsapikey.Setup,
		applicationinsightssmartdetectionrule.Setup,
		applicationinsightswebtest.Setup,
		applicationsecuritygroup.Setup,
		appservice.Setup,
		appserviceactiveslot.Setup,
		appservicecertificate.Setup,
		appservicecertificatebinding.Setup,
		appservicecertificateorder.Setup,
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
		attestationprovider.Setup,
		automationaccount.Setup,
		automationcertificate.Setup,
		automationconnection.Setup,
		automationconnectioncertificate.Setup,
		automationconnectionclassiccertificate.Setup,
		automationconnectionserviceprincipal.Setup,
		automationcredential.Setup,
		automationdscnodeconfiguration.Setup,
		automationjobschedule.Setup,
		automationmodule.Setup,
		automationrunbook.Setup,
		automationschedule.Setup,
		automationvariablebool.Setup,
		automationvariabledatetime.Setup,
		automationvariableint.Setup,
		automationvariablestring.Setup,
		availabilityset.Setup,
		backupcontainerstorageaccount.Setup,
		backuppolicyfileshare.Setup,
		backuppolicyvm.Setup,
		backupprotectedfileshare.Setup,
		backupprotectedvm.Setup,
		bastionhost.Setup,
		batchaccount.Setup,
		batchapplication.Setup,
		batchcertificate.Setup,
		batchjob.Setup,
		batchpool.Setup,
		blueprintassignment.Setup,
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
		cdnendpoint.Setup,
		cdnendpointcustomdomain.Setup,
		cdnprofile.Setup,
		cognitiveaccount.Setup,
		cognitiveaccountcustomermanagedkey.Setup,
		communicationservice.Setup,
		config.Setup,
		consumptionbudgetresourcegroup.Setup,
		consumptionbudgetsubscription.Setup,
		containerregistry.Setup,
		containerregistryscopemap.Setup,
		containerregistrytoken.Setup,
		containerregistrywebhook.Setup,
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
		costmanagementexportresourcegroup.Setup,
		customprovider.Setup,
		dashboard.Setup,
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
		databoxedgedevice.Setup,
		databoxedgeorder.Setup,
		databricksworkspace.Setup,
		databricksworkspacecustomermanagedkey.Setup,
		datafactory.Setup,
		datafactorycustomdataset.Setup,
		datafactorydataflow.Setup,
		datafactorydatasetazureblob.Setup,
		datafactorydatasetbinary.Setup,
		datafactorydatasetcosmosdbsqlapi.Setup,
		datafactorydatasetdelimitedtext.Setup,
		datafactorydatasethttp.Setup,
		datafactorydatasetjson.Setup,
		datafactorydatasetmysql.Setup,
		datafactorydatasetparquet.Setup,
		datafactorydatasetpostgresql.Setup,
		datafactorydatasetsnowflake.Setup,
		datafactorydatasetsqlservertable.Setup,
		datafactoryintegrationruntimeazure.Setup,
		datafactoryintegrationruntimeazuressis.Setup,
		datafactoryintegrationruntimemanaged.Setup,
		datafactoryintegrationruntimeselfhosted.Setup,
		datafactorylinkedcustomservice.Setup,
		datafactorylinkedserviceazureblobstorage.Setup,
		datafactorylinkedserviceazuredatabricks.Setup,
		datafactorylinkedserviceazurefilestorage.Setup,
		datafactorylinkedserviceazurefunction.Setup,
		datafactorylinkedserviceazuresearch.Setup,
		datafactorylinkedserviceazuresqldatabase.Setup,
		datafactorylinkedserviceazuretablestorage.Setup,
		datafactorylinkedservicecosmosdb.Setup,
		datafactorylinkedservicedatalakestoragegen2.Setup,
		datafactorylinkedservicekeyvault.Setup,
		datafactorylinkedservicekusto.Setup,
		datafactorylinkedservicemysql.Setup,
		datafactorylinkedserviceodata.Setup,
		datafactorylinkedservicepostgresql.Setup,
		datafactorylinkedservicesftp.Setup,
		datafactorylinkedservicesnowflake.Setup,
		datafactorylinkedservicesqlserver.Setup,
		datafactorylinkedservicesynapse.Setup,
		datafactorylinkedserviceweb.Setup,
		datafactorymanagedprivateendpoint.Setup,
		datafactorypipeline.Setup,
		datafactorytriggerblobevent.Setup,
		datafactorytriggercustomevent.Setup,
		datafactorytriggerschedule.Setup,
		datalakeanalyticsfirewallrule.Setup,
		datalakestore.Setup,
		datalakestorefile.Setup,
		datalakestorefirewallrule.Setup,
		datalakestorevirtualnetworkrule.Setup,
		dataprotectionbackupinstanceblobstorage.Setup,
		dataprotectionbackupinstancedisk.Setup,
		dataprotectionbackupinstancepostgresql.Setup,
		dataprotectionbackuppolicyblobstorage.Setup,
		dataprotectionbackuppolicydisk.Setup,
		dataprotectionbackuppolicypostgresql.Setup,
		dataprotectionbackupvault.Setup,
		datashare.Setup,
		datashareaccount.Setup,
		datasharedatasetblobstorage.Setup,
		datasharedatasetdatalakegen1.Setup,
		datasharedatasetdatalakegen2.Setup,
		datasharedatasetkustocluster.Setup,
		datasharedatasetkustodatabase.Setup,
		dedicatedhardwaresecuritymodule.Setup,
		dedicatedhost.Setup,
		devspacecontroller.Setup,
		devtestglobalvmshutdownschedule.Setup,
		devtestlab.Setup,
		devtestlinuxvirtualmachine.Setup,
		devtestpolicy.Setup,
		devtestschedule.Setup,
		devtestvirtualnetwork.Setup,
		devtestwindowsvirtualmachine.Setup,
		digitaltwinsendpointeventgrid.Setup,
		digitaltwinsendpointeventhub.Setup,
		digitaltwinsendpointservicebus.Setup,
		digitaltwinsinstance.Setup,
		diskaccess.Setup,
		diskencryptionset.Setup,
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
		eventgriddomain.Setup,
		eventgriddomaintopic.Setup,
		eventgrideventsubscription.Setup,
		eventgridsystemtopic.Setup,
		eventgridsystemtopiceventsubscription.Setup,
		eventgridtopic.Setup,
		eventhub.Setup,
		eventhubauthorizationrule.Setup,
		eventhubcluster.Setup,
		eventhubconsumergroup.Setup,
		eventhubnamespace.Setup,
		eventhubnamespaceauthorizationrule.Setup,
		eventhubnamespacecustomermanagedkey.Setup,
		eventhubnamespacedisasterrecoveryconfig.Setup,
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
		functionapp.Setup,
		functionappslot.Setup,
		hdinsighthadoopcluster.Setup,
		hdinsighthbasecluster.Setup,
		hdinsightinteractivequerycluster.Setup,
		hdinsightkafkacluster.Setup,
		hdinsightmlservicescluster.Setup,
		hdinsightrservercluster.Setup,
		hdinsightsparkcluster.Setup,
		hdinsightstormcluster.Setup,
		healthbot.Setup,
		healthcareservice.Setup,
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
		image.Setup,
		integrationserviceenvironment.Setup,
		iotcentralapplication.Setup,
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
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		iottimeseriesinsightsaccesspolicy.Setup,
		iottimeseriesinsightseventsourceiothub.Setup,
		iottimeseriesinsightsgen2environment.Setup,
		iottimeseriesinsightsreferencedataset.Setup,
		iottimeseriesinsightsstandardenvironment.Setup,
		ipgroup.Setup,
		keyvault.Setup,
		keyvaultaccesspolicy.Setup,
		keyvaultcertificate.Setup,
		keyvaultcertificateissuer.Setup,
		keyvaultkey.Setup,
		keyvaultmanagedhardwaresecuritymodule.Setup,
		keyvaultmanagedstorageaccount.Setup,
		keyvaultmanagedstorageaccountsastokendefinition.Setup,
		keyvaultsecret.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		kustoattacheddatabaseconfiguration.Setup,
		kustocluster.Setup,
		kustoclustercustomermanagedkey.Setup,
		kustoclusterprincipalassignment.Setup,
		kustodatabase.Setup,
		kustodatabaseprincipal.Setup,
		kustodatabaseprincipalassignment.Setup,
		kustoeventgriddataconnection.Setup,
		kustoeventhubdataconnection.Setup,
		kustoiothubdataconnection.Setup,
		lbbackendaddresspool.Setup,
		lbbackendaddresspooladdress.Setup,
		lbnatpool.Setup,
		lbnatrule.Setup,
		lboutboundrule.Setup,
		lbprobe.Setup,
		lbrule.Setup,
		lighthouseassignment.Setup,
		lighthousedefinition.Setup,
		linuxvirtualmachine.Setup,
		linuxvirtualmachinescaleset.Setup,
		loadbalancer.Setup,
		localnetworkgateway.Setup,
		loganalyticscluster.Setup,
		loganalyticsclustercustomermanagedkey.Setup,
		loganalyticsdataexportrule.Setup,
		loganalyticsdatasourcewindowsevent.Setup,
		loganalyticsdatasourcewindowsperformancecounter.Setup,
		loganalyticslinkedservice.Setup,
		loganalyticslinkedstorageaccount.Setup,
		loganalyticssavedsearch.Setup,
		loganalyticssolution.Setup,
		loganalyticsworkspace.Setup,
		logicappactioncustom.Setup,
		logicappactionhttp.Setup,
		logicappintegrationaccount.Setup,
		logicappintegrationaccountagreement.Setup,
		logicappintegrationaccountassembly.Setup,
		logicappintegrationaccountbatchconfiguration.Setup,
		logicappintegrationaccountmap.Setup,
		logicappintegrationaccountpartner.Setup,
		logicappintegrationaccountschema.Setup,
		logicappintegrationaccountsession.Setup,
		logicapptriggercustom.Setup,
		logicapptriggerhttprequest.Setup,
		logicapptriggerrecurrence.Setup,
		logicappworkflow.Setup,
		machinelearningcomputecluster.Setup,
		machinelearningcomputeinstance.Setup,
		machinelearningsynapsespark.Setup,
		machinelearningworkspace.Setup,
		maintenanceassignmentdedicatedhost.Setup,
		maintenanceassignmentvirtualmachine.Setup,
		maintenanceassignmentvirtualmachinescaleset.Setup,
		maintenanceconfiguration.Setup,
		managedapplication.Setup,
		managedapplicationdefinition.Setup,
		manageddisk.Setup,
		managementgroup.Setup,
		managementgrouppolicyassignment.Setup,
		managementgroupsubscriptionassociation.Setup,
		managementgrouptemplatedeployment.Setup,
		managementlock.Setup,
		mapsaccount.Setup,
		mariadbconfiguration.Setup,
		mariadbdatabase.Setup,
		mariadbfirewallrule.Setup,
		mariadbserver.Setup,
		mariadbvirtualnetworkrule.Setup,
		marketplaceagreement.Setup,
		mediaasset.Setup,
		mediaassetfilter.Setup,
		mediacontentkeypolicy.Setup,
		mediajob.Setup,
		medialiveevent.Setup,
		medialiveeventoutput.Setup,
		mediaservicesaccount.Setup,
		mediastreamingendpoint.Setup,
		mediastreaminglocator.Setup,
		mediastreamingpolicy.Setup,
		mediatransform.Setup,
		monitoraaddiagnosticsetting.Setup,
		monitoractiongroup.Setup,
		monitoractionruleactiongroup.Setup,
		monitoractionrulesuppression.Setup,
		monitoractivitylogalert.Setup,
		monitorautoscalesetting.Setup,
		monitordiagnosticsetting.Setup,
		monitormetricalert.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryruleslog.Setup,
		monitorsmartdetectoralertrule.Setup,
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
		mssqlvirtualmachine.Setup,
		mssqlvirtualnetworkrule.Setup,
		mysqlactivedirectoryadministrator.Setup,
		mysqlconfiguration.Setup,
		mysqldatabase.Setup,
		mysqlfirewallrule.Setup,
		mysqlserver.Setup,
		mysqlserverkey.Setup,
		mysqlvirtualnetworkrule.Setup,
		natgateway.Setup,
		natgatewaypublicipassociation.Setup,
		natgatewaypublicipprefixassociation.Setup,
		netappaccount.Setup,
		netapppool.Setup,
		netappsnapshot.Setup,
		netappvolume.Setup,
		networkconnectionmonitor.Setup,
		networkddosprotectionplan.Setup,
		networkinterface.Setup,
		networkinterfaceapplicationsecuritygroupassociation.Setup,
		networkinterfacebackendaddresspoolassociation.Setup,
		networkinterfacenatruleassociation.Setup,
		networkinterfacesecuritygroupassociation.Setup,
		networkpacketcapture.Setup,
		networkprofile.Setup,
		networksecuritygroup.Setup,
		networksecurityrule.Setup,
		networkwatcher.Setup,
		networkwatcherflowlog.Setup,
		notificationhub.Setup,
		notificationhubauthorizationrule.Setup,
		notificationhubnamespace.Setup,
		orchestratedvirtualmachinescaleset.Setup,
		packetcapture.Setup,
		pointtositevpngateway.Setup,
		policyassignment.Setup,
		policydefinition.Setup,
		policyremediation.Setup,
		policysetdefinition.Setup,
		policyvirtualmachineconfigurationassignment.Setup,
		portaltenantconfiguration.Setup,
		postgresqlactivedirectoryadministrator.Setup,
		postgresqlconfiguration.Setup,
		postgresqldatabase.Setup,
		postgresqlfirewallrule.Setup,
		postgresqlflexibleserver.Setup,
		postgresqlflexibleserverconfiguration.Setup,
		postgresqlflexibleserverdatabase.Setup,
		postgresqlflexibleserverfirewallrule.Setup,
		postgresqlserver.Setup,
		postgresqlvirtualnetworkrule.Setup,
		powerbiembedded.Setup,
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
		proximityplacementgroup.Setup,
		publicip.Setup,
		publicipprefix.Setup,
		purviewaccount.Setup,
		recoveryservicesvault.Setup,
		rediscache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		relayhybridconnection.Setup,
		relayhybridconnectionauthorizationrule.Setup,
		relaynamespace.Setup,
		relaynamespaceauthorizationrule.Setup,
		resourcegroup.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcegrouptemplatedeployment.Setup,
		resourcepolicyassignment.Setup,
		resourceproviderregistration.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		route.Setup,
		routefilter.Setup,
		routetable.Setup,
		searchservice.Setup,
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
		servicebusnamespace.Setup,
		servicebusnamespaceauthorizationrule.Setup,
		servicebusnamespacedisasterrecoveryconfig.Setup,
		servicebusnamespacenetworkruleset.Setup,
		servicebusqueue.Setup,
		servicebusqueueauthorizationrule.Setup,
		servicebussubscription.Setup,
		servicebussubscriptionrule.Setup,
		servicebustopic.Setup,
		servicebustopicauthorizationrule.Setup,
		servicefabriccluster.Setup,
		servicefabricmeshapplication.Setup,
		servicefabricmeshlocalnetwork.Setup,
		servicefabricmeshsecret.Setup,
		servicefabricmeshsecretvalue.Setup,
		sharedimage.Setup,
		sharedimagegallery.Setup,
		sharedimageversion.Setup,
		signalrservice.Setup,
		signalrservicenetworkacl.Setup,
		siterecoveryfabric.Setup,
		siterecoverynetworkmapping.Setup,
		siterecoveryprotectioncontainer.Setup,
		siterecoveryprotectioncontainermapping.Setup,
		siterecoveryreplicatedvm.Setup,
		siterecoveryreplicationpolicy.Setup,
		snapshot.Setup,
		spatialanchorsaccount.Setup,
		springcloudactivedeployment.Setup,
		springcloudapp.Setup,
		springcloudappcosmosdbassociation.Setup,
		springcloudappmysqlassociation.Setup,
		springcloudappredisassociation.Setup,
		springcloudcertificate.Setup,
		springcloudcustomdomain.Setup,
		springcloudjavadeployment.Setup,
		springcloudservice.Setup,
		sqlactivedirectoryadministrator.Setup,
		sqldatabase.Setup,
		sqlelasticpool.Setup,
		sqlfirewallrule.Setup,
		sqlmanageddatabase.Setup,
		sqlmanagedinstance.Setup,
		sqlserver.Setup,
		sshpublickey.Setup,
		stackhcicluster.Setup,
		staticsite.Setup,
		storageaccount.Setup,
		storageaccountcustomermanagedkey.Setup,
		storageaccountnetworkrules.Setup,
		storageblob.Setup,
		storageblobinventorypolicy.Setup,
		storagecontainer.Setup,
		storagedatalakegen2filesystem.Setup,
		storagedatalakegen2path.Setup,
		storageencryptionscope.Setup,
		storagemanagementpolicy.Setup,
		storageobjectreplication.Setup,
		storagequeue.Setup,
		storageshare.Setup,
		storagesharedirectory.Setup,
		storagesync.Setup,
		storagesynccloudendpoint.Setup,
		storagetable.Setup,
		storagetableentity.Setup,
		streamanalyticsfunctionjavascriptudf.Setup,
		streamanalyticsjob.Setup,
		streamanalyticsoutputblob.Setup,
		streamanalyticsoutputeventhub.Setup,
		streamanalyticsoutputmssql.Setup,
		streamanalyticsoutputservicebusqueue.Setup,
		streamanalyticsoutputservicebustopic.Setup,
		streamanalyticsreferenceinputblob.Setup,
		streamanalyticsstreaminputblob.Setup,
		streamanalyticsstreaminputeventhub.Setup,
		streamanalyticsstreaminputiothub.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		subscription.Setup,
		subscriptionpolicyassignment.Setup,
		subscriptiontemplatedeployment.Setup,
		synapsefirewallrule.Setup,
		synapseintegrationruntimeazure.Setup,
		synapseintegrationruntimeselfhosted.Setup,
		synapselinkedservice.Setup,
		synapsemanagedprivateendpoint.Setup,
		synapseprivatelinkhub.Setup,
		synapseroleassignment.Setup,
		synapsesparkpool.Setup,
		synapsesqlpool.Setup,
		synapsesqlpoolextendedauditingpolicy.Setup,
		synapsesqlpoolsecurityalertpolicy.Setup,
		synapsesqlpoolvulnerabilityassessment.Setup,
		synapseworkspace.Setup,
		synapseworkspaceextendedauditingpolicy.Setup,
		synapseworkspacekey.Setup,
		synapseworkspacesecurityalertpolicy.Setup,
		synapseworkspacevulnerabilityassessment.Setup,
		templatedeployment.Setup,
		tenanttemplatedeployment.Setup,
		trafficmanagerendpoint.Setup,
		trafficmanagerprofile.Setup,
		userassignedidentity.Setup,
		videoanalyzer.Setup,
		videoanalyzeredgemodule.Setup,
		virtualdesktopapplication.Setup,
		virtualdesktophostpool.Setup,
		virtualhubconnection.Setup,
		virtualhubip.Setup,
		virtualhubroutetable.Setup,
		virtualhubsecuritypartnerprovider.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		vmwarecluster.Setup,
		vmwareexpressrouteauthorization.Setup,
		vmwareprivatecloud.Setup,
		vpngateway.Setup,
		vpngatewayconnection.Setup,
		vpnserverconfiguration.Setup,
		vpnsite.Setup,
		webapplicationfirewallpolicy.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, concurrency); err != nil {
			return err
		}
	}
	return nil
}
