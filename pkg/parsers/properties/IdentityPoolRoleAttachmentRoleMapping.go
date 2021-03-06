package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// IdentityPoolRoleAttachmentRoleMapping Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html
type IdentityPoolRoleAttachmentRoleMapping struct {
	AmbiguousRoleResolution interface{}                                       `yaml:"AmbiguousRoleResolution,omitempty"`
	Type                    interface{}                                       `yaml:"Type"`
	RulesConfiguration      *IdentityPoolRoleAttachmentRulesConfigurationType `yaml:"RulesConfiguration,omitempty"`
}

// IdentityPoolRoleAttachmentRoleMapping validation
func (resource IdentityPoolRoleAttachmentRoleMapping) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
