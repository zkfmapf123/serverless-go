package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/aws"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/filesystem"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

var deployCmd = &cobra.Command{
	Use:   "dep",
	Short: "deploy Lambda Function",
	Long:  "deploy Lambda",
	Run: func(cmd *cobra.Command, args []string) {
		functionPath, isExit := filesystem.SelectBoxDirectory(viper.GetString("path"))
		if isExit {
			os.Exit(0)
		}

		// get yml config
		path, _ := os.Getwd()
		ymlConfig := utils.GetYmlProperties[FunctionConfig](fmt.Sprintf("%s/config.yml", functionPath))
		globalConfig := utils.GetYmlProperties[GlobalConfig](fmt.Sprintf("%s/agent.yml", path))

		profile, fnName, s3Name := viper.GetString("profile"), ymlConfig.Config.FunctionName, ymlConfig.Config.StateS3Bucket
		lambda, s3 := aws.NewLambda(profile), aws.NewS3(profile)

		// inspect configs
		if !lambda.API.IsExist(fnName) {
			log.Fatalf("%s is Not Exist Lambda Function", ymlConfig.Config.FunctionName)
		}

		// [x] S3가 없으면 생성해야 함
		if !s3.API.IsExist(s3Name) {
			s3.API.Create(aws.S3Info{
				Name:   s3Name,
				Region: globalConfig.Config.Region,
			})

			fmt.Printf("%s bucket Create", s3Name)
		}

		filesystem.MakeZip(functionPath)
		err := lambda.API.Deploy(aws.LambdaInfo{
			FunctionName: ymlConfig.Config.FunctionName,
			DeployPath:   functionPath,
		})

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	deployCmd.Flags().String("deploy", "deploy", "")
	rootCmd.AddCommand(deployCmd)
}
