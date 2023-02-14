/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ocif "github.com/nicklasfrahm/kubestack/pkg/openconfig/interfaces"
)

// VLANMode is the VLAN mode.
type VLANMode string

const (
	// VLANModeAccess is the access VLAN mode. Sometimes also referred
	// to as `untagged` as the frames do not have the 802.1Q tag.
	VLANModeAccess VLANMode = "Access"
	// VLANModeTrunk is the trunk VLAN mode. Sometimes also referred
	// to as `tagged` as the frames have the 802.1Q tag.
	VLANModeTrunk VLANMode = "Trunk"
)

// ConnectionReference is the reference to the connection
// that the interface is associated with.
type ConnectionReference struct {
	// Namespace of the connection.
	Namespace string `json:"namespace,omitempty"`
	// Name of the connection.
	//+kubebuilder:validation:Required
	Name string `json:"name"`
}

// InterfaceSelector is the selector for the interface.
type InterfaceSelector struct {
	// Name is the name of the interface.
	Name string `json:"name,omitempty"`
	// MACAddress is the MAC address of the interface.
	//+kubebuilder:validation:Pattern=`^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$`
	MACAddress string `json:"macAddress,omitempty"`
}

// IPv4IP is the IPv4 address configuration of an interface.
type IPv4IP struct {
	// Address is the IPv4 address.
	IP ocif.IPv4Address `json:"ip,omitempty"`
	// PrefixLength is the prefix length of the IPv4 address.
	PrefixLength uint8 `json:"prefixLength,omitempty"`
}

// IPv4Neighbor is an IPv4 neighbor.
type IPv4Neighbor struct {
	// IP is the IPv4 address of the neighbor.
	IP ocif.IPv4Address `json:"ip,omitempty"`
	// LinkLayerAddress is the link layer address of the neighbor.
	// TODO: Implement physical address type and use it here.
	LinkLayerAddress string `json:"linkLayerAddress,omitempty"`
}

// IPv4 is the IPv4 configuration of an interface.
type IPv4 struct {
	// Addresses is the list of IPv4 addresses assigned to the interface.
	Addresses []IPv4IP `json:"addresses,omitempty"`

	// TODO: Add proxy ARP.
	// ProxyARP enables proxy ARP on the interface.
	// ProxyARP ProxyARP `json:"proxyARP,omitempty"`

	// Neighbors is the list of IPv4 neighbors.
	Neighbors []IPv4Neighbor `json:"neighbors,omitempty"`

	// TODO: Add enabled flag.
	// Enabled controls the state of the IPv4 configuration.
	// Enabled bool

	// TODO: Add MTU.
	// MTU is the maximum transmission unit of the interface.
	// MTU uint16

	// DHCPClient enables DHCP client on the interface.
	DHCPClient bool `json:"dhcpClient,omitempty"`

	// TODO: Add support for IPv4 unnumbered.

	// TODO: Add support for IPv6.
}

// SwitchedVLANSpec defines the desired state of a switched VLAN.
type SwitchedVLANSpec struct {
	// InterfaceMode is the VLAN mode.
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Enum=Access;Trunk
	InterfaceMode VLANMode `json:"interfaceMode"`
	// NativeVLAN is the native VLAN.
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=4094
	NativeVLAN int `json:"nativeVLAN,omitempty"`
	// AccessVLAN is the access VLAN.
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=4094
	AccessVLAN int `json:"accessVLAN,omitempty"`
	// TrunkVLANs is the list of trunk VLANs.
	TrunkVLANs []int `json:"trunkVLANs,omitempty"`
}

// RoutedVLANSpec defines the desired state of a routed VLAN.
type RoutedVLANSpec struct {
	// VLANID is the VLAN ID of the routed VLAN.
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=4094
	VLAN int `json:"vlan,omitempty"`

	// IPv4 configures the IPv4 address of the interface.
	IPv4 IPv4 `json:"ipv4,omitempty"`
}

