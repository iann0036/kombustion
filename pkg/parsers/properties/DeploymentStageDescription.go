package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentStageDescription Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html
type DeploymentStageDescription struct {
	CacheClusterEnabled  interface{} `yaml:"CacheClusterEnabled,omitempty"`
	CacheClusterSize     interface{} `yaml:"CacheClusterSize,omitempty"`
	CacheDataEncrypted   interface{} `yaml:"CacheDataEncrypted,omitempty"`
	CacheTtlInSeconds    interface{} `yaml:"CacheTtlInSeconds,omitempty"`
	CachingEnabled       interface{} `yaml:"CachingEnabled,omitempty"`
	ClientCertificateId  interface{} `yaml:"ClientCertificateId,omitempty"`
	DataTraceEnabled     interface{} `yaml:"DataTraceEnabled,omitempty"`
	Description          interface{} `yaml:"Description,omitempty"`
	DocumentationVersion interface{} `yaml:"DocumentationVersion,omitempty"`
	LoggingLevel         interface{} `yaml:"LoggingLevel,omitempty"`
	MetricsEnabled       interface{} `yaml:"MetricsEnabled,omitempty"`
	ThrottlingBurstLimit interface{} `yaml:"ThrottlingBurstLimit,omitempty"`
	ThrottlingRateLimit  interface{} `yaml:"ThrottlingRateLimit,omitempty"`
	Variables            interface{} `yaml:"Variables,omitempty"`
	MethodSettings       interface{} `yaml:"MethodSettings,omitempty"`
}

// DeploymentStageDescription validation
func (resource DeploymentStageDescription) Validate() []error {
	errors := []error{}

	return errors
}
