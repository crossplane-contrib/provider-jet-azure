package clients

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane-contrib/provider-tf-azure/apis/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig = "no providerConfigRef provided"
	errGetProviderConfig = "cannot get referenced ProviderConfig"
	errTrackUsage = "cannot track ProviderConfig usage"
	errExtractCredentials = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	errMarshalProviderBlock = "cannot marshal Terraform Azurerm provider block as JSON"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID = "clientId"
	keyAzureClientSecret = "clientSecret"
	keyAzureTenantID = "tenantId"
	// Terraform Provider configuration keys
	keyTerraformSubscriptionID = "subscription_id"
	keyTerraformClientID = "client_id"
	keyTerraformClientSecret = "client_secret"
	keyTerraformTenantID = "tenant_id"
	keyTerraformFeatures = "features"
)

// ProviderConfigBuilder returns provider specific configuration like provider
// credentials used to connect to cloud APIs in the expected form of a Terraform
// provider.
func ProviderConfigBuilder(ctx context.Context, client client.Client, mg xpresource.Managed) ([]byte, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &v1alpha1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := xpresource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}
	azureCreds := map[string]string{}
	if err := json.Unmarshal(data, &azureCreds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}

	tfProviderBlock := map[string]interface{}{}
	tfProviderBlock[keyTerraformSubscriptionID] = azureCreds[keyAzureSubscriptionID]
	tfProviderBlock[keyTerraformTenantID] = azureCreds[keyAzureTenantID]
	tfProviderBlock[keyTerraformClientID] = azureCreds[keyAzureClientID]
	tfProviderBlock[keyTerraformClientSecret] = azureCreds[keyAzureClientSecret]
	tfProviderBlock[keyTerraformFeatures] = struct {}{}
	result, err := json.Marshal(tfProviderBlock)
	return result, errors.Wrap(err, errMarshalProviderBlock)
}
