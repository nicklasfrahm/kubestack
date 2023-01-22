# SSH

This section describes how to connect to appliances using SSH.

## Secret

Below you may find an example of a `Secret` for an SSH connection using all possible keys. The usage of an encrypted private key and a jump host is optional.

```yaml title="ssh-secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: ssh-credentials
type: Opaque
stringData:
  # Passphrase if the private key is encrypted.
  passphrase: dont-check-this-into-your-repo-please
  # Password for authentication.
  # This is only used if the private key is not provided.
  passwordInsecure: possible-but-not-recommended
  # Private key for authentication. This takes precedence over the password.
  key: |
    -----BEGIN OPENSSH PRIVATE KEY-----
    REDACTED
    -----END OPENSSH PRIVATE KEY-----
  # Passphrase for an encrypted private key while using a jump host.
  proxyPassphrase: dont-check-this-into-your-repo-please
  # Password for authentication while using a jump host.
  # This is only used if the private key is not provided.
  proxyPasswordInsecure: possible-but-not-recommended
  # Private key for authentication while using a jump host.
  # This takes precedence over the password.
  proxyKey: |
    -----BEGIN OPENSSH PRIVATE KEY-----
    REDACTED
    -----END OPENSSH PRIVATE KEY-----
```

## Connection

Below, you may find a simple example where the controller will connect directly to the appliance.

```yaml title="distswitch00.yaml"
--8<-- "config/samples/kubestack_v1alpha1_connection_distswitch00.yaml"
```

In more complex setup, you many need to use a jump host to connect to your appliance. This can be done as well.

```yaml title="alfa.yaml"
--8<-- "config/samples/kubestack_v1alpha1_connection_alfa.yaml"
```
