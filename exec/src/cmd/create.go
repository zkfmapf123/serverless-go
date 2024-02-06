package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/aws"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/filesystem"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/interaction"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

var createCmd = &cobra.Command{
	Use:   "cr",
	Short: "Create Lambda Function",
	Long:  "Create Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {

		functionPath, isExit := filesystem.SelectBoxDirectory(viper.GetString("path"))
		if isExit {
			os.Exit(0)
		}

		// get yml config
		ymlConfig := utils.GetYmlProperties[FunctionConfig](fmt.Sprintf("%s/config.yml", functionPath))

		// inspect lambda
		cfg := aws.New(viper.GetString("profile"))
		if cfg.IsExistLambda(ymlConfig.Config.FunctionName) {
			fmt.Printf("%s is Already Exist", ymlConfig.Config.FunctionName)
		}

		/*
			Terraform 활용해서 만들면 됨
		*/

	},
}

func createLambdaFunction() {

}

func init() {
	interaction.Clear()
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
