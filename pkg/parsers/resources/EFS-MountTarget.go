package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EFSMountTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
type EFSMountTarget struct {
	Type       string                   `yaml:"Type"`
	Properties EFSMountTargetProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// EFSMountTarget Properties
type EFSMountTargetProperties struct {
	FileSystemId   interface{} `yaml:"FileSystemId"`
	IpAddress      interface{} `yaml:"IpAddress,omitempty"`
	SubnetId       interface{} `yaml:"SubnetId"`
	SecurityGroups interface{} `yaml:"SecurityGroups"`
}

// NewEFSMountTarget constructor creates a new EFSMountTarget
func NewEFSMountTarget(properties EFSMountTargetProperties, deps ...interface{}) EFSMountTarget {
	return EFSMountTarget{
		Type:       "AWS::EFS::MountTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEFSMountTarget parses EFSMountTarget
func ParseEFSMountTarget(
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
	var resource EFSMountTarget
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
					"Fn::Sub": "${AWS::StackName}-EFSMountTarget-" + name,
				},
			},
		},
	}

	return
}

// ParseEFSMountTarget validator
func (resource EFSMountTarget) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEFSMountTargetProperties validator
func (resource EFSMountTargetProperties) Validate() []error {
	errors := []error{}
	if resource.FileSystemId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'FileSystemId'"))
	}
	if resource.SubnetId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	if resource.SecurityGroups == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SecurityGroups'"))
	}
	return errors
}