// EthernetSpec defines the desired state of Ethernet.
type EthernetSpec struct {
	// MACAddress allows to override the system-assigned address of the interface.
	//+kubebuilder:validation:Pattern=`^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$`
	MACAddress string `json:"macAddress,omitempty"`

	// AutoNegotiate specifies whether auto-negotiation is enabled.
	AutoNegotiate bool `json:"autoNegotiate,omitempty"`

	// TODO: Add support for standalone link training.
	// StandaloneLinkTraining specifies whether standalone link training is enabled.
	// StandaloneLinkTraining bool `json:"standaloneLinkTraining,omitempty"`

	// Duplex mode configures the duplex mode being advertised by the interface.
	// When auto-negotiate is TRUE, this optionally sets the
	// duplex mode that will be advertised to the peer. If
	// unspecified, the interface should negotiate the duplex mode
	// directly (typically full-duplex). When auto-negotiate is
	// FALSE, this sets the duplex mode on the interface directly.
	//+kubebuilder:validation:Enum=Half;Full
	DuplexMode ocif.DuplexMode `json:"duplexMode,omitempty"`

	// PortSpeed configures the port speed being advertised by the interface.
	// When auto-negotiate is TRUE, this optionally sets the
	// port-speed mode that will be advertised to the peer for
	// negotiation. If unspecified, it is expected that the
	// interface will select the highest speed available based on
	// negotiation. When auto-negotiate is set to FALSE, sets the
	// link speed to a fixed value.
	PortSpeed ocif.EthernetSpeed `json:"portSpeed,omitempty"`

	// EnableFlowControl specifies whether flow control is enabled.
	// If not specified, the default value is `false`.
	EnableFlowControl bool `json:"enableFlowControl,omitempty"`

	// FECMode configures the forward error correction mode.
	// FECMode ocif.FECMode `json:"fecMode,omitempty"`

	// TODO: Add support for aggregate interfaces.
	// AggregateID specifies the ID of the aggregate to which the interface belongs.
	// AggregateID int `json:"aggregateId,omitempty"`

	// SwitchedVLAN specifies the VLAN configuration of the interface.
	SwitchedVLAN SwitchedVLANSpec `json:"switchedVLAN,omitempty"`

	// TODO: Add support for PoE.

	// TODO: Add support for dot1x.

	// TODO: Add support for authenticated sessions.
}

// InterfaceSpec defines the desired state of Interface
type InterfaceSpec struct {
	// ConnectionRef is the reference to the connection that this interface is associated with.
	//+kubebuilder:validation:Required
	ConnectionRef ConnectionReference `json:"connectionRef,omitempty"`

	// Selector allows to select the interface based on certain criteria.
	//+kubebuilder:validation:Required
	Selector InterfaceSelector `json:"selector,omitempty"`

	// Name allows to override the interface name.
	Name string `json:"name,omitempty"`

	// Type specifies the type of the interface.
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Enum=EthernetCSMACD
	Type ocif.InterfaceType `json:"type"`

	// MTU specifies the maximum transmission unit of the interface.
	MTU int `json:"mtu,omitempty"`

	// TODO: Add loopback mode. Question though: Is this not implied-ish by the name?
	// LoopbackMode specifies whether the interface is in loopback mode.
	// LoopbackMode bool `json:"loopbackMode,omitempty"`

	// Description provides human-readable information about the interface.
	//+kubebuilder:validation:Pattern=`^[a-zA-Z0-9.,-_ ]{1,255}$`
	Description string `json:"description,omitempty"`

	// Enabled specifies whether the interface is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// TODO: Add support for Q-in-Q. Prerequisite: Understand Q-in-Q, LoL.
	// Reference: https://github.com/openconfig/public/blob/master/release/models/vlan/openconfig-vlan-types.yang#L87
	// Cisco: https://www.cisco.com/en/US/docs/ios/lanswitch/configuration/guide/lsw_ieee_802.1q.pdf
	// TPID configures the tag protocol identifier (TPID) for the interface
	// that is accepted on the VLAN.
	// TPID TPID `json:"tpid,omitempty"`

	// TODO: Add support for forwarding viable flag.
	// ForwardingViable indicates whether the interface may route traffic or not. L3?
	// ForwardingViable bool `json:"forwardingViable,omitempty"`

	// TODO: Add support for hold time.

	// Ethernet configures the Ethernet-specific properties of the interface.
	Ethernet EthernetSpec `json:"ethernet,omitempty"`

	// TODO: Add support for aggregation.

	// TODO: Add support for subinterfaces.

	// RoutedVLAN specifies the VLAN configuration of the interface.
	RoutedVLAN RoutedVLANSpec `json:"routedVLAN,omitempty"`
}

// InterfaceStatus defines the observed state of Interface
type InterfaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:categories={net,network,networking},shortName=if,path=interfaces,singular=interface
//+kubebuilder:printcolumn:name="Connection-Name",type=string,JSONPath=`.spec.connectionRef.name`
//+kubebuilder:printcolumn:name="Selector-Name",type=string,JSONPath=`.spec.selector.name`
//+kubebuilder:printcolumn:name="VLAN-Mode",type=string,JSONPath=`.spec.ethernet.switchedVLAN.interfaceMode`

// Interface is the Schema for the interfaces API
type Interface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InterfaceList contains a list of Interface
type InterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Interface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Interface{}, &InterfaceList{})
}
