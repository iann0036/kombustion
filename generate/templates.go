package main

// ./parsers/resources.go
const parserMapTemplate = `{{$MainPackageName := .MainPackageName}}package {{$MainPackageName}}
{{$PackageName := .PackageName}}
{{$PackageNameTitle := .PackageNameTitle}}
// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"


import (
  "github.com/KablamoOSS/kombustion/types"
  "github.com/KablamoOSS/kombustion/pkg/{{$MainPackageName}}/{{$PackageName}}"
)

// GetParsers{{$PackageNameTitle}} returns parser functions
func GetParsers{{$PackageNameTitle}}() map[string]types.ParserFunc {
	return map[string]types.ParserFunc{
		{{range $ResourceType, $ResourceName := .ResourceTypes}}"{{$ResourceType}}": {{$PackageName}}.Parse{{$ResourceName}},
		{{end}}
	}
}
`

// ./parsers/properties/*.go
const propertyTemplate = `package properties
// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

{{$PropertyName := .PropertyName}}

{{- if .NeedsFmtImport}}
	import "fmt"
{{- end}}

// {{$PropertyName}} Documentation: {{.Documentation}}
type {{$PropertyName}} struct {
	{{- range $property := .PropertyStrings}}
	{{$property}}
	{{- end}}
}

// {{$PropertyName}} validation
func (resource {{$PropertyName}}) Validate() []error {
	errors := []error{}
	{{- range $validator := .ValidatorStrings}}
	{{$validator}}
	{{- end}}
	return errors
}
`

// ./parsers/resources/*.go
const resourceTemplate = `package resources
// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

{{$MainPackageName := .MainPackageName}}
{{$ResourceName := .ResourceName -}}
{{- $BT := "` + "`" + `" -}}

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	{{- if .NeedsFmtImport}}
	"fmt"
	{{- end}}
	{{- if .NeedsPropertiesImport}}
	"github.com/KablamoOSS/kombustion/pkg/{{$MainPackageName}}/properties"
	{{- end}}
)

// {{$ResourceName}} Documentation: {{.Documentation}}
type {{$ResourceName}} struct {
	Type       string                      {{$BT}}yaml:"Type"{{$BT}}
	Properties {{$ResourceName}}Properties {{$BT}}yaml:"Properties"{{$BT}}
	Condition  interface{}                 {{$BT}}yaml:"Condition,omitempty"{{$BT}}
	Metadata   interface{}                 {{$BT}}yaml:"Metadata,omitempty"{{$BT}}
	DependsOn  interface{}                 {{$BT}}yaml:"DependsOn,omitempty"{{$BT}}
}

// {{$ResourceName}} Properties
type {{$ResourceName}}Properties struct {
	{{- range $property := .PropertyStrings}}
	{{$property}}
	{{- end}}
}

// New{{$ResourceName}} constructor creates a new {{$ResourceName}}
func New{{$ResourceName}}(properties {{$ResourceName}}Properties, deps ...interface{}) {{$ResourceName}} {
	return {{$ResourceName}}{
		Type:       "{{.Type}}",
		Properties: properties,
		DependsOn:  deps,
	}
}

// Parse{{$ResourceName}} parses {{$ResourceName}}
func Parse{{$ResourceName}}(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource {{$ResourceName}}
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
			errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	{{if .Attributes}}
	var resource, output types.TemplateObject

	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	{{end}}
	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-{{$ResourceName}}-" + name,
				},
			},
		},
	}

	{{range $attrName, $attr := .Attributes}}
	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "{{$attrName}}"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-{{$ResourceName}}-" + name + "-{{$attr}}",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"{{$attr}}"] = output
	{{end}}

	return
}

// Parse{{$ResourceName}} validator
func (resource {{$ResourceName}}) Validate() []error {
	return resource.Properties.Validate()
}

// Parse{{$ResourceName}}Properties validator
func (resource {{$ResourceName}}Properties) Validate() []error {
	errors := []error{}
	{{- range $validator := .ValidatorStrings}}
	{{$validator}}
	{{- end}}
	return errors
}
`

const validatorTemplate = `
	{{- if .PrimitiveType -}}
	if resource.{{.Name}} == nil {
		errors = append(errors, fmt.Errorf("Missing required field '{{.Name}}'"))
	}
	{{- else if .ListMapType -}}
	if resource.{{.Name}} == nil {
		errors = append(errors, fmt.Errorf("Missing required field '{{.Name}}'"))
	} else {
		errors = append(errors, resource.{{.Name}}.Validate()...)
	}
	{{- else -}}
	if resource.{{.Name}} == nil {
		errors = append(errors, fmt.Errorf("Missing required field '{{.Name}}'"))
	}
	{{- end -}}
`
