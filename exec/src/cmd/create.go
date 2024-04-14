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
		profile, region, fnName := viper.GetString("profile"), viper.GetString("region"), ymlConfig.Config.FunctionName
		lambda, iam := aws.NewLambda(profile, region), aws.NewIAM(profile, region)

		// inspect configs
		isExist := false
		if lambda.API.IsExist(fnName) {
			log.Printf("%s is Already Exist Lambda Function", ymlConfig.Config.FunctionName)
			isExist = true
		}

		// [ ] Get IAM Role ARN
		roleArn := iam.API.Retrieve(ymlConfig.Config.RoleARN)
		if roleArn.Arn == "" {
			log.Fatalf("%s is Not Exist", roleArn.Name)
		}

		if !isExist {
			// [x] Create LambdaFunction
			filesystem.MakeZip(functionPath)
			lambda.API.Create(aws.LambdaInfo{
				FunctionName: ymlConfig.Config.FunctionName,
				HandlerName:  "bootstrap",
				IamRoleArn:   roleArn.Arn,
				DeployPath:   functionPath,
				EnvList:      ymlConfig.Envs,
				TagList:      ymlConfig.Tags,
				HandlerConfig: aws.HandlerConfigInfo{
					Timeout:    ymlConfig.HandlerConfig.Timeout,
					MemorySize: ymlConfig.HandlerConfig.MemorySize,
				},
			})
		}

		defer func() {

			info := lambda.API.Retrieve(ymlConfig.Config.FunctionName)
			fmt.Printf(`FunctionName : %s\nRepositoryType : %s\nRoleArn : %s\nLastModified : %s\nMemorySize : %d\n`, info.FunctionName, info.RepositoryType, info.Role, info.LastUpdated, info.MemorySize)
		}()

		fmt.Printf("Deploy Success %s\n", ymlConfig.Config.FunctionName)
	},
}

func init() {
	interaction.Clear()
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
