package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Bucket_RedirectRule struct {
	HostName             interface{} `yaml:"HostName,omitempty"`
	HttpRedirectCode     interface{} `yaml:"HttpRedirectCode,omitempty"`
	Protocol             interface{} `yaml:"Protocol,omitempty"`
	ReplaceKeyPrefixWith interface{} `yaml:"ReplaceKeyPrefixWith,omitempty"`
	ReplaceKeyWith       interface{} `yaml:"ReplaceKeyWith,omitempty"`
}

func (resource Bucket_RedirectRule) Validate() []error {
	errs := []error{}

	return errs
}