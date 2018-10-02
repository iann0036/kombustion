package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// AutoScalingGroupLaunchTemplateSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html
type AutoScalingGroupLaunchTemplateSpecification struct {
	LaunchTemplateId   interface{} `yaml:"LaunchTemplateId,omitempty"`
	LaunchTemplateName interface{} `yaml:"LaunchTemplateName,omitempty"`
	Version            interface{} `yaml:"Version"`
}

// AutoScalingGroupLaunchTemplateSpecification validation
func (resource AutoScalingGroupLaunchTemplateSpecification) Validate() []error {
	errors := []error{}

	if resource.Version == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Version'"))
	}
	return errors
}