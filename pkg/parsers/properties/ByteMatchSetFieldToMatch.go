package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ByteMatchSetFieldToMatch Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type ByteMatchSetFieldToMatch struct {
	Data interface{} `yaml:"Data,omitempty"`
	Type interface{} `yaml:"Type"`
}

// ByteMatchSetFieldToMatch validation
func (resource ByteMatchSetFieldToMatch) Validate() []error {
	errors := []error{}

	if resource.Type == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Type'"))
	}
	return errors
}
