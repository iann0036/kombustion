package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationApplicationVersionLifecycleConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html
type ApplicationApplicationVersionLifecycleConfig struct {
	MaxCountRule *ApplicationMaxCountRule `yaml:"MaxCountRule,omitempty"`
	MaxAgeRule   *ApplicationMaxAgeRule   `yaml:"MaxAgeRule,omitempty"`
}

// ApplicationApplicationVersionLifecycleConfig validation
func (resource ApplicationApplicationVersionLifecycleConfig) Validate() []error {
	errors := []error{}

	return errors
}
