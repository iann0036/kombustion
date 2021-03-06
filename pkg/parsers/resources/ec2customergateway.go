package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2CustomerGateway Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html
type EC2CustomerGateway struct {
	Type       string                       `yaml:"Type"`
	Properties EC2CustomerGatewayProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// EC2CustomerGateway Properties
type EC2CustomerGatewayProperties struct {
	BgpAsn    interface{} `yaml:"BgpAsn"`
	IpAddress interface{} `yaml:"IpAddress"`
	Type      interface{} `yaml:"Type"`
	Tags      interface{} `yaml:"Tags,omitempty"`
}

// NewEC2CustomerGateway constructor creates a new EC2CustomerGateway
func NewEC2CustomerGateway(properties EC2CustomerGatewayProperties, deps ...interface{}) EC2CustomerGateway {
	return EC2CustomerGateway{
		Type:       "AWS::EC2::CustomerGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2CustomerGateway parses EC2CustomerGateway
func ParseEC2CustomerGateway(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2CustomerGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2CustomerGateway - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2CustomerGateway validator
func (resource EC2CustomerGateway) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2CustomerGatewayProperties validator
func (resource EC2CustomerGatewayProperties) Validate() []error {
	errs := []error{}
	if resource.BgpAsn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BgpAsn'"))
	}
	if resource.IpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpAddress'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
