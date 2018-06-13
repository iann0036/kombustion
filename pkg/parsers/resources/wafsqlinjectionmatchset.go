package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFSqlInjectionMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
type WAFSqlInjectionMatchSet struct {
	Type       string                            `yaml:"Type"`
	Properties WAFSqlInjectionMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// WAFSqlInjectionMatchSet Properties
type WAFSqlInjectionMatchSetProperties struct {
	Name                    interface{} `yaml:"Name"`
	SqlInjectionMatchTuples interface{} `yaml:"SqlInjectionMatchTuples,omitempty"`
}

// NewWAFSqlInjectionMatchSet constructor creates a new WAFSqlInjectionMatchSet
func NewWAFSqlInjectionMatchSet(properties WAFSqlInjectionMatchSetProperties, deps ...interface{}) WAFSqlInjectionMatchSet {
	return WAFSqlInjectionMatchSet{
		Type:       "AWS::WAF::SqlInjectionMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFSqlInjectionMatchSet parses WAFSqlInjectionMatchSet
func ParseWAFSqlInjectionMatchSet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFSqlInjectionMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFSqlInjectionMatchSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFSqlInjectionMatchSet validator
func (resource WAFSqlInjectionMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFSqlInjectionMatchSetProperties validator
func (resource WAFSqlInjectionMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}