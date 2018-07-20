package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html
type BucketRule struct {
	ExpirationDate                    interface{}                           `yaml:"ExpirationDate,omitempty"`
	ExpirationInDays                  interface{}                           `yaml:"ExpirationInDays,omitempty"`
	Id                                interface{}                           `yaml:"Id,omitempty"`
	NoncurrentVersionExpirationInDays interface{}                           `yaml:"NoncurrentVersionExpirationInDays,omitempty"`
	Prefix                            interface{}                           `yaml:"Prefix,omitempty"`
	Status                            interface{}                           `yaml:"Status"`
	Transition                        *BucketTransition                     `yaml:"Transition,omitempty"`
	NoncurrentVersionTransition       *BucketNoncurrentVersionTransition    `yaml:"NoncurrentVersionTransition,omitempty"`
	NoncurrentVersionTransitions      interface{}                           `yaml:"NoncurrentVersionTransitions,omitempty"`
	TagFilters                        interface{}                           `yaml:"TagFilters,omitempty"`
	Transitions                       interface{}                           `yaml:"Transitions,omitempty"`
	AbortIncompleteMultipartUpload    *BucketAbortIncompleteMultipartUpload `yaml:"AbortIncompleteMultipartUpload,omitempty"`
}

// BucketRule validation
func (resource BucketRule) Validate() []error {
	errors := []error{}

	if resource.Status == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Status'"))
	}
	return errors
}
