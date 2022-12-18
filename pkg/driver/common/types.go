package common

type Driver interface {
	// Connect establishes a connection to the host.
	Connect() error
	// Ping checks if the connection is still alive.
	Ping() error
	// Disconnect closes the connection to the host.
	Disconnect() error
}
