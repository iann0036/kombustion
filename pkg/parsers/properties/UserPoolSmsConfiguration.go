package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolSmsConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html
type UserPoolSmsConfiguration struct {
	ExternalId   interface{} `yaml:"ExternalId,omitempty"`
	SnsCallerArn interface{} `yaml:"SnsCallerArn,omitempty"`
}

// UserPoolSmsConfiguration validation
func (resource UserPoolSmsConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
