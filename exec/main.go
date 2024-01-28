package main

import "github.com/zkfmapf123/serverless-go-deploy-agent/cmd"

func main() {

	cmd.Init()
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
