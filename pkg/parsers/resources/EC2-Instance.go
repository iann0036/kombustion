package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2Instance Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
type EC2Instance struct {
	Type       string                `yaml:"Type"`
	Properties EC2InstanceProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// EC2Instance Properties
type EC2InstanceProperties struct {
	AdditionalInfo                    interface{}                                     `yaml:"AdditionalInfo,omitempty"`
	Affinity                          interface{}                                     `yaml:"Affinity,omitempty"`
	AvailabilityZone                  interface{}                                     `yaml:"AvailabilityZone,omitempty"`
	DisableApiTermination             interface{}                                     `yaml:"DisableApiTermination,omitempty"`
	EbsOptimized                      interface{}                                     `yaml:"EbsOptimized,omitempty"`
	HostId                            interface{}                                     `yaml:"HostId,omitempty"`
	IamInstanceProfile                interface{}                                     `yaml:"IamInstanceProfile,omitempty"`
	ImageId                           interface{}                                     `yaml:"ImageId,omitempty"`
	InstanceInitiatedShutdownBehavior interface{}                                     `yaml:"InstanceInitiatedShutdownBehavior,omitempty"`
	InstanceType                      interface{}                                     `yaml:"InstanceType,omitempty"`
	Ipv6AddressCount                  interface{}                                     `yaml:"Ipv6AddressCount,omitempty"`
	KernelId                          interface{}                                     `yaml:"KernelId,omitempty"`
	KeyName                           interface{}                                     `yaml:"KeyName,omitempty"`
	Monitoring                        interface{}                                     `yaml:"Monitoring,omitempty"`
	PlacementGroupName                interface{}                                     `yaml:"PlacementGroupName,omitempty"`
	PrivateIpAddress                  interface{}                                     `yaml:"PrivateIpAddress,omitempty"`
	RamdiskId                         interface{}                                     `yaml:"RamdiskId,omitempty"`
	SourceDestCheck                   interface{}                                     `yaml:"SourceDestCheck,omitempty"`
	SubnetId                          interface{}                                     `yaml:"SubnetId,omitempty"`
	Tenancy                           interface{}                                     `yaml:"Tenancy,omitempty"`
	UserData                          interface{}                                     `yaml:"UserData,omitempty"`
	BlockDeviceMappings               interface{}                                     `yaml:"BlockDeviceMappings,omitempty"`
	SecurityGroups                    interface{}                                     `yaml:"SecurityGroups,omitempty"`
	ElasticGpuSpecifications          interface{}                                     `yaml:"ElasticGpuSpecifications,omitempty"`
	Ipv6Addresses                     interface{}                                     `yaml:"Ipv6Addresses,omitempty"`
	Volumes                           interface{}                                     `yaml:"Volumes,omitempty"`
	NetworkInterfaces                 interface{}                                     `yaml:"NetworkInterfaces,omitempty"`
	SecurityGroupIds                  interface{}                                     `yaml:"SecurityGroupIds,omitempty"`
	SsmAssociations                   interface{}                                     `yaml:"SsmAssociations,omitempty"`
	Tags                              interface{}                                     `yaml:"Tags,omitempty"`
	LaunchTemplate                    *properties.InstanceLaunchTemplateSpecification `yaml:"LaunchTemplate,omitempty"`
	CreditSpecification               *properties.InstanceCreditSpecification         `yaml:"CreditSpecification,omitempty"`
}

// NewEC2Instance constructor creates a new EC2Instance
func NewEC2Instance(properties EC2InstanceProperties, deps ...interface{}) EC2Instance {
	return EC2Instance{
		Type:       "AWS::EC2::Instance",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2Instance parses EC2Instance
func ParseEC2Instance(
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
	var resource EC2Instance
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
					"Fn::Sub": "${AWS::StackName}-EC2Instance-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2Instance validator
func (resource EC2Instance) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2InstanceProperties validator
func (resource EC2InstanceProperties) Validate() []error {
	errors := []error{}
	return errors
}
