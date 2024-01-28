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
	cobra.OnInitialize(initConfig)

}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	// initConfig

	b, err := os.ReadFile("~/.aws/configure")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
