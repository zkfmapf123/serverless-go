package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrgin(t *testing.T) {

	text := Concat("A", "B", "C", "D")
	assert.Equal(t, text, "ABCD")
}
