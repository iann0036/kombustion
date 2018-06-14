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

// ElasticLoadBalancingV2TargetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html
type ElasticLoadBalancingV2TargetGroup struct {
	Type       string                                      `yaml:"Type"`
	Properties ElasticLoadBalancingV2TargetGroupProperties `yaml:"Properties"`
	Condition  interface{}                                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                 `yaml:"DependsOn,omitempty"`
}

// ElasticLoadBalancingV2TargetGroup Properties
type ElasticLoadBalancingV2TargetGroupProperties struct {
	HealthCheckIntervalSeconds interface{}                    `yaml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            interface{}                    `yaml:"HealthCheckPath,omitempty"`
	HealthCheckPort            interface{}                    `yaml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        interface{}                    `yaml:"HealthCheckProtocol,omitempty"`
	HealthCheckTimeoutSeconds  interface{}                    `yaml:"HealthCheckTimeoutSeconds,omitempty"`
	HealthyThresholdCount      interface{}                    `yaml:"HealthyThresholdCount,omitempty"`
	Name                       interface{}                    `yaml:"Name,omitempty"`
	Port                       interface{}                    `yaml:"Port"`
	Protocol                   interface{}                    `yaml:"Protocol"`
	TargetType                 interface{}                    `yaml:"TargetType,omitempty"`
	UnhealthyThresholdCount    interface{}                    `yaml:"UnhealthyThresholdCount,omitempty"`
	VpcId                      interface{}                    `yaml:"VpcId"`
	Matcher                    *properties.TargetGroupMatcher `yaml:"Matcher,omitempty"`
	Tags                       interface{}                    `yaml:"Tags,omitempty"`
	TargetGroupAttributes      interface{}                    `yaml:"TargetGroupAttributes,omitempty"`
	Targets                    interface{}                    `yaml:"Targets,omitempty"`
}

// NewElasticLoadBalancingV2TargetGroup constructor creates a new ElasticLoadBalancingV2TargetGroup
func NewElasticLoadBalancingV2TargetGroup(properties ElasticLoadBalancingV2TargetGroupProperties, deps ...interface{}) ElasticLoadBalancingV2TargetGroup {
	return ElasticLoadBalancingV2TargetGroup{
		Type:       "AWS::ElasticLoadBalancingV2::TargetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticLoadBalancingV2TargetGroup parses ElasticLoadBalancingV2TargetGroup
func ParseElasticLoadBalancingV2TargetGroup(name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticLoadBalancingV2TargetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2TargetGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticLoadBalancingV2TargetGroup validator
func (resource ElasticLoadBalancingV2TargetGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticLoadBalancingV2TargetGroupProperties validator
func (resource ElasticLoadBalancingV2TargetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Port == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Port'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
