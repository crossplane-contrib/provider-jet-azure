package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/config"

func kubernetesClusterExternalNameConfigure(_ map[string]interface{}, _ string) {}

func init() {
	config.Store.SetForResource("azurerm_kubernetes_cluster", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath:  "kubernetesClusterExternalNameConfigure",
			DisableNameInitializer: true,
		},
		LateInitializer: config.LateInitializer{
			IgnoredFields: []string{"KubeletIdentity"},
		},
		UseAsync: true,
	})
}
