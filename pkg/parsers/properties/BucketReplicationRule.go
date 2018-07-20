package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketReplicationRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html
type BucketReplicationRule struct {
	Id                      interface{}                    `yaml:"Id,omitempty"`
	Prefix                  interface{}                    `yaml:"Prefix"`
	Status                  interface{}                    `yaml:"Status"`
	SourceSelectionCriteria *BucketSourceSelectionCriteria `yaml:"SourceSelectionCriteria,omitempty"`
	Destination             *BucketReplicationDestination  `yaml:"Destination"`
}

// BucketReplicationRule validation
func (resource BucketReplicationRule) Validate() []error {
	errors := []error{}

	if resource.Prefix == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Prefix'"))
	}
	if resource.Status == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Status'"))
	}
	if resource.Destination == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Destination'"))
	} else {
		errors = append(errors, resource.Destination.Validate()...)
	}
	return errors
}
