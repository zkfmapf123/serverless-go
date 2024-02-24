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
		path, _ := os.Getwd()
		ymlConfig := utils.GetYmlProperties[FunctionConfig](fmt.Sprintf("%s/config.yml", functionPath))
		globalConfig := utils.GetYmlProperties[GlobalConfig](fmt.Sprintf("%s/agent.yml", path))

		fmt.Println(globalConfig)

		profile, fnName, s3Name := viper.GetString("profile"), ymlConfig.Config.FunctionName, ymlConfig.Config.StateS3Bucket
		lambda, s3 := aws.NewLambda(profile), aws.NewS3(profile)

		// inspect configs
		if lambda.API.IsExist(fnName) {
			log.Fatalln(fmt.Sprintf("%s is Already Exist Lambda Function", ymlConfig.Config.FunctionName))
		}

		// [x] S3가 없으면 생성해야 함
		if !s3.API.IsExist(s3Name) {
			isCreate := s3.API.Create(aws.S3Info{
				Name:   s3Name,
				Region: globalConfig.Config.Region,
			})

			fmt.Println(isCreate)
		}

		// [x] Create LambdaFunction
		lambda.API.Create(aws.LambdaInfo{
			FunctionName: ymlConfig.Config.FunctionName,
			HandlerName:  "bootstrap",
			IamRoleArn:   ymlConfig.Config.RoleARN,
			DeployPath:   functionPath,
		})

	},
}

func init() {
	interaction.Clear()
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
