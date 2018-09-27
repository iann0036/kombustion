package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2EIP Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type EC2EIP struct {
	Type       string           `yaml:"Type"`
	Properties EC2EIPProperties `yaml:"Properties"`
	Condition  interface{}      `yaml:"Condition,omitempty"`
	Metadata   interface{}      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}      `yaml:"DependsOn,omitempty"`
}

// EC2EIP Properties
type EC2EIPProperties struct {
	Domain     interface{} `yaml:"Domain,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
}

// NewEC2EIP constructor creates a new EC2EIP
func NewEC2EIP(properties EC2EIPProperties, deps ...interface{}) EC2EIP {
	return EC2EIP{
		Type:       "AWS::EC2::EIP",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2EIP parses EC2EIP
func ParseEC2EIP(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource EC2EIP
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EC2EIP-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2EIP validator
func (resource EC2EIP) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2EIPProperties validator
func (resource EC2EIPProperties) Validate() []error {
	errors := []error{}
	return errors
}
