package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BudgetNotification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html
type BudgetNotification struct {
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	NotificationType   interface{} `yaml:"NotificationType"`
	Threshold          interface{} `yaml:"Threshold"`
	ThresholdType      interface{} `yaml:"ThresholdType,omitempty"`
}

// BudgetNotification validation
func (resource BudgetNotification) Validate() []error {
	errs := []error{}

	if resource.ComparisonOperator == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComparisonOperator'"))
	}
	if resource.NotificationType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NotificationType'"))
	}
	if resource.Threshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Threshold'"))
	}
	return errs
}