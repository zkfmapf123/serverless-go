package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_PROFILE = "test"
	TEST_REGION  = "ap-northeast-2"
)

func TestGetLambdaList(t *testing.T) {
	client := NewLambda(TEST_PROFILE, TEST_REGION)
	list := client.API.GetList()

	if list == nil {
		t.Errorf("%v is error", list)
	}
}

// func TestIsExist(t *testing.T) {
// 	client := NewLambda(TEST_PROFILE)

// 	isExist_1 := client.API.IsExist("bucket")
// 	isExist_2 := client.API.IsExist("bucket-aaa")

// 	assert.Equal(t, isExist_1, true)
// 	assert.Equal(t, isExist_2, false)
// }

// func TestMakeEnvValues(t *testing.T) {
// 	m := map[string]interface{}{
// 		"a": "10",
// 		"b": 10,
// 		"c": true,
// 	}

// 	res := makeLamdaConfigValues(m, true)
// 	res2 := makeLamdaConfigValues(m, true)

// 	assert.Equal(t, res["a"], "10")
// 	assert.Equal(t, res2["A"], "10")

// 	assert.Equal(t, res["b"], "10")
// 	assert.Equal(t, res["c"], "true")
// }

func Test_Retrive(t *testing.T) {
	client := NewLambda(TEST_PROFILE, TEST_REGION)
	info := client.API.Retrieve("add_function")

	assert.NotNil(t, info.FunctionName, true)
	assert.NotNil(t, info.RepositoryType, true)
	assert.NotNil(t, info.Role, true)
	assert.NotNil(t, info.LastUpdated, true)
	assert.NotNil(t, info.MemorySize, true)
}
