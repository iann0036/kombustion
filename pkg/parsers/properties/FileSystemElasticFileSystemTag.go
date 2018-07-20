package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// FileSystemElasticFileSystemTag Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
type FileSystemElasticFileSystemTag struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

// FileSystemElasticFileSystemTag validation
func (resource FileSystemElasticFileSystemTag) Validate() []error {
	errors := []error{}

	if resource.Key == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
