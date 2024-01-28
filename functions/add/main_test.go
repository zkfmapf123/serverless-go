package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	r1, r2 := add(1, 2), add(10, 20)

	assert.Equal(t, r1, 3)
	assert.Equal(t, r2, 30)
}
