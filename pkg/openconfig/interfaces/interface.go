// Package interfaces contains the OpenConfig interfaces data structures.
// This is a manual and incomplete implementation of:
// https://github.com/openconfig/public/blob/master/release/models/interfaces/openconfig-interfaces.yang
package interfaces

// InterfaceType is the type of the interface.
type InterfaceType string

const (
	// InterfaceTypeEthernetCSMACD is the Ethernet CSMA/CD interface type.
	InterfaceTypeEthernetCSMACD InterfaceType = "EthernetCSMACD"
)
