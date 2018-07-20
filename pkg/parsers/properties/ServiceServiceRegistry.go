package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ServiceServiceRegistry Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceregistry.html
type ServiceServiceRegistry struct {
	Port        interface{} `yaml:"Port,omitempty"`
	RegistryArn interface{} `yaml:"RegistryArn,omitempty"`
}

// ServiceServiceRegistry validation
func (resource ServiceServiceRegistry) Validate() []error {
	errors := []error{}

	return errors
}
