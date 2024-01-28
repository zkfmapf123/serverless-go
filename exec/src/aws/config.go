package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type AWSConfig struct {
	cfg    aws.Config
	lambda *lambda.Client
}

func New(profile string) AWSConfig {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile))
	if err != nil {
		panic(err)
	}

	return AWSConfig{
		cfg:    cfg,
		lambda: lambda.NewFromConfig(cfg),
	}
}
