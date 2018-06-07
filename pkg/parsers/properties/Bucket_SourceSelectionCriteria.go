package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_SourceSelectionCriteria struct {
	SseKmsEncryptedObjects *Bucket_SseKmsEncryptedObjects `yaml:"SseKmsEncryptedObjects"`
}

func (resource Bucket_SourceSelectionCriteria) Validate() []error {
	errs := []error{}

	if resource.SseKmsEncryptedObjects == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SseKmsEncryptedObjects'"))
	} else {
		errs = append(errs, resource.SseKmsEncryptedObjects.Validate()...)
	}
	return errs
}