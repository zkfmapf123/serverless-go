package cmd

type FunctionConfig struct {
	Config struct {
		StateS3Bucket string `yaml:"state_s3_bucket"`
		RoleARN       string `yaml:"role_arn"`
		FunctionName  string `yaml:"function_name"`
	} `yaml:"configs"`

	HandlerConfig struct {
		Timeout    string `yaml:"timeout"`
		MemorySize string `yaml:"memory_size"`
		Runtime    string `yaml:"runtime"`
	} `yaml:"handler_config"`

	Tags map[string]interface{} `yaml:"tags"`
	Envs map[string]interface{} `yaml:"envs"`
}
