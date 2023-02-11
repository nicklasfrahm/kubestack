package common

import (
	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
)

// ClientFactory is a function that creates a new client.
type ClientFactory func(*v1alpha1.Connection, ...Option) (Client, error)

// Client is the interface for a client.
type Client interface {
	// Connect connects to the host.
	Connect() error
	// Disconnect disconnects from the host.
	Disconnect() error

	// OS returns information about the operating system the host.
	OS() *v1alpha1.OSInfo
}
