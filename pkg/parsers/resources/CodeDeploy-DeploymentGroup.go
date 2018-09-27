package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CodeDeployDeploymentGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
type CodeDeployDeploymentGroup struct {
	Type       string                              `yaml:"Type"`
	Properties CodeDeployDeploymentGroupProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// CodeDeployDeploymentGroup Properties
type CodeDeployDeploymentGroupProperties struct {
	ApplicationName              interface{}                                          `yaml:"ApplicationName"`
	DeploymentConfigName         interface{}                                          `yaml:"DeploymentConfigName,omitempty"`
	DeploymentGroupName          interface{}                                          `yaml:"DeploymentGroupName,omitempty"`
	ServiceRoleArn               interface{}                                          `yaml:"ServiceRoleArn"`
	LoadBalancerInfo             *properties.DeploymentGroupLoadBalancerInfo          `yaml:"LoadBalancerInfo,omitempty"`
	OnPremisesInstanceTagFilters interface{}                                          `yaml:"OnPremisesInstanceTagFilters,omitempty"`
	AutoScalingGroups            interface{}                                          `yaml:"AutoScalingGroups,omitempty"`
	Ec2TagFilters                interface{}                                          `yaml:"Ec2TagFilters,omitempty"`
	TriggerConfigurations        interface{}                                          `yaml:"TriggerConfigurations,omitempty"`
	DeploymentStyle              *properties.DeploymentGroupDeploymentStyle           `yaml:"DeploymentStyle,omitempty"`
	Deployment                   *properties.DeploymentGroupDeployment                `yaml:"Deployment,omitempty"`
	AutoRollbackConfiguration    *properties.DeploymentGroupAutoRollbackConfiguration `yaml:"AutoRollbackConfiguration,omitempty"`
	AlarmConfiguration           *properties.DeploymentGroupAlarmConfiguration        `yaml:"AlarmConfiguration,omitempty"`
}

// NewCodeDeployDeploymentGroup constructor creates a new CodeDeployDeploymentGroup
func NewCodeDeployDeploymentGroup(properties CodeDeployDeploymentGroupProperties, deps ...interface{}) CodeDeployDeploymentGroup {
	return CodeDeployDeploymentGroup{
		Type:       "AWS::CodeDeploy::DeploymentGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodeDeployDeploymentGroup parses CodeDeployDeploymentGroup
func ParseCodeDeployDeploymentGroup(
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
	var resource CodeDeployDeploymentGroup
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
					"Fn::Sub": "${AWS::StackName}-CodeDeployDeploymentGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseCodeDeployDeploymentGroup validator
func (resource CodeDeployDeploymentGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodeDeployDeploymentGroupProperties validator
func (resource CodeDeployDeploymentGroupProperties) Validate() []error {
	errors := []error{}
	if resource.ApplicationName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.ServiceRoleArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ServiceRoleArn'"))
	}
	return errors
}
