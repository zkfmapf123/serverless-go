package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

var (
	PROFILE    = "default"
	REGION     = "ap-northeast-2"
	PATH       = "functions"
	YML_PREFIX = "agent"
)

var rootCmd = &cobra.Command{
	Use:   "agent",
	Short: "Lambda Agent",
	Long:  "Lambda Agent",
}

func Execute() error {
	initial()
	return rootCmd.Execute()
}

func initial() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("profile", "p", "", "[Optional] aws profile")
	rootCmd.PersistentFlags().StringP("region", "r", "", "[Optional] aws region")
	rootCmd.PersistentFlags().StringP("path", "f", "", "[Required] function path")

	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}

func initConfig() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	inspectParameter(homeDir, viper.GetString("profile"), viper.GetString("region"))

	path, err := os.Getwd()
	if err != nil {
		panic(path)
	}

	inspectHomeYML(path, YML_PREFIX)
}

func inspectParameter(homeDir, profile, region string) {

	// profile
	viper.Set("profile", utils.InjectDefaultValueNotExist[string](viper.GetString("profile"), func(v string) bool {
		if v != "" {
			return true
		}

		return false
	}, PROFILE))

	// region
	viper.Set("region", utils.InjectDefaultValueNotExist[string](viper.GetString("region"), func(v string) bool {
		if v != "" {
			return true
		}

		return false
	}, REGION))

	// function path
	viper.Set("path", utils.InjectDefaultValueNotExist[string](viper.GetString("path"), func(v string) bool {
		if v != "" {
			return true
		}

		return false
	}, PATH))

	configFilePath := filepath.Join(homeDir, ".aws", "credentials")
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	if !strings.Contains(string(b), viper.GetString("profile")) {
		panic(fmt.Sprintf("Not Exist Profile %s", viper.GetString("profile")))
	}
}

func inspectHomeYML(execPath, ymlPrefix string) {

	_, err := os.Stat(fmt.Sprintf("%s/%s.yml", execPath, ymlPrefix))
	if err != nil {
		panic(fmt.Sprintf("Does Not Exist %s.yml", ymlPrefix))
	}
}
