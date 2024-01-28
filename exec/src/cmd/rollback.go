package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "ro",
	Short: "rollback Lambda Function Use ECR",
	Long:  "rollback Lambda",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rollback lambda")
	},
}

func init() {
	rollbackCmd.Flags().String("rollback", "rollback", "")
	rootCmd.AddCommand(rollbackCmd)
}
