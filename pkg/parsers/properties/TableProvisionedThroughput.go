package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TableProvisionedThroughput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type TableProvisionedThroughput struct {
	ReadCapacityUnits  interface{} `yaml:"ReadCapacityUnits"`
	WriteCapacityUnits interface{} `yaml:"WriteCapacityUnits"`
}

// TableProvisionedThroughput validation
func (resource TableProvisionedThroughput) Validate() []error {
	errors := []error{}

	if resource.ReadCapacityUnits == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ReadCapacityUnits'"))
	}
	if resource.WriteCapacityUnits == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'WriteCapacityUnits'"))
	}
	return errors
}
