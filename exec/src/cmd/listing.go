package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/aws"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/filesystem"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/interaction"
)

var listingCmd = &cobra.Command{
	Use:   "li",
	Short: "List Lambda Function",
	Long:  "List Lambda",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := aws.NewLambda(viper.GetString("profile"))
		list := cfg.API.GetList()
		if len(list) == 0 {
			fmt.Println("Lambda List is Not Exists")
			os.Exit(0)
		}

		filesystem.PrintTable[aws.LambdaInfo](list, []string{"Name", "Desc", "Env", "Size", "Last Updated"},
			func(k string, v aws.LambdaInfo) []string {
				envString := strconv.Itoa(v.Env)
				sizeString := strconv.Itoa(int(v.Size))

				return []string{k, v.Desc, envString, fmt.Sprintf("%s MB", sizeString), v.LastUpdated}
			})
	},
}

func init() {
	interaction.Clear()
	rollbackCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listingCmd)
}
