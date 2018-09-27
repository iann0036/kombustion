package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayRestApi Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
type ApiGatewayRestApi struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayRestApiProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// ApiGatewayRestApi Properties
type ApiGatewayRestApiProperties struct {
	ApiKeySourceType       interface{}                              `yaml:"ApiKeySourceType,omitempty"`
	Body                   interface{}                              `yaml:"Body,omitempty"`
	CloneFrom              interface{}                              `yaml:"CloneFrom,omitempty"`
	Description            interface{}                              `yaml:"Description,omitempty"`
	FailOnWarnings         interface{}                              `yaml:"FailOnWarnings,omitempty"`
	MinimumCompressionSize interface{}                              `yaml:"MinimumCompressionSize,omitempty"`
	Name                   interface{}                              `yaml:"Name,omitempty"`
	Policy                 interface{}                              `yaml:"Policy,omitempty"`
	BodyS3Location         *properties.RestApiS3Location            `yaml:"BodyS3Location,omitempty"`
	Parameters             interface{}                              `yaml:"Parameters,omitempty"`
	BinaryMediaTypes       interface{}                              `yaml:"BinaryMediaTypes,omitempty"`
	EndpointConfiguration  *properties.RestApiEndpointConfiguration `yaml:"EndpointConfiguration,omitempty"`
}

// NewApiGatewayRestApi constructor creates a new ApiGatewayRestApi
func NewApiGatewayRestApi(properties ApiGatewayRestApiProperties, deps ...interface{}) ApiGatewayRestApi {
	return ApiGatewayRestApi{
		Type:       "AWS::ApiGateway::RestApi",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayRestApi parses ApiGatewayRestApi
func ParseApiGatewayRestApi(
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
	var resource ApiGatewayRestApi
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayRestApi-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayRestApi validator
func (resource ApiGatewayRestApi) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayRestApiProperties validator
func (resource ApiGatewayRestApiProperties) Validate() []error {
	errors := []error{}
	return errors
}
