package aws

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

type HandlerConfigInfo struct {
	Timeout    string
	MemorySize string
	// Runtime    string ## deprecated
}

type LambdaInfo struct {
	Desc        string
	Env         int // envSize
	Size        float64
	LastUpdated string

	// Create Config
	FunctionName  string
	HandlerName   string
	IamRoleArn    string
	DeployPath    string
	EnvList       map[string]interface{}
	TagList       map[string]interface{}
	HandlerConfig HandlerConfigInfo
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

	timeout, _ := strconv.Atoi(info.HandlerConfig.Timeout)
	memorySize, _ := strconv.Atoi(info.HandlerConfig.MemorySize)

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
		Environment: &types.Environment{
			Variables: makeLamdaConfigValues(info.EnvList, false),
		},
		Tags:       makeLamdaConfigValues(info.TagList, true),
		Timeout:    aws.Int32(int32(timeout)),
		MemorySize: aws.Int32(int32(memorySize)),
	})

	if err != nil {
		panic(err)
	}

	return true
}

// [x] 간단한 배포 (함수 코드만)
// [ ] 세부사항 변경 배포
// [ ] S3 Versioning
func (l lambdaConfig) Deploy(info LambdaInfo) error {
	lambdaName := l.Retrieve(info.FunctionName)
	if lambdaName[info.FunctionName].FunctionName == "" {
		return fmt.Errorf("not exists %s", info.FunctionName)
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/%s", info.DeployPath, "bootstrap.zip"))
	if err != nil {
		panic(err)
	}

	_, err = l.config.lambda.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(info.FunctionName),
		ZipFile:      file,
	})

	return err
}

func (l lambdaConfig) Retrieve(name string) map[string]LambdaInfo {
	_, err := l.config.lambda.GetFunction(context.TODO(), &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	})

	if err != nil {
		return nil
	}

	return map[string]LambdaInfo{
		name: {
			FunctionName: name,
		},
	}
}

func (l lambdaConfig) Delete(name string) error {

	LambdaName := l.Retrieve(name)
	if LambdaName[name].FunctionName == "" {
		return fmt.Errorf("not exists %s", name)
	}

	_, err := l.config.lambda.DeleteFunction(context.TODO(), &lambda.DeleteFunctionInput{
		FunctionName: aws.String(name),
	})

	return err
}

func getEnvSize(t *types.EnvironmentResponse) int {

	if t == nil {
		return 0
	}

	return len(t.Variables)
}

// Envs, Tags
func makeLamdaConfigValues(value map[string]interface{}, isPascal bool) map[string]string {
	m := make(map[string]string)

	for k, v := range value {

		if isPascal {
			k = strings.Title(k)
		} else {
			k = strings.ToUpper(k)
		}

		switch reflect.TypeOf(v).Kind() {
		case reflect.Bool:
			if v.(bool) {
				m[k] = "true"
			} else {
				m[k] = "false"
			}
		case reflect.Int:
			i := strconv.Itoa(v.(int))
			m[k] = i
		case reflect.String:
			m[k] = v.(string)
		default:
			log.Fatalf("%s is invalid type", reflect.TypeOf(k).String())
		}
	}
	return m
}
