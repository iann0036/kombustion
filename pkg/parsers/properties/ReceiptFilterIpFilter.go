package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ReceiptFilterIpFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-ipfilter.html
type ReceiptFilterIpFilter struct {
	Cidr   interface{} `yaml:"Cidr"`
	Policy interface{} `yaml:"Policy"`
}

// ReceiptFilterIpFilter validation
func (resource ReceiptFilterIpFilter) Validate() []error {
	errs := []error{}

	if resource.Cidr == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Cidr'"))
	}
	if resource.Policy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Policy'"))
	}
	return errs
}
