# provider-jet-azure

`provider-jet-azure` is a [Crossplane](https://crossplane.io/) provider that is
built using [Terrajet](https://github.com/crossplane-contrib/terrajet) code
generation tools and exposes XRM-conformant managed resources for
[Microsoft Azure](https://azure.microsoft.com/).

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-jet-azure/releases):
```
kubectl crossplane install provider crossplane/provider-jet-azure:${release-version}
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-azure).

## Contributing

Please refer to the [Adding New Resources](/docs/adding-resources.md) guide.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane/provider-jet-azure/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-jet-azure is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-azure adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-azure is under the Apache 2.0 license.
