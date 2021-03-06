package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// RedshiftClusterSecurityGroupIngress Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
type RedshiftClusterSecurityGroupIngress struct {
	Type       string                                        `yaml:"Type"`
	Properties RedshiftClusterSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// RedshiftClusterSecurityGroupIngress Properties
type RedshiftClusterSecurityGroupIngressProperties struct {
	CIDRIP                   interface{} `yaml:"CIDRIP,omitempty"`
	ClusterSecurityGroupName interface{} `yaml:"ClusterSecurityGroupName"`
	EC2SecurityGroupName     interface{} `yaml:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId  interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

// NewRedshiftClusterSecurityGroupIngress constructor creates a new RedshiftClusterSecurityGroupIngress
func NewRedshiftClusterSecurityGroupIngress(properties RedshiftClusterSecurityGroupIngressProperties, deps ...interface{}) RedshiftClusterSecurityGroupIngress {
	return RedshiftClusterSecurityGroupIngress{
		Type:       "AWS::Redshift::ClusterSecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRedshiftClusterSecurityGroupIngress parses RedshiftClusterSecurityGroupIngress
func ParseRedshiftClusterSecurityGroupIngress(name string, data string) (cf types.TemplateObject, err error) {
	var resource RedshiftClusterSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftClusterSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseRedshiftClusterSecurityGroupIngress validator
func (resource RedshiftClusterSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRedshiftClusterSecurityGroupIngressProperties validator
func (resource RedshiftClusterSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.ClusterSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ClusterSecurityGroupName'"))
	}
	return errs
}
