package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ECSService Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
type ECSService struct {
	Type       string               `yaml:"Type"`
	Properties ECSServiceProperties `yaml:"Properties"`
	Condition  interface{}          `yaml:"Condition,omitempty"`
	Metadata   interface{}          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}          `yaml:"DependsOn,omitempty"`
}

// ECSService Properties
type ECSServiceProperties struct {
	Cluster                       interface{}                                `yaml:"Cluster,omitempty"`
	DesiredCount                  interface{}                                `yaml:"DesiredCount,omitempty"`
	HealthCheckGracePeriodSeconds interface{}                                `yaml:"HealthCheckGracePeriodSeconds,omitempty"`
	LaunchType                    interface{}                                `yaml:"LaunchType,omitempty"`
	PlatformVersion               interface{}                                `yaml:"PlatformVersion,omitempty"`
	Role                          interface{}                                `yaml:"Role,omitempty"`
	ServiceName                   interface{}                                `yaml:"ServiceName,omitempty"`
	TaskDefinition                interface{}                                `yaml:"TaskDefinition"`
	NetworkConfiguration          *properties.ServiceNetworkConfiguration    `yaml:"NetworkConfiguration,omitempty"`
	LoadBalancers                 interface{}                                `yaml:"LoadBalancers,omitempty"`
	PlacementConstraints          interface{}                                `yaml:"PlacementConstraints,omitempty"`
	PlacementStrategies           interface{}                                `yaml:"PlacementStrategies,omitempty"`
	ServiceRegistries             interface{}                                `yaml:"ServiceRegistries,omitempty"`
	DeploymentConfiguration       *properties.ServiceDeploymentConfiguration `yaml:"DeploymentConfiguration,omitempty"`
}

// NewECSService constructor creates a new ECSService
func NewECSService(properties ECSServiceProperties, deps ...interface{}) ECSService {
	return ECSService{
		Type:       "AWS::ECS::Service",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseECSService parses ECSService
func ParseECSService(
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
	var resource ECSService
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
					"Fn::Sub": "${AWS::StackName}-ECSService-" + name,
				},
			},
		},
	}

	return
}

// ParseECSService validator
func (resource ECSService) Validate() []error {
	return resource.Properties.Validate()
}

// ParseECSServiceProperties validator
func (resource ECSServiceProperties) Validate() []error {
	errors := []error{}
	if resource.TaskDefinition == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TaskDefinition'"))
	}
	return errors
}
