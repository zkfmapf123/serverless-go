package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type YML struct {
	Config struct {
		RoleARN      string `yaml:"role_arn"`
		FunctionName string `yaml:"function_name"`
	} `yaml:"config"`

	Env map[string]interface{} `yaml:"env"`
}

func TestGetYmlProperties(t *testing.T) {

	m := GetYmlProperties[YML]("./test.yml")

	assert.Equal(t, m.Config.FunctionName, "func_b")
	assert.Equal(t, m.Config.RoleARN, "role_a")
	assert.Equal(t, len(m.Env), 4)
}
