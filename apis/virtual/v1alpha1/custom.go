package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/config"

func virtualNetworkExternalNameConfigure(_ map[string]interface{}, _ string) {}

func init() {
	config.Store.SetForResource("azurerm_virtual_network", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath:  "virtualNetworkExternalNameConfigure",
			DisableNameInitializer: true,
		},
		UseAsync: true,
	})
}
