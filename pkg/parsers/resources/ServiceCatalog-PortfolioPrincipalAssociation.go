package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogPortfolioPrincipalAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html
type ServiceCatalogPortfolioPrincipalAssociation struct {
	Type       string                                                `yaml:"Type"`
	Properties ServiceCatalogPortfolioPrincipalAssociationProperties `yaml:"Properties"`
	Condition  interface{}                                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                           `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogPortfolioPrincipalAssociation Properties
type ServiceCatalogPortfolioPrincipalAssociationProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	PortfolioId    interface{} `yaml:"PortfolioId"`
	PrincipalARN   interface{} `yaml:"PrincipalARN"`
	PrincipalType  interface{} `yaml:"PrincipalType"`
}

// NewServiceCatalogPortfolioPrincipalAssociation constructor creates a new ServiceCatalogPortfolioPrincipalAssociation
func NewServiceCatalogPortfolioPrincipalAssociation(properties ServiceCatalogPortfolioPrincipalAssociationProperties, deps ...interface{}) ServiceCatalogPortfolioPrincipalAssociation {
	return ServiceCatalogPortfolioPrincipalAssociation{
		Type:       "AWS::ServiceCatalog::PortfolioPrincipalAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogPortfolioPrincipalAssociation parses ServiceCatalogPortfolioPrincipalAssociation
func ParseServiceCatalogPortfolioPrincipalAssociation(
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
	var resource ServiceCatalogPortfolioPrincipalAssociation
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
					"Fn::Sub": "${AWS::StackName}-ServiceCatalogPortfolioPrincipalAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseServiceCatalogPortfolioPrincipalAssociation validator
func (resource ServiceCatalogPortfolioPrincipalAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogPortfolioPrincipalAssociationProperties validator
func (resource ServiceCatalogPortfolioPrincipalAssociationProperties) Validate() []error {
	errors := []error{}
	if resource.PortfolioId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PortfolioId'"))
	}
	if resource.PrincipalARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PrincipalARN'"))
	}
	if resource.PrincipalType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PrincipalType'"))
	}
	return errors
}
