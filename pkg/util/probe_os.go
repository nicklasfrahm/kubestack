package util

import (
	"fmt"
	"strconv"

	"gopkg.in/ini.v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/nicklasfrahm/k3se/pkg/sshx"
	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
)

// osProbe is a function that probes the operating system of the system.
type osProbe func(*corev1.Secret) (*v1alpha1.OSInfo, error)

// ProbeOS probes the operating system of the system.
func ProbeOS(conn *v1alpha1.Connection, secret *corev1.Secret) (*v1alpha1.OSInfo, error) {
	probes := map[v1alpha1.Protocol]osProbe{
		v1alpha1.ProtocolSSH: probeOSViaSSH,
	}

	probe := probes[conn.Spec.Protocol]

	if probe == nil {
		return nil, fmt.Errorf("unknown protocol: %s", conn.Spec.Protocol)
	}

	return probe(secret)
}

// probeOSViaSSH probes the operating system of the system via SSH.
func probeOSViaSSH(secret *corev1.Secret) (*v1alpha1.OSInfo, error) {
	// TODO: Allow usage of SSH proxy or bastion host.
	port, err := strconv.Atoi(string(secret.Data["port"]))
	if err != nil {
		port = 22
	}

	config := &sshx.Config{
		Host:        string(secret.Data["host"]),
		Port:        port,
		User:        string(secret.Data["user"]),
		Key:         string(secret.Data["key"]),
		Fingerprint: string(secret.Data["fingerprint"]),
	}

	client, err := sshx.NewClient(config)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	session, err := client.SSH.NewSession()
	if err != nil {
		return nil, err
	}

	osReleaseRaw, err := session.Output("cat /etc/os-release")
	if err != nil {
		return nil, err
	}

	osRelease, err := ini.Load(osReleaseRaw)
	if err != nil {
		return nil, err
	}

	section := osRelease.Section("")
	if section == nil {
		return nil, fmt.Errorf("failed to parse file: /etc/os-release")
	}

	osNameKey := section.Key("NAME")
	if osNameKey == nil {
		return nil, fmt.Errorf("failed to parse OS name from file: /etc/os-release")
	}

	osVersionKey := section.Key("VERSION_ID")
	if osVersionKey == nil {
		return nil, fmt.Errorf("failed to parse OS version from file: /etc/os-release")
	}

	return &v1alpha1.OSInfo{
		Name:    osNameKey.MustString("Unknown"),
		Version: osVersionKey.MustString("Unknown"),
	}, nil
}
