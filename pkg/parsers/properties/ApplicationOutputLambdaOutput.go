package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationOutputLambdaOutput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-lambdaoutput.html
type ApplicationOutputLambdaOutput struct {
	ResourceARN interface{} `yaml:"ResourceARN"`
	RoleARN     interface{} `yaml:"RoleARN"`
}

// ApplicationOutputLambdaOutput validation
func (resource ApplicationOutputLambdaOutput) Validate() []error {
	errors := []error{}

	if resource.ResourceARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ResourceARN'"))
	}
	if resource.RoleARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errors
}
