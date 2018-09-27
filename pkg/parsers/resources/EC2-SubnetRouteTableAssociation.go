package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2SubnetRouteTableAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
type EC2SubnetRouteTableAssociation struct {
	Type       string                                   `yaml:"Type"`
	Properties EC2SubnetRouteTableAssociationProperties `yaml:"Properties"`
	Condition  interface{}                              `yaml:"Condition,omitempty"`
	Metadata   interface{}                              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                              `yaml:"DependsOn,omitempty"`
}

// EC2SubnetRouteTableAssociation Properties
type EC2SubnetRouteTableAssociationProperties struct {
	RouteTableId interface{} `yaml:"RouteTableId"`
	SubnetId     interface{} `yaml:"SubnetId"`
}

// NewEC2SubnetRouteTableAssociation constructor creates a new EC2SubnetRouteTableAssociation
func NewEC2SubnetRouteTableAssociation(properties EC2SubnetRouteTableAssociationProperties, deps ...interface{}) EC2SubnetRouteTableAssociation {
	return EC2SubnetRouteTableAssociation{
		Type:       "AWS::EC2::SubnetRouteTableAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SubnetRouteTableAssociation parses EC2SubnetRouteTableAssociation
func ParseEC2SubnetRouteTableAssociation(
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
	var resource EC2SubnetRouteTableAssociation
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
					"Fn::Sub": "${AWS::StackName}-EC2SubnetRouteTableAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2SubnetRouteTableAssociation validator
func (resource EC2SubnetRouteTableAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SubnetRouteTableAssociationProperties validator
func (resource EC2SubnetRouteTableAssociationProperties) Validate() []error {
	errors := []error{}
	if resource.RouteTableId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RouteTableId'"))
	}
	if resource.SubnetId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errors
}
