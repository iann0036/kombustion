package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// IAMInstanceProfile Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type IAMInstanceProfile struct {
	Type       string                       `yaml:"Type"`
	Properties IAMInstanceProfileProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// IAMInstanceProfile Properties
type IAMInstanceProfileProperties struct {
	InstanceProfileName interface{} `yaml:"InstanceProfileName,omitempty"`
	Path                interface{} `yaml:"Path,omitempty"`
	Roles               interface{} `yaml:"Roles"`
}

// NewIAMInstanceProfile constructor creates a new IAMInstanceProfile
func NewIAMInstanceProfile(properties IAMInstanceProfileProperties, deps ...interface{}) IAMInstanceProfile {
	return IAMInstanceProfile{
		Type:       "AWS::IAM::InstanceProfile",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIAMInstanceProfile parses IAMInstanceProfile
func ParseIAMInstanceProfile(name string, data string) (cf types.TemplateObject, err error) {
	var resource IAMInstanceProfile
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMInstanceProfile - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseIAMInstanceProfile validator
func (resource IAMInstanceProfile) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIAMInstanceProfileProperties validator
func (resource IAMInstanceProfileProperties) Validate() []error {
	errs := []error{}
	if resource.Roles == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Roles'"))
	}
	return errs
}
