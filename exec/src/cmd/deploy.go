package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "de",
	Short: "deploy Lambda Function",
	Long:  "deploy Lambda",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy lambda")
	},
}

func init() {
	deployCmd.Flags().String("deploy", "deploy", "")
	rootCmd.AddCommand(deployCmd)
}
