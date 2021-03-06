package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CognitoUserPoolGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html
type CognitoUserPoolGroup struct {
	Type       string                         `yaml:"Type"`
	Properties CognitoUserPoolGroupProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// CognitoUserPoolGroup Properties
type CognitoUserPoolGroupProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	GroupName   interface{} `yaml:"GroupName,omitempty"`
	Precedence  interface{} `yaml:"Precedence,omitempty"`
	RoleArn     interface{} `yaml:"RoleArn,omitempty"`
	UserPoolId  interface{} `yaml:"UserPoolId"`
}

// NewCognitoUserPoolGroup constructor creates a new CognitoUserPoolGroup
func NewCognitoUserPoolGroup(properties CognitoUserPoolGroupProperties, deps ...interface{}) CognitoUserPoolGroup {
	return CognitoUserPoolGroup{
		Type:       "AWS::Cognito::UserPoolGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPoolGroup parses CognitoUserPoolGroup
func ParseCognitoUserPoolGroup(name string, data string) (cf types.TemplateObject, err error) {
	var resource CognitoUserPoolGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCognitoUserPoolGroup validator
func (resource CognitoUserPoolGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoUserPoolGroupProperties validator
func (resource CognitoUserPoolGroupProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
