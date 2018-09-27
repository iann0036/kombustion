package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// RedshiftClusterSecurityGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type RedshiftClusterSecurityGroup struct {
	Type       string                                 `yaml:"Type"`
	Properties RedshiftClusterSecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// RedshiftClusterSecurityGroup Properties
type RedshiftClusterSecurityGroupProperties struct {
	Description interface{} `yaml:"Description"`
	Tags        interface{} `yaml:"Tags,omitempty"`
}

// NewRedshiftClusterSecurityGroup constructor creates a new RedshiftClusterSecurityGroup
func NewRedshiftClusterSecurityGroup(properties RedshiftClusterSecurityGroupProperties, deps ...interface{}) RedshiftClusterSecurityGroup {
	return RedshiftClusterSecurityGroup{
		Type:       "AWS::Redshift::ClusterSecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRedshiftClusterSecurityGroup parses RedshiftClusterSecurityGroup
func ParseRedshiftClusterSecurityGroup(
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
	var resource RedshiftClusterSecurityGroup
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
					"Fn::Sub": "${AWS::StackName}-RedshiftClusterSecurityGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseRedshiftClusterSecurityGroup validator
func (resource RedshiftClusterSecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRedshiftClusterSecurityGroupProperties validator
func (resource RedshiftClusterSecurityGroupProperties) Validate() []error {
	errors := []error{}
	if resource.Description == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Description'"))
	}
	return errors
}
