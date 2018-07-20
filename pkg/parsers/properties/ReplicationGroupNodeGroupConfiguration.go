package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ReplicationGroupNodeGroupConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html
type ReplicationGroupNodeGroupConfiguration struct {
	PrimaryAvailabilityZone  interface{} `yaml:"PrimaryAvailabilityZone,omitempty"`
	ReplicaCount             interface{} `yaml:"ReplicaCount,omitempty"`
	Slots                    interface{} `yaml:"Slots,omitempty"`
	ReplicaAvailabilityZones interface{} `yaml:"ReplicaAvailabilityZones,omitempty"`
}

// ReplicationGroupNodeGroupConfiguration validation
func (resource ReplicationGroupNodeGroupConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
