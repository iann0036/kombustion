package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceInstanceIpv6Address Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-instanceipv6address.html
type InstanceInstanceIpv6Address struct {
	Ipv6Address interface{} `yaml:"Ipv6Address"`
}

// InstanceInstanceIpv6Address validation
func (resource InstanceInstanceIpv6Address) Validate() []error {
	errs := []error{}

	if resource.Ipv6Address == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Ipv6Address'"))
	}
	return errs
}
