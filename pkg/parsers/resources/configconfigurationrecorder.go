package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ConfigConfigurationRecorder Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
type ConfigConfigurationRecorder struct {
	Type       string                                `yaml:"Type"`
	Properties ConfigConfigurationRecorderProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ConfigConfigurationRecorder Properties
type ConfigConfigurationRecorderProperties struct {
	Name           interface{}                                     `yaml:"Name,omitempty"`
	RoleARN        interface{}                                     `yaml:"RoleARN"`
	RecordingGroup *properties.ConfigurationRecorderRecordingGroup `yaml:"RecordingGroup,omitempty"`
}

// NewConfigConfigurationRecorder constructor creates a new ConfigConfigurationRecorder
func NewConfigConfigurationRecorder(properties ConfigConfigurationRecorderProperties, deps ...interface{}) ConfigConfigurationRecorder {
	return ConfigConfigurationRecorder{
		Type:       "AWS::Config::ConfigurationRecorder",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseConfigConfigurationRecorder parses ConfigConfigurationRecorder
func ParseConfigConfigurationRecorder(name string, data string) (cf types.TemplateObject, err error) {
	var resource ConfigConfigurationRecorder
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ConfigConfigurationRecorder - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseConfigConfigurationRecorder validator
func (resource ConfigConfigurationRecorder) Validate() []error {
	return resource.Properties.Validate()
}

// ParseConfigConfigurationRecorderProperties validator
func (resource ConfigConfigurationRecorderProperties) Validate() []error {
	errs := []error{}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
