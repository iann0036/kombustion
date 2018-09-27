package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceDiscoveryService Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html
type ServiceDiscoveryService struct {
	Type       string                            `yaml:"Type"`
	Properties ServiceDiscoveryServiceProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// ServiceDiscoveryService Properties
type ServiceDiscoveryServiceProperties struct {
	Description             interface{}                                `yaml:"Description,omitempty"`
	Name                    interface{}                                `yaml:"Name,omitempty"`
	HealthCheckCustomConfig *properties.ServiceHealthCheckCustomConfig `yaml:"HealthCheckCustomConfig,omitempty"`
	HealthCheckConfig       *properties.ServiceHealthCheckConfig       `yaml:"HealthCheckConfig,omitempty"`
	DnsConfig               *properties.ServiceDnsConfig               `yaml:"DnsConfig"`
}

// NewServiceDiscoveryService constructor creates a new ServiceDiscoveryService
func NewServiceDiscoveryService(properties ServiceDiscoveryServiceProperties, deps ...interface{}) ServiceDiscoveryService {
	return ServiceDiscoveryService{
		Type:       "AWS::ServiceDiscovery::Service",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceDiscoveryService parses ServiceDiscoveryService
func ParseServiceDiscoveryService(
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
	var resource ServiceDiscoveryService
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
					"Fn::Sub": "${AWS::StackName}-ServiceDiscoveryService-" + name,
				},
			},
		},
	}

	return
}

// ParseServiceDiscoveryService validator
func (resource ServiceDiscoveryService) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceDiscoveryServiceProperties validator
func (resource ServiceDiscoveryServiceProperties) Validate() []error {
	errors := []error{}
	if resource.DnsConfig == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DnsConfig'"))
	} else {
		errors = append(errors, resource.DnsConfig.Validate()...)
	}
	return errors
}
