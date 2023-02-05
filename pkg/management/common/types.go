package common

import (
	v1alpha1 "github.com/nicklasfrahm/kubestack/api/v1alpha1"
)

// ClientFactory is a function that creates a new client.
type ClientFactory func(*v1alpha1.Connection, ...Option) (Client, error)

// Client is the interface for a client.
type Client interface {
	// Connect connects to the host.
	Connect() error
	// Disconnect disconnects from the host.
	Disconnect() error

	// ProbeOS probes the operating system of the host.
	ProbeOS() (*v1alpha1.OSInfo, error)

	// TODO: ReconcileInterface reconciles an interface.
	// ReconcileInterface(iface *v1alpha1.Interface) error
}
