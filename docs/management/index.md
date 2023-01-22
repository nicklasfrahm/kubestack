# Overview

This section describes how to connect to existing appliances to manage them.

## Protocols

To configure the remote management of an appliance, you need to create a `Connection`, which defines a management protocol and a reference to a `Secret` that contains the credentials to connect to the appliance. As the authentication methods for each protocol vary, so do the supported keys in the `Secret` based on the mangement protocol. Learn more about the configuration of a `Connection` for a supported management protocol and the structure of the corresponding `Secret` by following the links below.

- [`SSH`](./ssh.md)

## Troubleshooting

If you are having trouble connecting to your appliance, inspecting the event log may provide useful information.

```bash
kubectl get event \
  --field-selector involvedObject.kind=Connection \
  --field-selector involvedObject.name=alfa
```

Below you may find an example of the event log for a `Connection` that is not able to connect to an appliance.

```text
LAST SEEN   TYPE      REASON             OBJECT            MESSAGE
50m         Warning   ConnectionFailed   connection/alfa   dial tcp: lookup alfa.nicklasfrahm.dev: no such host
13m         Warning   ConnectionFailed   connection/alfa   Secret "connection-alfa" not found
10m         Normal    OSProbed           connection/alfa   OS information probed successfully.
```
