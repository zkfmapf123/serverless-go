package main

import "github.com/zkfmapf123/serverless-go-deploy-agent/src/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
