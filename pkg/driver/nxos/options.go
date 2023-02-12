package nxos

import (
	"golang.org/x/crypto/ssh"
)

// Option is a function that configures a client.
type Option func(*Client) error

// WithSSH uses SSH as the transport protocol.
func WithSSH(ssh *ssh.Client) Option {
	return func(c *Client) error {
		c.SSH = ssh
		return nil
	}
}

// Apply applies the options to the client.
func (c *Client) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}
