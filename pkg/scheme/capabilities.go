package scheme

// Capability scheme for the Synse Server "capabilities" endpoint response.
type Capability struct {
	Plugin  string             `json:"plugin" yaml:"plugin"`
	Devices []CapabilityDevice `json:"devices" yaml:"devices"`
}

// CapabilityDevice describes a device in device capability.
type CapabilityDevice struct {
	Kind    string   `json:"kind" yaml:"kind"`
	Outputs []string `json:"outputs" yaml:"outputs"`
}
