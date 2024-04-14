package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSConfig struct {
	cfg    aws.Config
	lambda *lambda.Client
	s3     *s3.Client
	iam    *iam.Client
}

type IAWS[T any] interface {
	IsExist(name string) bool
	GetList() map[string]T
	Create(info T) bool
	Retrieve(name string) T
	Delete(name string) error
	Deploy(info T) error
}

func New(profile string, region string) AWSConfig {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	return AWSConfig{
		cfg:    cfg,
		lambda: lambda.NewFromConfig(cfg),
		s3:     s3.NewFromConfig(cfg),
		iam:    iam.NewFromConfig(cfg),
	}
}
