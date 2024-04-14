package cmd

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInspectParamter(t *testing.T) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	viper.Set("profile", "")
	viper.Set("region", "")
	viper.Set("path", "")

	profile, region, path := viper.Get("profile"), viper.Get("region"), viper.Get("path")
	assert.Equal(t, profile, "")
	assert.Equal(t, region, "")
	assert.Equal(t, path, "")

	InspectParameter(homeDir)
	profile, region, path = viper.Get("profile"), viper.Get("region"), viper.Get("path")
	assert.Equal(t, profile, "default")
	assert.Equal(t, region, "ap-northeast-2")
	assert.Equal(t, path, "functions")
}
