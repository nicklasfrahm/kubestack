package interfaces

// DuplexMode is the duplex mode of the interface.
type DuplexMode string

const (
	// DuplexModeHalf is the half duplex mode.
	DuplexModeHalf DuplexMode = "Half"
	// DuplexModeFull is the full duplex mode.
	DuplexModeFull DuplexMode = "Full"
)

// EthernetSpeed is the speed of the interface.
type EthernetSpeed string

const (
	// EthernetSpeed10Mbps is the 10 Mbps Ethernet speed.
	EthernetSpeed10Mbps EthernetSpeed = "10Mbps"
	// EthernetSpeed100Mbps is the 100 Mbps Ethernet speed.
	EthernetSpeed100Mbps EthernetSpeed = "100Mbps"
	// EthernetSpeed1Gbps is the 1 Gbps Ethernet speed.
	EthernetSpeed1Gbps EthernetSpeed = "1Gbps"
	// EthernetSpeed2500Mbps is the 2.5 Gbps Ethernet speed.
	EthernetSpeed2500Mbps EthernetSpeed = "2500Mbps"
	// EthernetSpeed5Gbps is the 5 Gbps Ethernet speed.
	EthernetSpeed5Gbps EthernetSpeed = "5Gbps"
	// EthernetSpeed10Gbps is the 10 Gbps Ethernet speed.
	EthernetSpeed10Gbps EthernetSpeed = "10Gbps"
	// EthernetSpeed25Gbps is the 25 Gbps Ethernet speed.
	EthernetSpeed25Gbps EthernetSpeed = "25Gbps"
	// EthernetSpeed40Gbps is the 40 Gbps Ethernet speed.
	EthernetSpeed40Gbps EthernetSpeed = "40Gbps"
	// EthernetSpeed50Gbps is the 50 Gbps Ethernet speed.
	EthernetSpeed50Gbps EthernetSpeed = "50Gbps"
	// EthernetSpeed100Gbps is the 100 Gbps Ethernet speed.
	EthernetSpeed100Gbps EthernetSpeed = "100Gbps"
	// EthernetSpeed200Gbps is the 200 Gbps Ethernet speed.
	EthernetSpeed200Gbps EthernetSpeed = "200Gbps"
	// EthernetSpeed400Gbps is the 400 Gbps Ethernet speed.
	EthernetSpeed400Gbps EthernetSpeed = "400Gbps"
	// EthernetSpeed800Gbps is the 800 Gbps Ethernet speed.
	EthernetSpeed800Gbps EthernetSpeed = "800Gbps"
	// EthernetSpeedUnknown is the unknown Ethernet speed.
	EthernetSpeedUnknown EthernetSpeed = "Unknown"
)
