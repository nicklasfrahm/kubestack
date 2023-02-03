# Overview

This section describes how to connect to existing appliances to manage them.

## Protocols

To configure the remote management of an appliance, you need to create a `Connection`, which defines a management protocol and a reference to a `Secret` that contains the credentials to connect to the appliance. As the authentication methods for each protocol vary, so do the supported keys in the `Secret` based on the mangement protocol. Learn more about the configuration of a `Connection` for a supported management protocol and the structure of the corresponding `Secret` by following the links below.

- [`SSH`: Secure Shell protocol](./management/ssh.md)

## Verification

If a `Connection` is successfully created, the controller will probe for the operating system of the appliance. You may verify this by running the following command.

```shell
kubectl get conn
```

```text
NAME           PROTOCOL   OS-NAME   OS-VERSION
alfa           SSH        Ubuntu    22.04
distswitch00   SSH        Nexus     9.3(10)I9
```

## Troubleshooting

If you are having trouble connecting to your appliance, inspecting the event log may provide useful information.

```bash
kubectl get event \
  --field-selector involvedObject.kind=Connection \
  --field-selector involvedObject.name=alfa
```

```text
LAST SEEN   TYPE      REASON             OBJECT            MESSAGE
50m         Warning   ConnectionFailed   connection/alfa   dial tcp: lookup alfa.nicklasfrahm.dev: no such host
13m         Warning   ConnectionFailed   connection/alfa   Secret "connection-alfa" not found
10m         Normal    OSProbed           connection/alfa   OS information probed successfully.
```
