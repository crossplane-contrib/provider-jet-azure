package clients

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// error messages
	errNoProviderConfig = "no providerConfigRef provided"
	errGetProviderConfig = "cannot get referenced ProviderConfig"
	errTrackUsage = "cannot track ProviderConfig usage"
	errExtractCredentials = "cannot extract credentials"
)

// ProviderConfigBuilder returns provider specific configuration like provider
// credentials used to connect to cloud APIs in the expected form of a Terraform
// provider.
func ProviderConfigBuilder(ctx context.Context, client client.Client, mg xpresource.Managed) ([]byte, error) {
	return []byte(`{"features":{}}`), nil
}

