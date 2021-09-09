# provider-tf-azure

`provider-tf-azure` is a [Crossplane](https://crossplane.io/) provider that
wraps the Terraform CLI and exposes XRM-conformant managed resources for
[Microsoft Azure](https://azure.microsoft.com/).


## Developing

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```
