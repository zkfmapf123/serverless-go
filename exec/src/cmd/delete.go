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

var deleteCmd = &cobra.Command{
	Use:   "de",
	Short: "Delete Lambda Function",
	Long:  "Delete Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {

		functionPath, isExit := filesystem.SelectBoxDirectory(viper.GetString("path"))
		if isExit {
			os.Exit(0)
		}
		ymlConfig := utils.GetYmlProperties[FunctionConfig](fmt.Sprintf("%s/config.yml", functionPath))
		profile, fnName, _ := viper.GetString("profile"), ymlConfig.Config.FunctionName, ymlConfig.Config.StateS3Bucket

		lambda := aws.NewLambda(profile)
		err := lambda.API.Delete(fnName)
		if err != nil {
			log.Fatalln("Dont't Delete ", err)
		}

		fmt.Println("Delete Function", fnName)
	},
}

func init() {
	deleteCmd.Flags().String("delete", "delete", "")
	rootCmd.AddCommand(deleteCmd)
}
