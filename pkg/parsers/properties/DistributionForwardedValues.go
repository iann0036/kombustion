package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DistributionForwardedValues Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html
type DistributionForwardedValues struct {
	QueryString          interface{}          `yaml:"QueryString"`
	Headers              interface{}          `yaml:"Headers,omitempty"`
	QueryStringCacheKeys interface{}          `yaml:"QueryStringCacheKeys,omitempty"`
	Cookies              *DistributionCookies `yaml:"Cookies,omitempty"`
}

// DistributionForwardedValues validation
func (resource DistributionForwardedValues) Validate() []error {
	errs := []error{}

	if resource.QueryString == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'QueryString'"))
	}
	return errs
}
