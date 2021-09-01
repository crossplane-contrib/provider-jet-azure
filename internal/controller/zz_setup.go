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

	activedirectorydomainservice "github.com/ulucinar/provider-tf-azure/internal/controller/active/activedirectorydomainservice"
	activedirectorydomainservicereplicaset "github.com/ulucinar/provider-tf-azure/internal/controller/active/activedirectorydomainservicereplicaset"
	advancedthreatprotection "github.com/ulucinar/provider-tf-azure/internal/controller/advanced/advancedthreatprotection"
	analysisservicesserver "github.com/ulucinar/provider-tf-azure/internal/controller/analysis/analysisservicesserver"
	apimanagement "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagement"
	apimanagementapi "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapi"
	apimanagementapidiagnostic "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapidiagnostic"
	apimanagementapioperation "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapioperation"
	apimanagementapioperationpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapioperationpolicy"
	apimanagementapioperationtag "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapioperationtag"
	apimanagementapipolicy "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapipolicy"
	apimanagementapirelease "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapirelease"
	apimanagementapischema "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapischema"
	apimanagementapiversionset "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementapiversionset"
	apimanagementauthorizationserver "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementauthorizationserver"
	apimanagementbackend "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementbackend"
	apimanagementcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementcertificate"
	apimanagementcustomdomain "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementcustomdomain"
	apimanagementdiagnostic "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementdiagnostic"
	apimanagementemailtemplate "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementemailtemplate"
	apimanagementgateway "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementgateway"
	apimanagementgatewayapi "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementgatewayapi"
	apimanagementgroupuser "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementgroupuser"
	apimanagementidentityprovideraad "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityprovideraad"
	apimanagementidentityprovideraadb2c "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityprovideraadb2c"
	apimanagementidentityproviderfacebook "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityproviderfacebook"
	apimanagementidentityprovidergoogle "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityprovidergoogle"
	apimanagementidentityprovidermicrosoft "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityprovidermicrosoft"
	apimanagementidentityprovidertwitter "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementidentityprovidertwitter"
	apimanagementlogger "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementlogger"
	apimanagementnamedvalue "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementnamedvalue"
	apimanagementnotificationrecipientemail "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementnotificationrecipientemail"
	apimanagementopenidconnectprovider "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementopenidconnectprovider"
	apimanagementpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementpolicy"
	apimanagementproduct "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementproduct"
	apimanagementproductapi "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementproductapi"
	apimanagementproductpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementproductpolicy"
	apimanagementproperty "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementproperty"
	apimanagementrediscache "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementrediscache"
	apimanagementsubscription "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementsubscription"
	apimanagementtag "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementtag"
	apimanagementuser "github.com/ulucinar/provider-tf-azure/internal/controller/api/apimanagementuser"
	appconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/app/appconfiguration"
	appservice "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservice"
	appserviceactiveslot "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceactiveslot"
	appservicecertificate "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicecertificate"
	appservicecertificatebinding "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicecertificatebinding"
	appservicecertificateorder "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicecertificateorder"
	appservicecustomhostnamebinding "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicecustomhostnamebinding"
	appserviceenvironment "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceenvironment"
	appserviceenvironmentv3 "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceenvironmentv3"
	appservicehybridconnection "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicehybridconnection"
	appservicemanagedcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicemanagedcertificate"
	appserviceplan "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceplan"
	appserviceslot "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceslot"
	appserviceslotvirtualnetworkswiftconnection "github.com/ulucinar/provider-tf-azure/internal/controller/app/appserviceslotvirtualnetworkswiftconnection"
	appservicesourcecontroltoken "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicesourcecontroltoken"
	appservicevirtualnetworkswiftconnection "github.com/ulucinar/provider-tf-azure/internal/controller/app/appservicevirtualnetworkswiftconnection"
	applicationgateway "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationgateway"
	applicationinsights "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationinsights"
	applicationinsightsanalyticsitem "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationinsightssmartdetectionrule"
	applicationinsightswebtest "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationinsightswebtest"
	applicationsecuritygroup "github.com/ulucinar/provider-tf-azure/internal/controller/application/applicationsecuritygroup"
	attestationprovider "github.com/ulucinar/provider-tf-azure/internal/controller/attestation/attestationprovider"
	automationaccount "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationaccount"
	automationcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationcertificate"
	automationconnection "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationconnection"
	automationconnectioncertificate "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationconnectioncertificate"
	automationconnectionclassiccertificate "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationconnectionclassiccertificate"
	automationconnectionserviceprincipal "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationconnectionserviceprincipal"
	automationcredential "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationcredential"
	automationdscconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationdscconfiguration"
	automationdscnodeconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationdscnodeconfiguration"
	automationjobschedule "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationjobschedule"
	automationmodule "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationmodule"
	automationrunbook "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationrunbook"
	automationschedule "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationschedule"
	automationvariablebool "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationvariablebool"
	automationvariabledatetime "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationvariabledatetime"
	automationvariableint "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationvariableint"
	automationvariablestring "github.com/ulucinar/provider-tf-azure/internal/controller/automation/automationvariablestring"
	availabilityset "github.com/ulucinar/provider-tf-azure/internal/controller/availability/availabilityset"
	backupcontainerstorageaccount "github.com/ulucinar/provider-tf-azure/internal/controller/backup/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/ulucinar/provider-tf-azure/internal/controller/backup/backuppolicyfileshare"
	backuppolicyvm "github.com/ulucinar/provider-tf-azure/internal/controller/backup/backuppolicyvm"
	backupprotectedfileshare "github.com/ulucinar/provider-tf-azure/internal/controller/backup/backupprotectedfileshare"
	backupprotectedvm "github.com/ulucinar/provider-tf-azure/internal/controller/backup/backupprotectedvm"
	bastionhost "github.com/ulucinar/provider-tf-azure/internal/controller/bastion/bastionhost"
	batchaccount "github.com/ulucinar/provider-tf-azure/internal/controller/batch/batchaccount"
	batchapplication "github.com/ulucinar/provider-tf-azure/internal/controller/batch/batchapplication"
	batchcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/batch/batchcertificate"
	batchjob "github.com/ulucinar/provider-tf-azure/internal/controller/batch/batchjob"
	batchpool "github.com/ulucinar/provider-tf-azure/internal/controller/batch/batchpool"
	blueprintassignment "github.com/ulucinar/provider-tf-azure/internal/controller/blueprint/blueprintassignment"
	botchannelalexa "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelalexa"
	botchanneldirectline "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchanneldirectline"
	botchanneldirectlinespeech "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchanneldirectlinespeech"
	botchannelemail "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelemail"
	botchannelfacebook "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelfacebook"
	botchannelline "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelline"
	botchannelmsteams "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelmsteams"
	botchannelslack "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelslack"
	botchannelsms "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelsms"
	botchannelsregistration "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelsregistration"
	botchannelwebchat "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botchannelwebchat"
	botconnection "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botconnection"
	botwebapp "github.com/ulucinar/provider-tf-azure/internal/controller/bot/botwebapp"
	cdnendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/cdn/cdnendpoint"
	cdnendpointcustomdomain "github.com/ulucinar/provider-tf-azure/internal/controller/cdn/cdnendpointcustomdomain"
	cdnprofile "github.com/ulucinar/provider-tf-azure/internal/controller/cdn/cdnprofile"
	cognitiveaccount "github.com/ulucinar/provider-tf-azure/internal/controller/cognitive/cognitiveaccount"
	communicationservice "github.com/ulucinar/provider-tf-azure/internal/controller/communication/communicationservice"
	config "github.com/ulucinar/provider-tf-azure/internal/controller/config"
	consumptionbudgetresourcegroup "github.com/ulucinar/provider-tf-azure/internal/controller/consumption/consumptionbudgetresourcegroup"
	consumptionbudgetsubscription "github.com/ulucinar/provider-tf-azure/internal/controller/consumption/consumptionbudgetsubscription"
	containergroup "github.com/ulucinar/provider-tf-azure/internal/controller/container/containergroup"
	containerregistry "github.com/ulucinar/provider-tf-azure/internal/controller/container/containerregistry"
	containerregistryscopemap "github.com/ulucinar/provider-tf-azure/internal/controller/container/containerregistryscopemap"
	containerregistrytoken "github.com/ulucinar/provider-tf-azure/internal/controller/container/containerregistrytoken"
	containerregistrywebhook "github.com/ulucinar/provider-tf-azure/internal/controller/container/containerregistrywebhook"
	cosmosdbaccount "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbaccount"
	cosmosdbcassandrakeyspace "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbcassandrakeyspace"
	cosmosdbcassandratable "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbcassandratable"
	cosmosdbgremlindatabase "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbgremlindatabase"
	cosmosdbgremlingraph "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbgremlingraph"
	cosmosdbmongocollection "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbmongocollection"
	cosmosdbmongodatabase "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbmongodatabase"
	cosmosdbnotebookworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbnotebookworkspace"
	cosmosdbsqlcontainer "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlcontainer"
	cosmosdbsqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqldatabase"
	cosmosdbsqlfunction "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlfunction"
	cosmosdbsqlstoredprocedure "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqlstoredprocedure"
	cosmosdbsqltrigger "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbsqltrigger"
	cosmosdbtable "github.com/ulucinar/provider-tf-azure/internal/controller/cosmosdb/cosmosdbtable"
	costmanagementexportresourcegroup "github.com/ulucinar/provider-tf-azure/internal/controller/cost/costmanagementexportresourcegroup"
	customprovider "github.com/ulucinar/provider-tf-azure/internal/controller/custom/customprovider"
	dashboard "github.com/ulucinar/provider-tf-azure/internal/controller/dashboard/dashboard"
	datafactory "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactory"
	datafactorycustomdataset "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorycustomdataset"
	datafactorydataflow "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydataflow"
	datafactorydatasetazureblob "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetazureblob"
	datafactorydatasetbinary "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetbinary"
	datafactorydatasetcosmosdbsqlapi "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetcosmosdbsqlapi"
	datafactorydatasetdelimitedtext "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetdelimitedtext"
	datafactorydatasethttp "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasethttp"
	datafactorydatasetjson "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetjson"
	datafactorydatasetmysql "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetmysql"
	datafactorydatasetparquet "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetparquet"
	datafactorydatasetpostgresql "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetpostgresql"
	datafactorydatasetsnowflake "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetsnowflake"
	datafactorydatasetsqlservertable "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorydatasetsqlservertable"
	datafactoryintegrationruntimeazure "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeazure"
	datafactoryintegrationruntimeazuressis "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeazuressis"
	datafactoryintegrationruntimemanaged "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimemanaged"
	datafactoryintegrationruntimeselfhosted "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactoryintegrationruntimeselfhosted"
	datafactorylinkedcustomservice "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedcustomservice"
	datafactorylinkedserviceazureblobstorage "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazureblobstorage"
	datafactorylinkedserviceazuredatabricks "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuredatabricks"
	datafactorylinkedserviceazurefilestorage "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazurefilestorage"
	datafactorylinkedserviceazurefunction "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazurefunction"
	datafactorylinkedserviceazuresearch "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuresearch"
	datafactorylinkedserviceazuresqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuresqldatabase"
	datafactorylinkedserviceazuretablestorage "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceazuretablestorage"
	datafactorylinkedservicecosmosdb "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicecosmosdb"
	datafactorylinkedservicedatalakestoragegen2 "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicedatalakestoragegen2"
	datafactorylinkedservicekeyvault "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicekeyvault"
	datafactorylinkedservicekusto "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicekusto"
	datafactorylinkedservicemysql "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicemysql"
	datafactorylinkedserviceodata "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceodata"
	datafactorylinkedservicepostgresql "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicepostgresql"
	datafactorylinkedservicesftp "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicesftp"
	datafactorylinkedservicesnowflake "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicesnowflake"
	datafactorylinkedservicesqlserver "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicesqlserver"
	datafactorylinkedservicesynapse "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedservicesynapse"
	datafactorylinkedserviceweb "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorylinkedserviceweb"
	datafactorymanagedprivateendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorymanagedprivateendpoint"
	datafactorypipeline "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorypipeline"
	datafactorytriggerblobevent "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorytriggerblobevent"
	datafactorytriggercustomevent "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorytriggercustomevent"
	datafactorytriggerschedule "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorytriggerschedule"
	datafactorytriggertumblingwindow "github.com/ulucinar/provider-tf-azure/internal/controller/data/datafactorytriggertumblingwindow"
	datalakeanalyticsaccount "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakeanalyticsaccount"
	datalakeanalyticsfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakeanalyticsfirewallrule"
	datalakestore "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakestore"
	datalakestorefile "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakestorefile"
	datalakestorefirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakestorefirewallrule"
	datalakestorevirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/data/datalakestorevirtualnetworkrule"
	dataprotectionbackupinstancedisk "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackupinstancedisk"
	dataprotectionbackupinstancepostgresql "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackupinstancepostgresql"
	dataprotectionbackuppolicyblobstorage "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicyblobstorage"
	dataprotectionbackuppolicydisk "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicydisk"
	dataprotectionbackuppolicypostgresql "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackuppolicypostgresql"
	dataprotectionbackupvault "github.com/ulucinar/provider-tf-azure/internal/controller/data/dataprotectionbackupvault"
	datashare "github.com/ulucinar/provider-tf-azure/internal/controller/data/datashare"
	datashareaccount "github.com/ulucinar/provider-tf-azure/internal/controller/data/datashareaccount"
	datasharedatasetblobstorage "github.com/ulucinar/provider-tf-azure/internal/controller/data/datasharedatasetblobstorage"
	datasharedatasetdatalakegen1 "github.com/ulucinar/provider-tf-azure/internal/controller/data/datasharedatasetdatalakegen1"
	datasharedatasetdatalakegen2 "github.com/ulucinar/provider-tf-azure/internal/controller/data/datasharedatasetdatalakegen2"
	datasharedatasetkustocluster "github.com/ulucinar/provider-tf-azure/internal/controller/data/datasharedatasetkustocluster"
	datasharedatasetkustodatabase "github.com/ulucinar/provider-tf-azure/internal/controller/data/datasharedatasetkustodatabase"
	databasemigrationproject "github.com/ulucinar/provider-tf-azure/internal/controller/database/databasemigrationproject"
	databasemigrationservice "github.com/ulucinar/provider-tf-azure/internal/controller/database/databasemigrationservice"
	databoxedgedevice "github.com/ulucinar/provider-tf-azure/internal/controller/databox/databoxedgedevice"
	databoxedgeorder "github.com/ulucinar/provider-tf-azure/internal/controller/databox/databoxedgeorder"
	databricksworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/databricks/databricksworkspace"
	databricksworkspacecustomermanagedkey "github.com/ulucinar/provider-tf-azure/internal/controller/databricks/databricksworkspacecustomermanagedkey"
	dedicatedhardwaresecuritymodule "github.com/ulucinar/provider-tf-azure/internal/controller/dedicated/dedicatedhardwaresecuritymodule"
	dedicatedhost "github.com/ulucinar/provider-tf-azure/internal/controller/dedicated/dedicatedhost"
	devtestglobalvmshutdownschedule "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestglobalvmshutdownschedule"
	devtestlab "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestlab"
	devtestlinuxvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestlinuxvirtualmachine"
	devtestpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestpolicy"
	devtestschedule "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestschedule"
	devtestvirtualnetwork "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestvirtualnetwork"
	devtestwindowsvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/dev/devtestwindowsvirtualmachine"
	devspacecontroller "github.com/ulucinar/provider-tf-azure/internal/controller/devspace/devspacecontroller"
	digitaltwinsendpointeventgrid "github.com/ulucinar/provider-tf-azure/internal/controller/digital/digitaltwinsendpointeventgrid"
	digitaltwinsendpointeventhub "github.com/ulucinar/provider-tf-azure/internal/controller/digital/digitaltwinsendpointeventhub"
	digitaltwinsendpointservicebus "github.com/ulucinar/provider-tf-azure/internal/controller/digital/digitaltwinsendpointservicebus"
	digitaltwinsinstance "github.com/ulucinar/provider-tf-azure/internal/controller/digital/digitaltwinsinstance"
	diskaccess "github.com/ulucinar/provider-tf-azure/internal/controller/disk/diskaccess"
	diskencryptionset "github.com/ulucinar/provider-tf-azure/internal/controller/disk/diskencryptionset"
	dnsaaaarecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnsaaaarecord"
	dnsarecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnsarecord"
	dnscaarecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnscaarecord"
	dnscnamerecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnscnamerecord"
	dnsmxrecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnsmxrecord"
	dnsnsrecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnsnsrecord"
	dnsptrrecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnsptrrecord"
	dnssrvrecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnssrvrecord"
	dnstxtrecord "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnstxtrecord"
	dnszone "github.com/ulucinar/provider-tf-azure/internal/controller/dns/dnszone"
	eventgriddomain "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgriddomain"
	eventgriddomaintopic "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgriddomaintopic"
	eventgrideventsubscription "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgrideventsubscription"
	eventgridsystemtopic "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgridsystemtopic"
	eventgridsystemtopiceventsubscription "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgridsystemtopiceventsubscription"
	eventgridtopic "github.com/ulucinar/provider-tf-azure/internal/controller/eventgrid/eventgridtopic"
	eventhub "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhub"
	eventhubauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubauthorizationrule"
	eventhubcluster "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubcluster"
	eventhubconsumergroup "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubconsumergroup"
	eventhubnamespace "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubnamespace"
	eventhubnamespaceauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubnamespaceauthorizationrule"
	eventhubnamespacecustomermanagedkey "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubnamespacecustomermanagedkey"
	eventhubnamespacedisasterrecoveryconfig "github.com/ulucinar/provider-tf-azure/internal/controller/eventhub/eventhubnamespacedisasterrecoveryconfig"
	expressroutecircuit "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressroutecircuit"
	expressroutecircuitauthorization "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressroutecircuitpeering"
	expressrouteconnection "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressrouteconnection"
	expressroutegateway "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressroutegateway"
	expressrouteport "github.com/ulucinar/provider-tf-azure/internal/controller/express/expressrouteport"
	firewall "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewall"
	firewallapplicationrulecollection "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewallnetworkrulecollection"
	firewallpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/ulucinar/provider-tf-azure/internal/controller/firewall/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/ulucinar/provider-tf-azure/internal/controller/frontdoor/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/frontdoor/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/frontdoor/frontdoorfirewallpolicy"
	functionapp "github.com/ulucinar/provider-tf-azure/internal/controller/function/functionapp"
	functionappslot "github.com/ulucinar/provider-tf-azure/internal/controller/function/functionappslot"
	hdinsighthadoopcluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsighthadoopcluster"
	hdinsighthbasecluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsighthbasecluster"
	hdinsightinteractivequerycluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightinteractivequerycluster"
	hdinsightkafkacluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightkafkacluster"
	hdinsightmlservicescluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightmlservicescluster"
	hdinsightrservercluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightrservercluster"
	hdinsightsparkcluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightsparkcluster"
	hdinsightstormcluster "github.com/ulucinar/provider-tf-azure/internal/controller/hdinsight/hdinsightstormcluster"
	healthbot "github.com/ulucinar/provider-tf-azure/internal/controller/healthbot/healthbot"
	healthcareservice "github.com/ulucinar/provider-tf-azure/internal/controller/healthcare/healthcareservice"
	hpccache "github.com/ulucinar/provider-tf-azure/internal/controller/hpc/hpccache"
	hpccacheaccesspolicy "github.com/ulucinar/provider-tf-azure/internal/controller/hpc/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/ulucinar/provider-tf-azure/internal/controller/hpc/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/ulucinar/provider-tf-azure/internal/controller/hpc/hpccacheblobtarget"
	hpccachenfstarget "github.com/ulucinar/provider-tf-azure/internal/controller/hpc/hpccachenfstarget"
	image "github.com/ulucinar/provider-tf-azure/internal/controller/image/image"
	integrationserviceenvironment "github.com/ulucinar/provider-tf-azure/internal/controller/integration/integrationserviceenvironment"
	iotsecuritydevicegroup "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iotsecuritysolution"
	iottimeseriesinsightsaccesspolicy "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsaccesspolicy"
	iottimeseriesinsightseventsourceiothub "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iottimeseriesinsightseventsourceiothub"
	iottimeseriesinsightsgen2environment "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsgen2environment"
	iottimeseriesinsightsreferencedataset "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsreferencedataset"
	iottimeseriesinsightsstandardenvironment "github.com/ulucinar/provider-tf-azure/internal/controller/iot/iottimeseriesinsightsstandardenvironment"
	iotcentralapplication "github.com/ulucinar/provider-tf-azure/internal/controller/iotcentral/iotcentralapplication"
	iothub "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothub"
	iothubconsumergroup "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubconsumergroup"
	iothubdps "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubdps"
	iothubdpscertificate "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubendpointstoragecontainer"
	iothubenrichment "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubenrichment"
	iothubfallbackroute "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubfallbackroute"
	iothubroute "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubroute"
	iothubsharedaccesspolicy "github.com/ulucinar/provider-tf-azure/internal/controller/iothub/iothubsharedaccesspolicy"
	ipgroup "github.com/ulucinar/provider-tf-azure/internal/controller/ip/ipgroup"
	keyvault "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvault"
	keyvaultaccesspolicy "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultaccesspolicy"
	keyvaultcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultcertificate"
	keyvaultcertificateissuer "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultcertificateissuer"
	keyvaultkey "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultkey"
	keyvaultmanagedhardwaresecuritymodule "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultmanagedhardwaresecuritymodule"
	keyvaultsecret "github.com/ulucinar/provider-tf-azure/internal/controller/key/keyvaultsecret"
	kubernetescluster "github.com/ulucinar/provider-tf-azure/internal/controller/kubernetes/kubernetescluster"
	kubernetesclusternodepool "github.com/ulucinar/provider-tf-azure/internal/controller/kubernetes/kubernetesclusternodepool"
	kustoattacheddatabaseconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoattacheddatabaseconfiguration"
	kustocluster "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustocluster"
	kustoclustercustomermanagedkey "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoclustercustomermanagedkey"
	kustoclusterprincipalassignment "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoclusterprincipalassignment"
	kustodatabase "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustodatabase"
	kustodatabaseprincipal "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustodatabaseprincipal"
	kustodatabaseprincipalassignment "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustodatabaseprincipalassignment"
	kustoeventgriddataconnection "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoeventgriddataconnection"
	kustoeventhubdataconnection "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoeventhubdataconnection"
	kustoiothubdataconnection "github.com/ulucinar/provider-tf-azure/internal/controller/kusto/kustoiothubdataconnection"
	lb "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lb"
	lbbackendaddresspool "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbbackendaddresspool"
	lbbackendaddresspooladdress "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbbackendaddresspooladdress"
	lbnatpool "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbnatpool"
	lbnatrule "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbnatrule"
	lboutboundrule "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lboutboundrule"
	lbprobe "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbprobe"
	lbrule "github.com/ulucinar/provider-tf-azure/internal/controller/lb/lbrule"
	lighthouseassignment "github.com/ulucinar/provider-tf-azure/internal/controller/lighthouse/lighthouseassignment"
	lighthousedefinition "github.com/ulucinar/provider-tf-azure/internal/controller/lighthouse/lighthousedefinition"
	linuxvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/linux/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/ulucinar/provider-tf-azure/internal/controller/linux/linuxvirtualmachinescaleset"
	localnetworkgateway "github.com/ulucinar/provider-tf-azure/internal/controller/local/localnetworkgateway"
	loganalyticscluster "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticscluster"
	loganalyticsclustercustomermanagedkey "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsclustercustomermanagedkey"
	loganalyticsdataexportrule "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticslinkedstorageaccount"
	loganalyticssavedsearch "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticssavedsearch"
	loganalyticssolution "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticssolution"
	loganalyticsstorageinsights "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsstorageinsights"
	loganalyticsworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/log/loganalyticsworkspace"
	logicappactioncustom "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappactioncustom"
	logicappactionhttp "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappactionhttp"
	logicappintegrationaccount "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappintegrationaccount"
	logicappintegrationaccountcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappintegrationaccountcertificate"
	logicappintegrationaccountschema "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappintegrationaccountschema"
	logicappintegrationaccountsession "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappintegrationaccountsession"
	logicapptriggercustom "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicapptriggercustom"
	logicapptriggerhttprequest "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicapptriggerhttprequest"
	logicapptriggerrecurrence "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicapptriggerrecurrence"
	logicappworkflow "github.com/ulucinar/provider-tf-azure/internal/controller/logic/logicappworkflow"
	machinelearningcomputecluster "github.com/ulucinar/provider-tf-azure/internal/controller/machine/machinelearningcomputecluster"
	machinelearningcomputeinstance "github.com/ulucinar/provider-tf-azure/internal/controller/machine/machinelearningcomputeinstance"
	machinelearninginferencecluster "github.com/ulucinar/provider-tf-azure/internal/controller/machine/machinelearninginferencecluster"
	machinelearningsynapsespark "github.com/ulucinar/provider-tf-azure/internal/controller/machine/machinelearningsynapsespark"
	machinelearningworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/machine/machinelearningworkspace"
	maintenanceassignmentdedicatedhost "github.com/ulucinar/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceassignmentvirtualmachinescaleset "github.com/ulucinar/provider-tf-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachinescaleset"
	maintenanceconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/maintenance/maintenanceconfiguration"
	managedapplication "github.com/ulucinar/provider-tf-azure/internal/controller/managed/managedapplication"
	managedapplicationdefinition "github.com/ulucinar/provider-tf-azure/internal/controller/managed/managedapplicationdefinition"
	manageddisk "github.com/ulucinar/provider-tf-azure/internal/controller/managed/manageddisk"
	managementgroup "github.com/ulucinar/provider-tf-azure/internal/controller/management/managementgroup"
	managementgrouppolicyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/management/managementgrouppolicyassignment"
	managementgroupsubscriptionassociation "github.com/ulucinar/provider-tf-azure/internal/controller/management/managementgroupsubscriptionassociation"
	managementgrouptemplatedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/management/managementgrouptemplatedeployment"
	managementlock "github.com/ulucinar/provider-tf-azure/internal/controller/management/managementlock"
	mapsaccount "github.com/ulucinar/provider-tf-azure/internal/controller/maps/mapsaccount"
	mariadbconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/mariadb/mariadbconfiguration"
	mariadbdatabase "github.com/ulucinar/provider-tf-azure/internal/controller/mariadb/mariadbdatabase"
	mariadbfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/mariadb/mariadbfirewallrule"
	mariadbserver "github.com/ulucinar/provider-tf-azure/internal/controller/mariadb/mariadbserver"
	mariadbvirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/mariadb/mariadbvirtualnetworkrule"
	marketplaceagreement "github.com/ulucinar/provider-tf-azure/internal/controller/marketplace/marketplaceagreement"
	mediaasset "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediaasset"
	mediaassetfilter "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediaassetfilter"
	mediacontentkeypolicy "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediacontentkeypolicy"
	mediajob "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediajob"
	medialiveevent "github.com/ulucinar/provider-tf-azure/internal/controller/media/medialiveevent"
	medialiveeventoutput "github.com/ulucinar/provider-tf-azure/internal/controller/media/medialiveeventoutput"
	mediaservicesaccount "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediaservicesaccount"
	mediastreamingendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediastreamingendpoint"
	mediastreaminglocator "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediastreaminglocator"
	mediastreamingpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediastreamingpolicy"
	mediatransform "github.com/ulucinar/provider-tf-azure/internal/controller/media/mediatransform"
	monitoraaddiagnosticsetting "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitoraaddiagnosticsetting"
	monitoractiongroup "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitoractiongroup"
	monitoractionruleactiongroup "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitoractionrulesuppression"
	monitoractivitylogalert "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitoractivitylogalert"
	monitorautoscalesetting "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitorautoscalesetting"
	monitordiagnosticsetting "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitordiagnosticsetting"
	monitorlogprofile "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitorlogprofile"
	monitormetricalert "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitormetricalert"
	monitorscheduledqueryrulesalert "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitorscheduledqueryrulesalert"
	monitorscheduledqueryruleslog "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitorscheduledqueryruleslog"
	monitorsmartdetectoralertrule "github.com/ulucinar/provider-tf-azure/internal/controller/monitor/monitorsmartdetectoralertrule"
	mssqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlelasticpool"
	mssqlfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlfirewallrule"
	mssqljobagent "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqljobagent"
	mssqljobcredential "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqljobcredential"
	mssqlserver "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlserver"
	mssqlserversecurityalertpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlservervulnerabilityassessment"
	mssqlvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlvirtualmachine"
	mssqlvirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/mssql/mssqlvirtualnetworkrule"
	mysqlactivedirectoryadministrator "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlactivedirectoryadministrator"
	mysqlconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlconfiguration"
	mysqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqldatabase"
	mysqlfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlfirewallrule"
	mysqlserver "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlserver"
	mysqlserverkey "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlserverkey"
	mysqlvirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/mysql/mysqlvirtualnetworkrule"
	natgateway "github.com/ulucinar/provider-tf-azure/internal/controller/nat/natgateway"
	natgatewaypublicipassociation "github.com/ulucinar/provider-tf-azure/internal/controller/nat/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/ulucinar/provider-tf-azure/internal/controller/nat/natgatewaypublicipprefixassociation"
	netappaccount "github.com/ulucinar/provider-tf-azure/internal/controller/netapp/netappaccount"
	netapppool "github.com/ulucinar/provider-tf-azure/internal/controller/netapp/netapppool"
	netappsnapshot "github.com/ulucinar/provider-tf-azure/internal/controller/netapp/netappsnapshot"
	netappvolume "github.com/ulucinar/provider-tf-azure/internal/controller/netapp/netappvolume"
	networkconnectionmonitor "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkconnectionmonitor"
	networkddosprotectionplan "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkddosprotectionplan"
	networkinterface "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkinterfacesecuritygroupassociation"
	networkpacketcapture "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkpacketcapture"
	networkprofile "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkprofile"
	networksecuritygroup "github.com/ulucinar/provider-tf-azure/internal/controller/network/networksecuritygroup"
	networksecurityrule "github.com/ulucinar/provider-tf-azure/internal/controller/network/networksecurityrule"
	networkwatcher "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkwatcher"
	networkwatcherflowlog "github.com/ulucinar/provider-tf-azure/internal/controller/network/networkwatcherflowlog"
	notificationhub "github.com/ulucinar/provider-tf-azure/internal/controller/notification/notificationhub"
	notificationhubauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/notification/notificationhubauthorizationrule"
	notificationhubnamespace "github.com/ulucinar/provider-tf-azure/internal/controller/notification/notificationhubnamespace"
	orchestratedvirtualmachinescaleset "github.com/ulucinar/provider-tf-azure/internal/controller/orchestrated/orchestratedvirtualmachinescaleset"
	packetcapture "github.com/ulucinar/provider-tf-azure/internal/controller/packet/packetcapture"
	pointtositevpngateway "github.com/ulucinar/provider-tf-azure/internal/controller/point/pointtositevpngateway"
	policyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/policy/policyassignment"
	policydefinition "github.com/ulucinar/provider-tf-azure/internal/controller/policy/policydefinition"
	policyremediation "github.com/ulucinar/provider-tf-azure/internal/controller/policy/policyremediation"
	policysetdefinition "github.com/ulucinar/provider-tf-azure/internal/controller/policy/policysetdefinition"
	policyvirtualmachineconfigurationassignment "github.com/ulucinar/provider-tf-azure/internal/controller/policy/policyvirtualmachineconfigurationassignment"
	portaltenantconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/portal/portaltenantconfiguration"
	postgresqlactivedirectoryadministrator "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlactivedirectoryadministrator"
	postgresqlconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlconfiguration"
	postgresqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqldatabase"
	postgresqlfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlfirewallrule"
	postgresqlflexibleserver "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserver"
	postgresqlflexibleserverconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverconfiguration"
	postgresqlflexibleserverdatabase "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverdatabase"
	postgresqlflexibleserverfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlflexibleserverfirewallrule"
	postgresqlserver "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlserver"
	postgresqlserverkey "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlserverkey"
	postgresqlvirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/postgresql/postgresqlvirtualnetworkrule"
	powerbiembedded "github.com/ulucinar/provider-tf-azure/internal/controller/powerbi/powerbiembedded"
	privatednsaaaarecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednsaaaarecord"
	privatednsarecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednsarecord"
	privatednscnamerecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednscnamerecord"
	privatednsmxrecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednsmxrecord"
	privatednsptrrecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednsptrrecord"
	privatednssrvrecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednssrvrecord"
	privatednstxtrecord "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednstxtrecord"
	privatednszone "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednszone"
	privatednszonevirtualnetworklink "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatednszonevirtualnetworklink"
	privateendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/private/privateendpoint"
	privatelinkservice "github.com/ulucinar/provider-tf-azure/internal/controller/private/privatelinkservice"
	proximityplacementgroup "github.com/ulucinar/provider-tf-azure/internal/controller/proximity/proximityplacementgroup"
	publicip "github.com/ulucinar/provider-tf-azure/internal/controller/public/publicip"
	publicipprefix "github.com/ulucinar/provider-tf-azure/internal/controller/public/publicipprefix"
	purviewaccount "github.com/ulucinar/provider-tf-azure/internal/controller/purview/purviewaccount"
	recoveryservicesvault "github.com/ulucinar/provider-tf-azure/internal/controller/recovery/recoveryservicesvault"
	rediscache "github.com/ulucinar/provider-tf-azure/internal/controller/redis/rediscache"
	redisenterprisecluster "github.com/ulucinar/provider-tf-azure/internal/controller/redis/redisenterprisecluster"
	redisenterprisedatabase "github.com/ulucinar/provider-tf-azure/internal/controller/redis/redisenterprisedatabase"
	redisfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/redis/redisfirewallrule"
	redislinkedserver "github.com/ulucinar/provider-tf-azure/internal/controller/redis/redislinkedserver"
	relayhybridconnection "github.com/ulucinar/provider-tf-azure/internal/controller/relay/relayhybridconnection"
	relayhybridconnectionauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/relay/relayhybridconnectionauthorizationrule"
	relaynamespace "github.com/ulucinar/provider-tf-azure/internal/controller/relay/relaynamespace"
	relaynamespaceauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/relay/relaynamespaceauthorizationrule"
	resourcegroup "github.com/ulucinar/provider-tf-azure/internal/controller/resource/resourcegroup"
	resourcegrouppolicyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/resource/resourcegrouppolicyassignment"
	resourcegrouptemplatedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/resource/resourcegrouptemplatedeployment"
	resourcepolicyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/resource/resourcepolicyassignment"
	resourceproviderregistration "github.com/ulucinar/provider-tf-azure/internal/controller/resource/resourceproviderregistration"
	roleassignment "github.com/ulucinar/provider-tf-azure/internal/controller/role/roleassignment"
	roledefinition "github.com/ulucinar/provider-tf-azure/internal/controller/role/roledefinition"
	route "github.com/ulucinar/provider-tf-azure/internal/controller/route/route"
	routefilter "github.com/ulucinar/provider-tf-azure/internal/controller/route/routefilter"
	routetable "github.com/ulucinar/provider-tf-azure/internal/controller/route/routetable"
	searchservice "github.com/ulucinar/provider-tf-azure/internal/controller/search/searchservice"
	securitycenterassessment "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterassessment"
	securitycenterassessmentmetadata "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterassessmentmetadata"
	securitycenterassessmentpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterassessmentpolicy"
	securitycenterautomation "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterautomation"
	securitycenterautoprovisioning "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterservervulnerabilityassessment"
	securitycentersetting "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/security/securitycenterworkspace"
	sentinelalertrulefusion "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentinelalertrulemssecurityincident"
	sentinelalertrulescheduled "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentinelalertrulescheduled"
	sentineldataconnectorawscloudtrail "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorawscloudtrail"
	sentineldataconnectorazureactivedirectory "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazureactivedirectory"
	sentineldataconnectorazureadvancedthreatprotection "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazureadvancedthreatprotection"
	sentineldataconnectorazuresecuritycenter "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorazuresecuritycenter"
	sentineldataconnectormicrosoftcloudappsecurity "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectormicrosoftcloudappsecurity"
	sentineldataconnectoroffice365 "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectoroffice365"
	sentineldataconnectorthreatintelligence "github.com/ulucinar/provider-tf-azure/internal/controller/sentinel/sentineldataconnectorthreatintelligence"
	servicefabriccluster "github.com/ulucinar/provider-tf-azure/internal/controller/service/servicefabriccluster"
	servicefabricmeshapplication "github.com/ulucinar/provider-tf-azure/internal/controller/service/servicefabricmeshapplication"
	servicefabricmeshlocalnetwork "github.com/ulucinar/provider-tf-azure/internal/controller/service/servicefabricmeshlocalnetwork"
	servicefabricmeshsecret "github.com/ulucinar/provider-tf-azure/internal/controller/service/servicefabricmeshsecret"
	servicefabricmeshsecretvalue "github.com/ulucinar/provider-tf-azure/internal/controller/service/servicefabricmeshsecretvalue"
	servicebusnamespace "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusnamespace"
	servicebusnamespaceauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusnamespaceauthorizationrule"
	servicebusnamespacedisasterrecoveryconfig "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusnamespacedisasterrecoveryconfig"
	servicebusnamespacenetworkruleset "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusnamespacenetworkruleset"
	servicebusqueue "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusqueue"
	servicebusqueueauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebusqueueauthorizationrule"
	servicebussubscription "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebussubscription"
	servicebussubscriptionrule "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebussubscriptionrule"
	servicebustopic "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebustopic"
	servicebustopicauthorizationrule "github.com/ulucinar/provider-tf-azure/internal/controller/servicebus/servicebustopicauthorizationrule"
	sharedimage "github.com/ulucinar/provider-tf-azure/internal/controller/shared/sharedimage"
	sharedimagegallery "github.com/ulucinar/provider-tf-azure/internal/controller/shared/sharedimagegallery"
	sharedimageversion "github.com/ulucinar/provider-tf-azure/internal/controller/shared/sharedimageversion"
	signalrservice "github.com/ulucinar/provider-tf-azure/internal/controller/signalr/signalrservice"
	signalrservicenetworkacl "github.com/ulucinar/provider-tf-azure/internal/controller/signalr/signalrservicenetworkacl"
	siterecoveryfabric "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoveryprotectioncontainermapping"
	siterecoveryreplicatedvm "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoveryreplicatedvm"
	siterecoveryreplicationpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/site/siterecoveryreplicationpolicy"
	snapshot "github.com/ulucinar/provider-tf-azure/internal/controller/snapshot/snapshot"
	spatialanchorsaccount "github.com/ulucinar/provider-tf-azure/internal/controller/spatial/spatialanchorsaccount"
	springcloudactivedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudactivedeployment"
	springcloudapp "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudapp"
	springcloudappcosmosdbassociation "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudappredisassociation"
	springcloudcertificate "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudcertificate"
	springcloudcustomdomain "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudcustomdomain"
	springcloudjavadeployment "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudjavadeployment"
	springcloudservice "github.com/ulucinar/provider-tf-azure/internal/controller/spring/springcloudservice"
	sqlactivedirectoryadministrator "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlactivedirectoryadministrator"
	sqldatabase "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqldatabase"
	sqlelasticpool "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlelasticpool"
	sqlfailovergroup "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlfailovergroup"
	sqlfirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlfirewallrule"
	sqlserver "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlserver"
	sqlvirtualnetworkrule "github.com/ulucinar/provider-tf-azure/internal/controller/sql/sqlvirtualnetworkrule"
	sshpublickey "github.com/ulucinar/provider-tf-azure/internal/controller/ssh/sshpublickey"
	stackhcicluster "github.com/ulucinar/provider-tf-azure/internal/controller/stack/stackhcicluster"
	staticsite "github.com/ulucinar/provider-tf-azure/internal/controller/static/staticsite"
	storageaccount "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageaccount"
	storageaccountcustomermanagedkey "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageaccountcustomermanagedkey"
	storageaccountnetworkrules "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageaccountnetworkrules"
	storageblob "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageblob"
	storageblobinventorypolicy "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageblobinventorypolicy"
	storagecontainer "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagecontainer"
	storagedatalakegen2filesystem "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagedatalakegen2filesystem"
	storagedatalakegen2path "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagedatalakegen2path"
	storageencryptionscope "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageencryptionscope"
	storagemanagementpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagemanagementpolicy"
	storageobjectreplication "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageobjectreplication"
	storagequeue "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagequeue"
	storageshare "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storageshare"
	storagesharedirectory "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagesharedirectory"
	storagesharefile "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagesharefile"
	storagesync "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagesync"
	storagesynccloudendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagesynccloudendpoint"
	storagetable "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagetable"
	storagetableentity "github.com/ulucinar/provider-tf-azure/internal/controller/storage/storagetableentity"
	streamanalyticsfunctionjavascriptudf "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsfunctionjavascriptudf"
	streamanalyticsjob "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsjob"
	streamanalyticsoutputblob "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsoutputblob"
	streamanalyticsoutputeventhub "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsoutputeventhub"
	streamanalyticsoutputmssql "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsoutputmssql"
	streamanalyticsoutputservicebusqueue "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsoutputservicebusqueue"
	streamanalyticsoutputservicebustopic "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsoutputservicebustopic"
	streamanalyticsreferenceinputblob "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsreferenceinputblob"
	streamanalyticsstreaminputblob "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputblob"
	streamanalyticsstreaminputeventhub "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputeventhub"
	streamanalyticsstreaminputiothub "github.com/ulucinar/provider-tf-azure/internal/controller/stream/streamanalyticsstreaminputiothub"
	subnet "github.com/ulucinar/provider-tf-azure/internal/controller/subnet/subnet"
	subnetnatgatewayassociation "github.com/ulucinar/provider-tf-azure/internal/controller/subnet/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/ulucinar/provider-tf-azure/internal/controller/subnet/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/ulucinar/provider-tf-azure/internal/controller/subnet/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/ulucinar/provider-tf-azure/internal/controller/subnet/subnetserviceendpointstoragepolicy"
	subscription "github.com/ulucinar/provider-tf-azure/internal/controller/subscription/subscription"
	subscriptionpolicyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/subscription/subscriptionpolicyassignment"
	subscriptiontemplatedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/subscription/subscriptiontemplatedeployment"
	synapsefirewallrule "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapsefirewallrule"
	synapsemanagedprivateendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapsemanagedprivateendpoint"
	synapseprivatelinkhub "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapseprivatelinkhub"
	synapseroleassignment "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapseroleassignment"
	synapsesparkpool "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapsesparkpool"
	synapsesqlpool "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapsesqlpool"
	synapseworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/synapse/synapseworkspace"
	templatedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/template/templatedeployment"
	tenanttemplatedeployment "github.com/ulucinar/provider-tf-azure/internal/controller/tenant/tenanttemplatedeployment"
	trafficmanagerendpoint "github.com/ulucinar/provider-tf-azure/internal/controller/traffic/trafficmanagerendpoint"
	trafficmanagerprofile "github.com/ulucinar/provider-tf-azure/internal/controller/traffic/trafficmanagerprofile"
	userassignedidentity "github.com/ulucinar/provider-tf-azure/internal/controller/user/userassignedidentity"
	videoanalyzer "github.com/ulucinar/provider-tf-azure/internal/controller/video/videoanalyzer"
	videoanalyzeredgemodule "github.com/ulucinar/provider-tf-azure/internal/controller/video/videoanalyzeredgemodule"
	virtualdesktopapplication "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualdesktopapplication"
	virtualdesktophostpool "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualdesktophostpool"
	virtualdesktopworkspace "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualdesktopworkspace"
	virtualdesktopworkspaceapplicationgroupassociation "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualdesktopworkspaceapplicationgroupassociation"
	virtualhub "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhub"
	virtualhubbgpconnection "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhubbgpconnection"
	virtualhubconnection "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhubconnection"
	virtualhubip "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhubip"
	virtualhubroutetable "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhubroutetable"
	virtualhubsecuritypartnerprovider "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualhubsecuritypartnerprovider"
	virtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachine"
	virtualmachineconfigurationpolicyassignment "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachineconfigurationpolicyassignment"
	virtualmachinedatadiskattachment "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachineextension"
	virtualmachinescaleset "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachinescaleset"
	virtualmachinescalesetextension "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualmachinescalesetextension"
	virtualnetwork "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualnetwork"
	virtualnetworkdnsservers "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualnetworkdnsservers"
	virtualnetworkgateway "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualnetworkpeering"
	virtualwan "github.com/ulucinar/provider-tf-azure/internal/controller/virtual/virtualwan"
	vmwarecluster "github.com/ulucinar/provider-tf-azure/internal/controller/vmware/vmwarecluster"
	vmwareexpressrouteauthorization "github.com/ulucinar/provider-tf-azure/internal/controller/vmware/vmwareexpressrouteauthorization"
	vmwareprivatecloud "github.com/ulucinar/provider-tf-azure/internal/controller/vmware/vmwareprivatecloud"
	vpngateway "github.com/ulucinar/provider-tf-azure/internal/controller/vpn/vpngateway"
	vpngatewayconnection "github.com/ulucinar/provider-tf-azure/internal/controller/vpn/vpngatewayconnection"
	vpnserverconfiguration "github.com/ulucinar/provider-tf-azure/internal/controller/vpn/vpnserverconfiguration"
	vpnsite "github.com/ulucinar/provider-tf-azure/internal/controller/vpn/vpnsite"
	webapplicationfirewallpolicy "github.com/ulucinar/provider-tf-azure/internal/controller/web/webapplicationfirewallpolicy"
	windowsvirtualmachine "github.com/ulucinar/provider-tf-azure/internal/controller/windows/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/ulucinar/provider-tf-azure/internal/controller/windows/windowsvirtualmachinescaleset"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter) error{
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
		automationdscconfiguration.Setup,
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
		communicationservice.Setup,
		config.Setup,
		consumptionbudgetresourcegroup.Setup,
		consumptionbudgetsubscription.Setup,
		containergroup.Setup,
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
		datafactorytriggertumblingwindow.Setup,
		datalakeanalyticsaccount.Setup,
		datalakeanalyticsfirewallrule.Setup,
		datalakestore.Setup,
		datalakestorefile.Setup,
		datalakestorefirewallrule.Setup,
		datalakestorevirtualnetworkrule.Setup,
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
		lb.Setup,
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
		loganalyticsstorageinsights.Setup,
		loganalyticsworkspace.Setup,
		logicappactioncustom.Setup,
		logicappactionhttp.Setup,
		logicappintegrationaccount.Setup,
		logicappintegrationaccountcertificate.Setup,
		logicappintegrationaccountschema.Setup,
		logicappintegrationaccountsession.Setup,
		logicapptriggercustom.Setup,
		logicapptriggerhttprequest.Setup,
		logicapptriggerrecurrence.Setup,
		logicappworkflow.Setup,
		machinelearningcomputecluster.Setup,
		machinelearningcomputeinstance.Setup,
		machinelearninginferencecluster.Setup,
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
		monitorlogprofile.Setup,
		monitormetricalert.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryruleslog.Setup,
		monitorsmartdetectoralertrule.Setup,
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
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
		postgresqlserverkey.Setup,
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
		securitycenterautomation.Setup,
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
		sqlfailovergroup.Setup,
		sqlfirewallrule.Setup,
		sqlserver.Setup,
		sqlvirtualnetworkrule.Setup,
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
		storagesharefile.Setup,
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
		synapsemanagedprivateendpoint.Setup,
		synapseprivatelinkhub.Setup,
		synapseroleassignment.Setup,
		synapsesparkpool.Setup,
		synapsesqlpool.Setup,
		synapseworkspace.Setup,
		templatedeployment.Setup,
		tenanttemplatedeployment.Setup,
		trafficmanagerendpoint.Setup,
		trafficmanagerprofile.Setup,
		userassignedidentity.Setup,
		videoanalyzer.Setup,
		videoanalyzeredgemodule.Setup,
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
		if err := setup(mgr, l, wl); err != nil {
			return err
		}
	}
	return nil
}
