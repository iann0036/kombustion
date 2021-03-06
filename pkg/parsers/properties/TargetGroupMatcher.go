package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TargetGroupMatcher Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
type TargetGroupMatcher struct {
	HttpCode interface{} `yaml:"HttpCode"`
}

// TargetGroupMatcher validation
func (resource TargetGroupMatcher) Validate() []error {
	errs := []error{}

	if resource.HttpCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HttpCode'"))
	}
	return errs
}
