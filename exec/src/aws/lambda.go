package aws

import (
	"context"

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
}

type LambdaConfig struct {
	config AWSConfig
}

func NewLambda(profile string) LambdaConfig {
	return &LambdaHandler{
		handler: New(profile),
	}
}

func (c AWSConfig) IsExist(name string) bool {

	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	}
	_, err := c.lambda.GetFunction(context.TODO(), input)
	if err != nil {
		return false
	}

	return true
}

func (c AWSConfig) GetList() map[string]LambdaInfo {

	listInput := &lambda.ListFunctionsInput{}
	res, err := c.lambda.ListFunctions(context.TODO(), listInput)
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

func getEnvSize(t *types.EnvironmentResponse) int {

	if t == nil {
		return 0
	}

	return len(t.Variables)
}
