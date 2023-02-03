package interfaces

// Mode is the mode of the VLAN.
type Mode string

const (
	// ModeAccess is the access mode, sometimes called "untagged".
	ModeAccess Mode = "Access"
	// ModeTrunk is the trunk mode, sometimes called "tagged".
	ModeTrunk Mode = "Trunk"
)

// VLANID is the ID of the VLAN.
type VLANID int
