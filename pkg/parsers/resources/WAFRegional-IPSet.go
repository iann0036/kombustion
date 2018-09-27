package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRegionalIPSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html
type WAFRegionalIPSet struct {
	Type       string                     `yaml:"Type"`
	Properties WAFRegionalIPSetProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// WAFRegionalIPSet Properties
type WAFRegionalIPSetProperties struct {
	Name             interface{} `yaml:"Name"`
	IPSetDescriptors interface{} `yaml:"IPSetDescriptors,omitempty"`
}

// NewWAFRegionalIPSet constructor creates a new WAFRegionalIPSet
func NewWAFRegionalIPSet(properties WAFRegionalIPSetProperties, deps ...interface{}) WAFRegionalIPSet {
	return WAFRegionalIPSet{
		Type:       "AWS::WAFRegional::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalIPSet parses WAFRegionalIPSet
func ParseWAFRegionalIPSet(
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
	var resource WAFRegionalIPSet
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
					"Fn::Sub": "${AWS::StackName}-WAFRegionalIPSet-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFRegionalIPSet validator
func (resource WAFRegionalIPSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalIPSetProperties validator
func (resource WAFRegionalIPSetProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
