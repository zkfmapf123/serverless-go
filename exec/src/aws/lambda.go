package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

type LambdaInfo struct {
	Desc        string
	Env         int
	Size        float64
	LastUpdated string

	// Create Config
	FunctionName string
	HandlerName  string
	IamRoleArn   string
	DeployPath   string
}

type lambdaConfig struct {
	config AWSConfig
}

type NewLambdaAPI struct {
	API IAWS[LambdaInfo]
}

func NewLambda(profile string) NewLambdaAPI {
	return NewLambdaAPI{
		API: lambdaConfig{
			config: New(profile),
		},
	}
}

func (l lambdaConfig) IsExist(name string) bool {

	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	}
	_, err := l.config.lambda.GetFunction(context.TODO(), input)

	return err == nil
}

func (l lambdaConfig) GetList() map[string]LambdaInfo {

	listInput := &lambda.ListFunctionsInput{}
	res, err := l.config.lambda.ListFunctions(context.TODO(), listInput)
	if err != nil {
		panic(err)
	}

	funcInfo := make(map[string]LambdaInfo)
	for _, fn := range res.Functions {

		funcInfo[*fn.FunctionName] = LambdaInfo{
			Size:        utils.ByteToMB(fn.CodeSize),
			Env:         getEnvSize(fn.Environment),
			Desc:        *fn.Description,
			LastUpdated: *fn.LastModified,
		}
	}

	return funcInfo
}

func (l lambdaConfig) Create(info LambdaInfo) bool {
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", info.DeployPath, "bootstrap.zip"))
	if err != nil {
		panic(err)
	}

	_, err = l.config.lambda.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
		Code:         &types.FunctionCode{ZipFile: file},
		FunctionName: aws.String(info.FunctionName),
		Role:         &info.IamRoleArn,
		Handler:      aws.String(info.HandlerName),
		Publish:      true,
		Runtime:      types.RuntimeProvidedal2023,
		Architectures: []types.Architecture{
			types.ArchitectureArm64,
		},
	})

	if err != nil {
		panic(err)
	}

	return true
}

func (l lambdaConfig) Retrieve(name string) map[string]LambdaInfo {
	return nil
}

func getEnvSize(t *types.EnvironmentResponse) int {

	if t == nil {
		return 0
	}

	return len(t.Variables)
}
