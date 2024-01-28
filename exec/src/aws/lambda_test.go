package aws

import (
	"testing"
)

var TEST_PROFILE = "zkfmapf123"

func TestGetLambdaList(t *testing.T) {
	client := New(TEST_PROFILE)
	list := client.GetLambdaList()

	if list == nil {
		t.Errorf("%v is error", list)
	}

}
