# Kubestack ♻️

Kubestack is a infrastructure orchestrator built on top of Kubernetes. The goal is to provide APIs to build private clouds using Kubernetes as a control plane and API. You may find more information in the documentation at [kubestack.nicklasfrahm.dev][docs-kubestack].

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Create a secret for the connection credentials:

```yaml
# Location: config/samples/secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  labels:
    app.kubernetes.io/name: connection
    app.kubernetes.io/instance: delta
    app.kubernetes.io/part-of: kubestack
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: kubestack
  name: connection-delta
  namespace: default
type: Opaque
stringData:
  host: xx.xx.xx.xx
  user: xxxx
  key: |
    -----BEGIN OPENSSH PRIVATE KEY-----
    REDACTED
    -----END OPENSSH PRIVATE KEY-----
  fingerprint: SHA256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

1. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/kubestack:tag
```

1. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/kubestack:tag
```

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing

// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/)
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster

### Test It Out

1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

This projects is licensed under the terms of the [MIT license](./LICENSE.md).

[docs-kubestack]: https://kubestack.nicklasfrahm.dev
