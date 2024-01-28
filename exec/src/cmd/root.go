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
	PROFILE = "default"
	REGION  = "ap-northeast-2"
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

	rootCmd.PersistentFlags().StringP("profile", "p", "", "[Required] aws profile")
	rootCmd.PersistentFlags().StringP("region", "r", "", "[Required] aws region")

	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
}

func initConfig() {

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

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configFilePath := filepath.Join(homeDir, ".aws", "credentials")
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	if !strings.Contains(string(b), viper.GetString("profile")) {
		panic(fmt.Sprintf("Not Exist Profile %s", viper.GetString("profile")))
	}
}
