package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrive(t *testing.T) {

	iam := NewIAM(TEST_PROFILE)

	roleName1, roleName2 := "Basic-Lambda-Role", "asdf"

	info_1 := iam.API.Retrieve(roleName1)
	info_2 := iam.API.Retrieve(roleName2)

	assert.NotEqual(t, info_1[roleName1].Arn, "")
	assert.Equal(t, info_2[roleName2].Arn, "")
}
