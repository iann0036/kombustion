package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayBasePathMapping Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html
type ApiGatewayBasePathMapping struct {
	Type       string                              `yaml:"Type"`
	Properties ApiGatewayBasePathMappingProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// ApiGatewayBasePathMapping Properties
type ApiGatewayBasePathMappingProperties struct {
	BasePath   interface{} `yaml:"BasePath,omitempty"`
	DomainName interface{} `yaml:"DomainName"`
	RestApiId  interface{} `yaml:"RestApiId,omitempty"`
	Stage      interface{} `yaml:"Stage,omitempty"`
}

// NewApiGatewayBasePathMapping constructor creates a new ApiGatewayBasePathMapping
func NewApiGatewayBasePathMapping(properties ApiGatewayBasePathMappingProperties, deps ...interface{}) ApiGatewayBasePathMapping {
	return ApiGatewayBasePathMapping{
		Type:       "AWS::ApiGateway::BasePathMapping",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayBasePathMapping parses ApiGatewayBasePathMapping
func ParseApiGatewayBasePathMapping(
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
	var resource ApiGatewayBasePathMapping
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayBasePathMapping-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayBasePathMapping validator
func (resource ApiGatewayBasePathMapping) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayBasePathMappingProperties validator
func (resource ApiGatewayBasePathMappingProperties) Validate() []error {
	errors := []error{}
	if resource.DomainName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errors
}
