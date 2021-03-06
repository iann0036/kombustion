package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SDBDomain Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type SDBDomain struct {
	Type       string              `yaml:"Type"`
	Properties SDBDomainProperties `yaml:"Properties"`
	Condition  interface{}         `yaml:"Condition,omitempty"`
	Metadata   interface{}         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}         `yaml:"DependsOn,omitempty"`
}

// SDBDomain Properties
type SDBDomainProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
}

// NewSDBDomain constructor creates a new SDBDomain
func NewSDBDomain(properties SDBDomainProperties, deps ...interface{}) SDBDomain {
	return SDBDomain{
		Type:       "AWS::SDB::Domain",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSDBDomain parses SDBDomain
func ParseSDBDomain(name string, data string) (cf types.TemplateObject, err error) {
	var resource SDBDomain
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SDBDomain - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSDBDomain validator
func (resource SDBDomain) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSDBDomainProperties validator
func (resource SDBDomainProperties) Validate() []error {
	errs := []error{}
	return errs
}
