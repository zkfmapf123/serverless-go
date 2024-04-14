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
		ymlConfig := utils.GetYmlProperties[FunctionConfig](fmt.Sprintf("%s/config.yml", functionPath))
		profile, region, fnName := viper.GetString("profile"), viper.GetString("region"), ymlConfig.Config.FunctionName
		lambda := aws.NewLambda(profile, region)

		// inspect configs
		if !lambda.API.IsExist(fnName) {
			log.Fatalf("%s is Not Exist Lambda Function", ymlConfig.Config.FunctionName)
		}

		filesystem.MakeZip(functionPath)
		err := lambda.API.Deploy(aws.LambdaInfo{
			FunctionName: ymlConfig.Config.FunctionName,
			DeployPath:   functionPath,
		})

		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			info := lambda.API.Retrieve(ymlConfig.Config.FunctionName)
			fmt.Printf("FunctionName : %s\nRepositoryType : %s\nRoleArn : %s\nLastModified : %s\nMemorySize : %d\n", info.FunctionName, info.RepositoryType, info.Role, info.LastUpdated, info.MemorySize)
		}()

		fmt.Printf("Deploy Success %s\n", ymlConfig.Config.FunctionName)

	},
}

func init() {
	deployCmd.Flags().String("deploy", "deploy", "")
	rootCmd.AddCommand(deployCmd)
}
