package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/config"

func postgresqlServerExternalNameConfigure(_ map[string]interface{}, _ string) {}

func init() {
	config.Store.SetForResource("azurerm_postgresql_server", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath:  "postgresqlServerExternalNameConfigure",
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"SslEnforcement", "StorageProfile"},
		},
		UseAsync: true,
	})
}
