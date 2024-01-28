package cmd

import (
	"github.com/spf13/cobra"
)

var listingCmd = &cobra.Command{
	Use:   "li",
	Short: "List Lambda Function",
	Long:  "List Lambda",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func getLambdaList() {

}

func init() {
	rollbackCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listingCmd)
}
