package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFRegionalByteMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
type WAFRegionalByteMatchSet struct {
	Type       string                            `yaml:"Type"`
	Properties WAFRegionalByteMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// WAFRegionalByteMatchSet Properties
type WAFRegionalByteMatchSetProperties struct {
	Name            interface{} `yaml:"Name"`
	ByteMatchTuples interface{} `yaml:"ByteMatchTuples,omitempty"`
}

// NewWAFRegionalByteMatchSet constructor creates a new WAFRegionalByteMatchSet
func NewWAFRegionalByteMatchSet(properties WAFRegionalByteMatchSetProperties, deps ...interface{}) WAFRegionalByteMatchSet {
	return WAFRegionalByteMatchSet{
		Type:       "AWS::WAFRegional::ByteMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalByteMatchSet parses WAFRegionalByteMatchSet
func ParseWAFRegionalByteMatchSet(name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalByteMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalByteMatchSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFRegionalByteMatchSet validator
func (resource WAFRegionalByteMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalByteMatchSetProperties validator
func (resource WAFRegionalByteMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}