package aws

import (
	"context"

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

func (c AWSConfig) GetLambdaList() map[string]LambdaInfo {

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
