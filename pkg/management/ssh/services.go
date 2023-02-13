package ssh

import (
	"fmt"

	"github.com/nicklasfrahm/kubestack/pkg/driver/nxos"
	"github.com/nicklasfrahm/kubestack/pkg/management/common"
)

// Interface attempts to load a service to manage interfaces.
// The implementation depends on the probed operating system.
func (c *Client) Interface() (common.InterfaceService, error) {
	if c.os.Name == common.OSNexus {
		return nxos.NewClient(nxos.WithSSH(c.ssh.SSH))
	}

	return nil, fmt.Errorf("failed to load service for operating system: %s", c.os.Name)
}
