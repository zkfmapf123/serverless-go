package cmd

import (
	"fmt"
	"log"
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
		profile, fnName, s3Name := viper.GetString("profile"), ymlConfig.Config.FunctionName, ymlConfig.Config.StateS3Bucket

		lambda, s3 := aws.NewLambda(profile), aws.NewS3(profile)

		if lambda.API.IsExist(fnName) {
			log.Fatalln(fmt.Sprintf("%s is Already Exist Lambda Function", ymlConfig.Config.FunctionName))
		}

		if !s3.API.IsExist(s3Name) {
			log.Fatalln(fmt.Sprintf("%s is Not Exist S3 Bucket", ymlConfig.Config.StateS3Bucket))
		}

		/*
			Terraform 활용해서 만들면 됨
		*/

	},
}

func init() {
	interaction.Clear()
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
