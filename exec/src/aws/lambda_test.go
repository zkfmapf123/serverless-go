package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_PROFILE = "default"

func TestGetLambdaList(t *testing.T) {
	client := New(TEST_PROFILE)
	list := client.GetLambdaList()

	if list == nil {
		t.Errorf("%v is error", list)
	}
}

func TestIsExist(t *testing.T) {
	client := New(TEST_PROFILE)

	isExist_1 := client.IsExistLambda("bucket")
	isExist_2 := client.IsExistLambda("bucket-aaa")

	assert.Equal(t, isExist_1, true)
	assert.Equal(t, isExist_2, false)
}
