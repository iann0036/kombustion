package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogTagOptionAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html
type ServiceCatalogTagOptionAssociation struct {
	Type       string                                       `yaml:"Type"`
	Properties ServiceCatalogTagOptionAssociationProperties `yaml:"Properties"`
	Condition  interface{}                                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                  `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogTagOptionAssociation Properties
type ServiceCatalogTagOptionAssociationProperties struct {
	ResourceId  interface{} `yaml:"ResourceId"`
	TagOptionId interface{} `yaml:"TagOptionId"`
}

// NewServiceCatalogTagOptionAssociation constructor creates a new ServiceCatalogTagOptionAssociation
func NewServiceCatalogTagOptionAssociation(properties ServiceCatalogTagOptionAssociationProperties, deps ...interface{}) ServiceCatalogTagOptionAssociation {
	return ServiceCatalogTagOptionAssociation{
		Type:       "AWS::ServiceCatalog::TagOptionAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogTagOptionAssociation parses ServiceCatalogTagOptionAssociation
func ParseServiceCatalogTagOptionAssociation(
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
	var resource ServiceCatalogTagOptionAssociation
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
					"Fn::Sub": "${AWS::StackName}-ServiceCatalogTagOptionAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseServiceCatalogTagOptionAssociation validator
func (resource ServiceCatalogTagOptionAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogTagOptionAssociationProperties validator
func (resource ServiceCatalogTagOptionAssociationProperties) Validate() []error {
	errors := []error{}
	if resource.ResourceId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ResourceId'"))
	}
	if resource.TagOptionId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TagOptionId'"))
	}
	return errors
}
