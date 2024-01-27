package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Lambda Agent",
	Short: "Lambda Agent",
	Long:  "Lambda Agent",
}

func Init() {
	cobra.OnInitialize()
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	fmt.Println(home)

}
