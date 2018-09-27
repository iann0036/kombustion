package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ElasticLoadBalancingLoadBalancer Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type ElasticLoadBalancingLoadBalancer struct {
	Type       string                                     `yaml:"Type"`
	Properties ElasticLoadBalancingLoadBalancerProperties `yaml:"Properties"`
	Condition  interface{}                                `yaml:"Condition,omitempty"`
	Metadata   interface{}                                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                `yaml:"DependsOn,omitempty"`
}

// ElasticLoadBalancingLoadBalancer Properties
type ElasticLoadBalancingLoadBalancerProperties struct {
	CrossZone                 interface{}                                      `yaml:"CrossZone,omitempty"`
	LoadBalancerName          interface{}                                      `yaml:"LoadBalancerName,omitempty"`
	Scheme                    interface{}                                      `yaml:"Scheme,omitempty"`
	Tags                      interface{}                                      `yaml:"Tags,omitempty"`
	AppCookieStickinessPolicy interface{}                                      `yaml:"AppCookieStickinessPolicy,omitempty"`
	AvailabilityZones         interface{}                                      `yaml:"AvailabilityZones,omitempty"`
	Subnets                   interface{}                                      `yaml:"Subnets,omitempty"`
	SecurityGroups            interface{}                                      `yaml:"SecurityGroups,omitempty"`
	Policies                  interface{}                                      `yaml:"Policies,omitempty"`
	Instances                 interface{}                                      `yaml:"Instances,omitempty"`
	LBCookieStickinessPolicy  interface{}                                      `yaml:"LBCookieStickinessPolicy,omitempty"`
	Listeners                 interface{}                                      `yaml:"Listeners"`
	HealthCheck               *properties.LoadBalancerHealthCheck              `yaml:"HealthCheck,omitempty"`
	ConnectionSettings        *properties.LoadBalancerConnectionSettings       `yaml:"ConnectionSettings,omitempty"`
	ConnectionDrainingPolicy  *properties.LoadBalancerConnectionDrainingPolicy `yaml:"ConnectionDrainingPolicy,omitempty"`
	AccessLoggingPolicy       *properties.LoadBalancerAccessLoggingPolicy      `yaml:"AccessLoggingPolicy,omitempty"`
}

// NewElasticLoadBalancingLoadBalancer constructor creates a new ElasticLoadBalancingLoadBalancer
func NewElasticLoadBalancingLoadBalancer(properties ElasticLoadBalancingLoadBalancerProperties, deps ...interface{}) ElasticLoadBalancingLoadBalancer {
	return ElasticLoadBalancingLoadBalancer{
		Type:       "AWS::ElasticLoadBalancing::LoadBalancer",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticLoadBalancingLoadBalancer parses ElasticLoadBalancingLoadBalancer
func ParseElasticLoadBalancingLoadBalancer(
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
	var resource ElasticLoadBalancingLoadBalancer
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
					"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name,
				},
			},
		},
	}

	return
}

// ParseElasticLoadBalancingLoadBalancer validator
func (resource ElasticLoadBalancingLoadBalancer) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticLoadBalancingLoadBalancerProperties validator
func (resource ElasticLoadBalancingLoadBalancerProperties) Validate() []error {
	errors := []error{}
	if resource.Listeners == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Listeners'"))
	}
	return errors
}
