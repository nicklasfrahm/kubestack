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

	// Interface loads the interface service of the appliance.
	Interface() (InterfaceService, error)
}

// InterfaceService provides methods for managing interfaces.
type InterfaceService interface {
	// ListInterfaces lists all interfaces.
	ListInterfaces(*v1alpha1.InterfaceSelector) (*v1alpha1.InterfaceList, error)
	// GetInterface reads an interface.
	GetInterface(*v1alpha1.InterfaceSelector) (*v1alpha1.Interface, error)
	// CreateInterface creates an interface.
	CreateInterface(*v1alpha1.Interface) (*v1alpha1.Interface, error)
	// UpdateInterface updates an interface.
	UpdateInterface(*v1alpha1.Interface) (*v1alpha1.Interface, error)
	// DeleteInterface deletes an interface.
	DeleteInterface(*v1alpha1.Interface) error
}
