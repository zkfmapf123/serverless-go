package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_PROFILE = "zkfmapf123"

func TestGetLambdaList(t *testing.T) {
	client := NewLambda(TEST_PROFILE)
	list := client.API.GetList()

	if list == nil {
		t.Errorf("%v is error", list)
	}
}

func TestIsExist(t *testing.T) {
	client := NewLambda(TEST_PROFILE)

	isExist_1 := client.API.IsExist("bucket")
	isExist_2 := client.API.IsExist("bucket-aaa")

	assert.Equal(t, isExist_1, true)
	assert.Equal(t, isExist_2, false)
}
