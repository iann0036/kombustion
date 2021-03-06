package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StreamingDistributionLogging Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-logging.html
type StreamingDistributionLogging struct {
	Bucket  interface{} `yaml:"Bucket"`
	Enabled interface{} `yaml:"Enabled"`
	Prefix  interface{} `yaml:"Prefix"`
}

// StreamingDistributionLogging validation
func (resource StreamingDistributionLogging) Validate() []error {
	errs := []error{}

	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	return errs
}
