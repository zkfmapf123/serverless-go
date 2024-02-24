package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSConfig struct {
	cfg    aws.Config
	lambda *lambda.Client
	s3     *s3.Client
}

type IAWS[T any] interface {
	IsExist(name string) bool
	GetList() map[string]T
	Create(info T) bool
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
		s3:     s3.NewFromConfig(cfg),
	}
}
