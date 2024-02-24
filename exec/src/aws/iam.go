package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IamInfo struct {
	Arn  string
	Desc string
}

type IamConfig struct {
	config AWSConfig
}

type NewIamAPI struct {
	API IAWS[IamInfo]
}

func NewIAM(profile string) NewIamAPI {
	return NewIamAPI{
		API: IamConfig{
			config: New(profile),
		},
	}
}

func (i IamConfig) Create(info IamInfo) bool {
	return false
}

func (i IamConfig) IsExist(name string) bool {

	return false
}

func (i IamConfig) GetList() map[string]IamInfo {

	return nil
}

func (i IamConfig) Retrieve(name string) map[string]IamInfo {

	res, err := i.config.iam.GetRole(context.TODO(),
		&iam.GetRoleInput{RoleName: aws.String(name)},
	)

	if err != nil {
		return nil
	}

	return map[string]IamInfo{
		name: {
			Arn:  *res.Role.Arn,
			Desc: *res.Role.Description,
		},
	}
}

func (i IamConfig) Delete(name string) error {
	return nil
}
