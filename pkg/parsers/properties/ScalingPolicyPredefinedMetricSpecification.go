package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPolicyPredefinedMetricSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html
type ScalingPolicyPredefinedMetricSpecification struct {
	PredefinedMetricType interface{} `yaml:"PredefinedMetricType"`
	ResourceLabel        interface{} `yaml:"ResourceLabel,omitempty"`
}

// ScalingPolicyPredefinedMetricSpecification validation
func (resource ScalingPolicyPredefinedMetricSpecification) Validate() []error {
	errs := []error{}

	if resource.PredefinedMetricType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PredefinedMetricType'"))
	}
	return errs
}