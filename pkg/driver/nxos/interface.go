package nxos

import (
	"fmt"

	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
)

var (
	// ErrMissingSelectorName is returned when the interface selector name is missing.
	ErrMissingSelectorName = fmt.Errorf("missing selector: .spec.selector.name")
)

// ValidateInterface validates the driver-specific interface configuration.
func ValidateInterface(iface *v1alpha1.Interface) error {
	if iface.Spec.Selector.Name == "" {
		return ErrMissingSelectorName
	}

	if iface.Spec.Name != "" {
		return fmt.Errorf("unsupported field: .spec.name")
	}

	// TODO: Validate interface configuration for obvious errors, e.g. trunk / access conflict.

	return nil
}

// UnmarshalInterface unmarshals the OpenConfig interface configuration into the driver-specific format.
func UnmarshalInterface(iface *v1alpha1.Interface) (*Config, error) {
	section := Section{
		Header: fmt.Sprintf("interface %s", iface.Name),
		Lines:  make(map[string]bool),
	}

	if iface.Spec.MTU != 0 {
		section.Lines[fmt.Sprintf("mtu %d", iface.Spec.MTU)] = true
	}

	if iface.Spec.Description != "" {
		section.Lines[fmt.Sprintf("description %s", iface.Spec.Description)] = true
	}

	if iface.Spec.Enabled {
		section.Lines["no shutdown"] = true
	} else {
		section.Lines["shutdown"] = true
	}

	if iface.Spec.Ethernet != nil {
		if iface.Spec.Ethernet.AutoNegotiate {
			section.Lines["negotiate auto"] = true
		}

		if iface.Spec.Ethernet.EnableFlowControl {
			section.Lines["flowcontrol receive on"] = true
			section.Lines["flowcontrol send on"] = true
		} else {
			section.Lines["flowcontrol receive off"] = true
			section.Lines["flowcontrol send off"] = true
		}

		if iface.Spec.Ethernet.SwitchedVLAN != nil {
			if iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode == v1alpha1.VLANModeTrunk {
				section.Lines["switchport mode trunk"] = true

				if iface.Spec.Ethernet.SwitchedVLAN.NativeVLAN != 0 {
					section.Lines[fmt.Sprintf("switchport trunk native vlan %d", iface.Spec.Ethernet.SwitchedVLAN.NativeVLAN)] = true
				}

				for _, vlan := range iface.Spec.Ethernet.SwitchedVLAN.TrunkVLANs {
					section.Lines[fmt.Sprintf("switchport trunk allowed vlan %d", vlan)] = true
				}
			}

			if iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode == v1alpha1.VLANModeAccess {
				section.Lines["switchport mode access"] = true

				if iface.Spec.Ethernet.SwitchedVLAN.AccessVLAN != 0 {
					section.Lines[fmt.Sprintf("switchport access vlan %d", iface.Spec.Ethernet.SwitchedVLAN.AccessVLAN)] = true
				}
			}
		}
	}

	return &Config{
		Sections: []Section{
			section,
		},
	}, nil
}

// MarshalInterface marshals the driver-specific interface configuration into the OpenConfig format.
func MarshalInterface(config *Config) (*v1alpha1.Interface, error) {

	// TODO: Implement this method.

	return nil, fmt.Errorf("method not implemented: MarshalInterface")
}

// ListInterfaces lists all interfaces.
func (c *Client) ListInterfaces(selector *v1alpha1.InterfaceSelector) (*v1alpha1.InterfaceList, error) {

	// TODO: Implement this method.

	return nil, fmt.Errorf("method not implemented: ListInterfaces")
}

// GetInterface reads an interface.
func (c *Client) GetInterface(selector *v1alpha1.InterfaceSelector) (*v1alpha1.Interface, error) {
	config, err := c.getInterface(selector.Name)
	if err != nil {
		return nil, err
	}

	iface, err := MarshalInterface(config)
	if err != nil {
		return nil, err
	}

	return iface, nil
}

// getInterface returns information about an interface.
func (c *Client) getInterface(name string) (*Config, error) {
	raw, err := c.runVSHCommand(fmt.Sprintf("show running-config interface %s", name))
	if err != nil {
		return nil, err
	}

	config, err := Parse(raw)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// CreateInterface creates an interface.
func (c *Client) CreateInterface(iface *v1alpha1.Interface) (*v1alpha1.Interface, error) {
	// NX-OS does not support the creation of interfaces. Instead, it will just write the
	// corresponding configuration, which is effectively an update.
	return c.UpdateInterface(iface)
}

// UpdateInterface updates an interface.
func (c *Client) UpdateInterface(iface *v1alpha1.Interface) (*v1alpha1.Interface, error) {
	if err := ValidateInterface(iface); err != nil {
		return nil, err
	}

	// Convert OpenConfig interface configuration into NX-OS specific format.
	config, err := UnmarshalInterface(iface)
	if err != nil {
		return nil, err
	}
	_ = config

	// TODO: Read current configuration in NX-OS specific format.

	// TODO: Diff target and current configuration.

	// TODO: Render configuration.

	// TODO: Apply surgical configuration update.

	return nil, fmt.Errorf("method not implemented: UpdateInterface")
}

// DeleteInterface deletes an interface.
func (c *Client) DeleteInterface(iface *v1alpha1.Interface) error {
	// Skip the detailed validation, but we
	// still need to know what to delete.
	if iface.Spec.Selector.Name == "" {
		return ErrMissingSelectorName
	}

	// TODO: Technically speaking, this should be a low-hanging fruit,
	//       because the deletion would only clear all current configuration.

	return fmt.Errorf("method not implemented: DeleteInterface")
}
