package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayDocumentationPart Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html
type ApiGatewayDocumentationPart struct {
	Type       string                                `yaml:"Type"`
	Properties ApiGatewayDocumentationPartProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ApiGatewayDocumentationPart Properties
type ApiGatewayDocumentationPartProperties struct {
	Properties interface{}                           `yaml:"Properties"`
	RestApiId  interface{}                           `yaml:"RestApiId"`
	Location   *properties.DocumentationPartLocation `yaml:"Location"`
}

// NewApiGatewayDocumentationPart constructor creates a new ApiGatewayDocumentationPart
func NewApiGatewayDocumentationPart(properties ApiGatewayDocumentationPartProperties, deps ...interface{}) ApiGatewayDocumentationPart {
	return ApiGatewayDocumentationPart{
		Type:       "AWS::ApiGateway::DocumentationPart",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayDocumentationPart parses ApiGatewayDocumentationPart
func ParseApiGatewayDocumentationPart(
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
	var resource ApiGatewayDocumentationPart
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayDocumentationPart-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayDocumentationPart validator
func (resource ApiGatewayDocumentationPart) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayDocumentationPartProperties validator
func (resource ApiGatewayDocumentationPartProperties) Validate() []error {
	errors := []error{}
	if resource.Properties == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Properties'"))
	}
	if resource.RestApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	if resource.Location == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Location'"))
	} else {
		errors = append(errors, resource.Location.Validate()...)
	}
	return errors
}
