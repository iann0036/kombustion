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

// SESReceiptRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html
type SESReceiptRule struct {
	Type       string                   `yaml:"Type"`
	Properties SESReceiptRuleProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// SESReceiptRule Properties
type SESReceiptRuleProperties struct {
	After       interface{}                 `yaml:"After,omitempty"`
	RuleSetName interface{}                 `yaml:"RuleSetName"`
	Rule        *properties.ReceiptRuleRule `yaml:"Rule"`
}

// NewSESReceiptRule constructor creates a new SESReceiptRule
func NewSESReceiptRule(properties SESReceiptRuleProperties, deps ...interface{}) SESReceiptRule {
	return SESReceiptRule{
		Type:       "AWS::SES::ReceiptRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESReceiptRule parses SESReceiptRule
func ParseSESReceiptRule(name string, data string) (cf types.TemplateObject, err error) {
	var resource SESReceiptRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESReceiptRule - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSESReceiptRule validator
func (resource SESReceiptRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESReceiptRuleProperties validator
func (resource SESReceiptRuleProperties) Validate() []error {
	errs := []error{}
	if resource.RuleSetName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleSetName'"))
	}
	if resource.Rule == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rule'"))
	} else {
		errs = append(errs, resource.Rule.Validate()...)
	}
	return errs
}
