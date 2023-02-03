# Interfaces

The `Interface` custom resource allows you to manage network interfaces on your network devices. All interfaces must reference a `Connection`. For more information about how to configure a `Connection`, please refer to the [management](../management.md) section.

You may get a list of network interfaces by running the following command.

```shell
kubectl get if
```

<!-- TODO: Add command output. -->

## VLAN

Virtual local area networks (VLANs) are used to segment a network into multiple broadcast domains. They allow you to build multiple networks using a single physical network. For more information about VLANs, please refer to the [VLAN Wikipedia article](https://en.wikipedia.org/wiki/Virtual_LAN).

### Access port

An access port is a link that is configured to only allow traffic for a single VLAN. This is useful to connect an end-user device to a specific VLAN.

```yaml title="distswitch00-eth1-1.yaml"
--8<-- "config/samples/kubestack_v1alpha1_interface_distswitch00-eth1-1.yaml"
```

### Trunk port

A trunk port is a link that is configured to allow traffic for multiple VLANs. This is useful to connect two VLAN-aware switches together.

```yaml title="distswitch00-eth1-3.yaml"
--8<-- "config/samples/kubestack_v1alpha1_interface_distswitch00-eth1-3.yaml"
```
