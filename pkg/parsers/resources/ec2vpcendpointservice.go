package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2VPCEndpointService Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html
type EC2VPCEndpointService struct {
	Type       string                          `yaml:"Type"`
	Properties EC2VPCEndpointServiceProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// EC2VPCEndpointService Properties
type EC2VPCEndpointServiceProperties struct {
	AcceptanceRequired      interface{} `yaml:"AcceptanceRequired,omitempty"`
	NetworkLoadBalancerArns interface{} `yaml:"NetworkLoadBalancerArns"`
}

// NewEC2VPCEndpointService constructor creates a new EC2VPCEndpointService
func NewEC2VPCEndpointService(properties EC2VPCEndpointServiceProperties, deps ...interface{}) EC2VPCEndpointService {
	return EC2VPCEndpointService{
		Type:       "AWS::EC2::VPCEndpointService",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCEndpointService parses EC2VPCEndpointService
func ParseEC2VPCEndpointService(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPCEndpointService
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCEndpointService - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2VPCEndpointService validator
func (resource EC2VPCEndpointService) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCEndpointServiceProperties validator
func (resource EC2VPCEndpointServiceProperties) Validate() []error {
	errs := []error{}
	if resource.NetworkLoadBalancerArns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NetworkLoadBalancerArns'"))
	}
	return errs
}