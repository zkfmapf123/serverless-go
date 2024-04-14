package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IamInfo struct {
	Arn  string
	Desc string
	Name string
}

type IamConfig struct {
	config AWSConfig
}

type NewIamAPI struct {
	API IAWS[IamInfo]
}

func NewIAM(profile, region string) NewIamAPI {
	return NewIamAPI{
		API: IamConfig{
			config: New(profile, region),
		},
	}
}

// 작업 중
func (i IamConfig) Create(info IamInfo) bool {

	i.config.iam.CreateRole(context.TODO(), &iam.CreateRoleInput{
		RoleName: &info.Name,
		AssumeRolePolicyDocument: aws.String(`
		{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Principal": {
						"Service": "lambda.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
				}
			]
		}`),
	})

	return false
}

func (i IamConfig) IsExist(name string) bool {

	return false
}

func (i IamConfig) GetList() map[string]IamInfo {

	return nil
}

func (i IamConfig) Retrieve(name string) IamInfo {

	res, err := i.config.iam.GetRole(context.TODO(),
		&iam.GetRoleInput{RoleName: aws.String(name)},
	)

	if err != nil {
		return IamInfo{}
	}

	return IamInfo{
		Arn:  *res.Role.Arn,
		Desc: *res.Role.Description,
		Name: *res.Role.RoleName,
	}
}

func (i IamConfig) Delete(name string) error {
	return nil
}

func (i IamConfig) Deploy(info IamInfo) error {
	return nil
}
