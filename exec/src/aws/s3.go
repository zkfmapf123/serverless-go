package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Info struct {
	Name string
}

type S3Config struct {
	config AWSConfig
}

type NewS3API struct {
	API IAWS[S3Info]
}

func NewS3(profile string) NewS3API {
	return NewS3API{
		API: S3Config{
			config: New(profile),
		},
	}
}

func (s S3Config) IsExist(name string) bool {

	for _, v := range s.GetList() {
		if v.Name == name {
			return true
		}
	}

	return false
}

func (s S3Config) GetList() map[string]S3Info {
	input := &s3.ListBucketsInput{}
	s3Buckets, err := s.config.s3.ListBuckets(context.TODO(), input)
	if err != nil {
		panic(err)
	}

	s3Items := make(map[string]S3Info)
	for _, v := range s3Buckets.Buckets {
		name := *v.Name
		s3Items[name] = S3Info{
			Name: name,
		}

	}

	return s3Items
}
