package ssh

import (
	"context"
	"fmt"

	"gopkg.in/ini.v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nicklasfrahm/k3se/pkg/sshx"
	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
)

// Client manages an appliance using SSH.
type Client struct {
	conn *v1alpha1.Connection
	ssh  *sshx.Client
	kube client.Client
	os   *v1alpha1.OSInfo
}

// NewClientFromConnection creates a new client from a connection.
func NewClientFromConnection(client client.Client, conn *v1alpha1.Connection) (*Client, error) {
	return &Client{
		conn: conn,
		kube: client,
	}, nil
}

// Connect connects to the host.
func (c *Client) Connect() error {
	// Fetch credentials from secret.
	namespace := c.conn.Spec.SecretRef.Namespace
	if namespace == "" {
		namespace = c.conn.ObjectMeta.Namespace
	}

	secret := new(corev1.Secret)
	err := c.kube.Get(context.TODO(), client.ObjectKey{
		Namespace: namespace,
		Name:      c.conn.Spec.SecretRef.Name,
	}, secret)
	if err != nil {
		// TODO: Signal user if secret does not exist.
		return err
	}

	var proxy *sshx.Client
	if c.conn.Spec.SSH.ProxyHost != "" {
		proxy, err = sshx.NewClient(&sshx.Config{
			Host:        c.conn.Spec.SSH.ProxyHost,
			Port:        c.conn.Spec.SSH.ProxyPort,
			Fingerprint: c.conn.Spec.SSH.ProxyFingerprint,
			User:        c.conn.Spec.SSH.ProxyUser,
			Key:         string(secret.Data["proxyKey"]),
			Passphrase:  string(secret.Data["proxyPassphrase"]),
			Password:    string(secret.Data["proxyPasswordInsecure"]),
		}, sshx.WithSTFPDisabled())
		if err != nil {
			return err
		}
	}

	c.ssh, err = sshx.NewClient(&sshx.Config{
		Host:        c.conn.Spec.Host,
		Port:        c.conn.Spec.Port,
		Fingerprint: c.conn.Spec.SSH.Fingerprint,
		User:        c.conn.Spec.SSH.User,
		Key:         string(secret.Data["key"]),
		Passphrase:  string(secret.Data["passphrase"]),
		Password:    string(secret.Data["passwordInsecure"]),
	}, sshx.WithProxy(proxy), sshx.WithSTFPDisabled())
	if err != nil {
		return err
	}

	return nil
}

// Disconnect disconnects from the host.
func (c *Client) Disconnect() error {
	return c.ssh.Close()
}

// ProbeOS probes the operating system of the host.
func (c *Client) ProbeOS() (*v1alpha1.OSInfo, error) {
	osReleaseFile := "/etc/os-release"

	session, err := c.ssh.SSH.NewSession()
	if err != nil {
		return nil, err
	}

	osReleaseRaw, err := session.Output(fmt.Sprintf("cat %s", osReleaseFile))
	if err != nil {
		return nil, err
	}
	osRelease, err := ini.Load(osReleaseRaw)
	if err != nil {
		return nil, err
	}

	// Parse the INI formatted file.
	section := osRelease.Section("")
	if section == nil {
		return nil, fmt.Errorf("failed to parse file: %s", osReleaseFile)
	}
	osNameKey := section.Key("NAME")
	if osNameKey == nil {
		return nil, fmt.Errorf("failed to parse OS name from file: %s", osReleaseFile)
	}
	osVersionKey := section.Key("VERSION_ID")
	if osVersionKey == nil {
		return nil, fmt.Errorf("failed to parse OS version from file: %s", osReleaseFile)
	}

	// Store the result internally for later use.
	c.os = &v1alpha1.OSInfo{
		Name:    osNameKey.MustString("Unknown"),
		Version: osVersionKey.MustString("Unknown"),
	}

	return c.os, nil
}
