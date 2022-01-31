module github.com/crossplane-contrib/provider-jet-azure

go 1.16

require (
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/terrajet v0.4.0-rc.0.0.20220128111246-e5aaa1790fe6
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	// github.com/hashicorp/terraform-provider-azurerm is replaced with  ./.work/.azurerm
	// and the dependency version is specified in the Makefile via make variable AZURERM_REFSPEC
	github.com/hashicorp/terraform-provider-azurerm v0.0.0
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

replace github.com/hashicorp/terraform-provider-azurerm => ./.work/.azurerm
