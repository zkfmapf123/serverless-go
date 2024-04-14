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

// Default
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

	InspectParameter(homeDir)

	path, err := os.Getwd()
	if err != nil {
		panic(path)
	}

	inspectHomeYML(path, YML_PREFIX)
}

func InspectParameter(homeDir string) {

	cond := func(v string) bool {
		return v != ""
	}

	viper.Set("profile", utils.InjectDefaultValueNotExist(viper.GetString("profile"), cond, PROFILE))
	viper.Set("region", utils.InjectDefaultValueNotExist(viper.GetString("region"), cond, REGION))
	viper.Set("path", utils.InjectDefaultValueNotExist(viper.GetString("path"), cond, PATH))

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
