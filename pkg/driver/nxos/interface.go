package nxos

import (
	"fmt"
)

// GetInterface returns information about an interface.
func (c *Client) GetInterface(name string) (*Config, error) {
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
