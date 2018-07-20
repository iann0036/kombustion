package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DataSourceHttpConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html
type DataSourceHttpConfig struct {
	Endpoint interface{} `yaml:"Endpoint"`
}

// DataSourceHttpConfig validation
func (resource DataSourceHttpConfig) Validate() []error {
	errors := []error{}

	if resource.Endpoint == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Endpoint'"))
	}
	return errors
}
