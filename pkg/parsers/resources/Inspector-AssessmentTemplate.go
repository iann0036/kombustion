package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// InspectorAssessmentTemplate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html
type InspectorAssessmentTemplate struct {
	Type       string                                `yaml:"Type"`
	Properties InspectorAssessmentTemplateProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// InspectorAssessmentTemplate Properties
type InspectorAssessmentTemplateProperties struct {
	AssessmentTargetArn       interface{} `yaml:"AssessmentTargetArn"`
	AssessmentTemplateName    interface{} `yaml:"AssessmentTemplateName,omitempty"`
	DurationInSeconds         interface{} `yaml:"DurationInSeconds"`
	RulesPackageArns          interface{} `yaml:"RulesPackageArns"`
	UserAttributesForFindings interface{} `yaml:"UserAttributesForFindings,omitempty"`
}

// NewInspectorAssessmentTemplate constructor creates a new InspectorAssessmentTemplate
func NewInspectorAssessmentTemplate(properties InspectorAssessmentTemplateProperties, deps ...interface{}) InspectorAssessmentTemplate {
	return InspectorAssessmentTemplate{
		Type:       "AWS::Inspector::AssessmentTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseInspectorAssessmentTemplate parses InspectorAssessmentTemplate
func ParseInspectorAssessmentTemplate(
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
	var resource InspectorAssessmentTemplate
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
					"Fn::Sub": "${AWS::StackName}-InspectorAssessmentTemplate-" + name,
				},
			},
		},
	}

	return
}

// ParseInspectorAssessmentTemplate validator
func (resource InspectorAssessmentTemplate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseInspectorAssessmentTemplateProperties validator
func (resource InspectorAssessmentTemplateProperties) Validate() []error {
	errors := []error{}
	if resource.AssessmentTargetArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AssessmentTargetArn'"))
	}
	if resource.DurationInSeconds == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DurationInSeconds'"))
	}
	if resource.RulesPackageArns == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RulesPackageArns'"))
	}
	return errors
}
