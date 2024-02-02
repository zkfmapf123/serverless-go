package cmd

type FunctionConfig struct {
	Config struct {
		RoleARN      string `yaml:"role_arn"`
		FunctionName string `yaml:"function_name"`
	} `yaml:"configs"`

	Env map[string]interface{} `yaml:"env"`
}
