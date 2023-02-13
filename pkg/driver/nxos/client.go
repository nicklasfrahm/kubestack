package nxos

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

const (
	// vshExecutable is the path to the VSH executable.
	vshExecutable = "/isan/bin/vsh.bin"
)

// Client is an NX-OS client.
type Client struct {
	SSH *ssh.Client
}

// NewClient creates a new client.
func NewClient(opts ...Option) (*Client, error) {
	client := &Client{}
	if err := client.Apply(opts...); err != nil {
		return nil, err
	}

	if client.SSH == nil {
		return nil, fmt.Errorf("missing required option: WithSSH()")
	}

	return client, nil
}

// runVSHCommand runs a command in the VSH on the host.
func (c *Client) runVSHCommand(cmd string) (string, error) {
	session, err := c.SSH.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	out, err := session.Output(fmt.Sprintf("%s -c '%s'", vshExecutable, cmd))
	if err != nil {
		return "", err
	}

	return string(out), nil
}
