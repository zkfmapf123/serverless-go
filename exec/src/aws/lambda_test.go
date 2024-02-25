package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_PROFILE = "test"

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

func TestMakeEnvValues(t *testing.T) {
	m := map[string]interface{}{
		"a": "10",
		"b": 10,
		"c": true,
	}

	res := makeLamdaConfigValues(m)
	res2 := makeLamdaConfigValues(m)

	assert.Equal(t, res["a"], "10")
	assert.Equal(t, res2["A"], "10")

	assert.Equal(t, res["b"], "10")
	assert.Equal(t, res["c"], "true")

}
