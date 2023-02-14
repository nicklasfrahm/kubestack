package nxos

import (
	"fmt"
	"strconv"
	"strings"

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
		Header: normalizeInterfaceNames(fmt.Sprintf("interface %s", iface.Name)),
		Lines:  make(map[string]bool),
	}

	if iface.Spec.MTU != 0 {
		section.Lines[fmt.Sprintf("mtu %d", iface.Spec.MTU)] = true
	}

	if iface.Spec.Description != "" {
		section.Lines[fmt.Sprintf("description %s", iface.Spec.Description)] = true
	}

	if !iface.Spec.Enabled {
		section.Lines["shutdown"] = true
	}

	if iface.Spec.Ethernet.AutoNegotiate {
		section.Lines["negotiate auto"] = true
	}

	// TODO: Not supported on n3k platform.
	//       How should we handle this?
	if iface.Spec.Ethernet.EnableFlowControl {
		section.Lines["flowcontrol receive on"] = true
		section.Lines["flowcontrol send on"] = true
	}

	if iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode == v1alpha1.VLANModeTrunk {
		section.Lines["switchport mode trunk"] = true

		if iface.Spec.Ethernet.SwitchedVLAN.NativeVLAN != 0 {
			section.Lines[fmt.Sprintf("switchport trunk native vlan %d", iface.Spec.Ethernet.SwitchedVLAN.NativeVLAN)] = true
		}

		for _, vlan := range iface.Spec.Ethernet.SwitchedVLAN.TrunkVLANs {
			section.Lines[fmt.Sprintf("switchport trunk allowed vlan %d", vlan)] = true
		}
	}

	// The user must explicitly configure the interface mode to be `Access`.
	// Skipping this will prevent the access VLAN ID from being configured.
	if iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode == v1alpha1.VLANModeAccess {
		section.Lines["switchport mode access"] = true

		if iface.Spec.Ethernet.SwitchedVLAN.AccessVLAN != 0 {
			section.Lines[fmt.Sprintf("switchport access vlan %d", iface.Spec.Ethernet.SwitchedVLAN.AccessVLAN)] = true
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
	if len(config.Sections) == 0 {
		return nil, fmt.Errorf("failed to marshal config: no sections found")
	}

	section := config.Sections[0]
	interfaceName := strings.TrimPrefix(normalizeInterfaceNames(section.Header), "interface ")

	iface := &v1alpha1.Interface{
		Spec: v1alpha1.InterfaceSpec{
			Selector: v1alpha1.InterfaceSelector{
				Name: interfaceName,
			},
		},
	}

	for line := range section.Lines {

		if strings.HasPrefix(line, "mtu ") {
			mtu, err := strconv.Atoi(strings.TrimPrefix(line, "mtu "))
			if err != nil {
				return nil, fmt.Errorf("failed to parse MTU: %w", err)
			}
			iface.Spec.MTU = mtu
			continue
		}

		if strings.HasPrefix(line, "description ") {
			iface.Spec.Description = strings.TrimPrefix(line, "description ")
			continue
		}

		if strings.HasPrefix(line, "shutdown") {
			iface.Spec.Enabled = false
			continue
		}

		if strings.HasPrefix(line, "negotiate auto") {
			iface.Spec.Ethernet.AutoNegotiate = true
			continue
		}

		if strings.HasPrefix(line, "flowcontrol receive on") {
			iface.Spec.Ethernet.EnableFlowControl = true
			continue
		}

		if strings.HasPrefix(line, "flowcontrol send on") {
			iface.Spec.Ethernet.EnableFlowControl = true
			continue
		}

		if strings.HasPrefix(line, "switchport mode trunk") {
			iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode = v1alpha1.VLANModeTrunk
			continue
		}

		if strings.HasPrefix(line, "switchport trunk native vlan ") {
			iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode = v1alpha1.VLANModeTrunk
			vlan, err := strconv.Atoi(strings.TrimPrefix(line, "switchport trunk native vlan "))
			if err != nil {
				return nil, err
			}
			iface.Spec.Ethernet.SwitchedVLAN.NativeVLAN = vlan
			continue
		}

		if strings.HasPrefix(line, "switchport trunk allowed vlan ") {
			iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode = v1alpha1.VLANModeTrunk
			vlan, err := strconv.Atoi(strings.TrimPrefix(line, "switchport trunk allowed vlan "))
			if err != nil {
				return nil, fmt.Errorf("failed to parse trunk VLAN ID: %w", err)
			}
			iface.Spec.Ethernet.SwitchedVLAN.TrunkVLANs = append(iface.Spec.Ethernet.SwitchedVLAN.TrunkVLANs, vlan)
			continue
		}

		if strings.HasPrefix(line, "switchport access vlan") {
			iface.Spec.Ethernet.SwitchedVLAN.InterfaceMode = v1alpha1.VLANModeAccess
			vlan, err := strconv.Atoi(strings.TrimPrefix(line, "switchport access vlan "))
			if err != nil {
				return nil, fmt.Errorf("failed to parse access VLAN ID: %w", err)
			}
			iface.Spec.Ethernet.SwitchedVLAN.AccessVLAN = vlan
			continue
		}
	}

	return iface, nil
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
	desired, err := UnmarshalInterface(iface)
	if err != nil {
		return nil, err
	}
	_ = desired

	current, err := c.getInterface(iface.Spec.Selector.Name)
	if err != nil {
		return nil, err
	}
	_ = current

	// TODO: Diff target and current configuration.

	// TODO: Render commands to apply diff.

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

	desired := Config{
		Sections: []Section{
			{
				Header: fmt.Sprintf("interface %s", iface.Spec.Selector.Name),
			},
		},
	}
	_ = desired

	// TODO: Technically speaking, this should be a low-hanging fruit,
	//       because the deletion would only clear all current configuration.

	return fmt.Errorf("method not implemented: DeleteInterface")
}

// normalizeInterfaceNames normalizes interface names.
// This function lowercases the interface name and
// replaces "Ethernet" with "eth".
func normalizeInterfaceNames(raw string) string {
	if !strings.HasPrefix(raw, "interface") {
		return raw
	}

	raw = strings.ToLower(raw)
	raw = strings.ReplaceAll(raw, "ethernet", "eth")

	return raw
}
