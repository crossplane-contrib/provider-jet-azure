package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane-contrib/provider-jet-azure/apis/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
	keyAzureTenantID       = "tenantId"
	// Terraform Provider configuration keys
	keyTerraformFeatures = "features"
	// environment variable names for storing credentials
	envClientID       = "ARM_CLIENT_ID"
	envClientSecret   = "ARM_CLIENT_SECRET"
	envSubscriptionID = "ARM_SUBSCRIPTION_ID"
	envTenantID       = "ARM_TENANT_ID"
	// Terraform configuration file keys
	keyTerraformSubscriptionID = "subscription_id"

	fmtEnvVar = "%s=%s"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := xpresource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		azureCreds := map[string]string{}
		if err := json.Unmarshal(data, &azureCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures:       struct{}{},
			keyTerraformSubscriptionID: azureCreds[keyAzureSubscriptionID],
		}
		// set credentials environment
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envSubscriptionID, azureCreds[keyAzureSubscriptionID]),
			fmt.Sprintf(fmtEnvVar, envTenantID, azureCreds[keyAzureTenantID]),
			fmt.Sprintf(fmtEnvVar, envClientID, azureCreds[keyAzureClientID]),
			fmt.Sprintf(fmtEnvVar, envClientSecret, azureCreds[keyAzureClientSecret]),
		}
		return ps, err
	}
}
