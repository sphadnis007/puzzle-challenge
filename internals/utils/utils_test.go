package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRotatedElem(t *testing.T) {
	var output byte

	output = 'f'
	b := GetRotatedElem('a', 5)
	assert.Equal(t, output, b)

	output = 'v'
	b = GetRotatedElem('a', -5)
	assert.Equal(t, output, b)

	output = 'b'
	b = GetRotatedElem('z', 2)
	assert.Equal(t, output, b)

	// get same element
	output = 'a'
	b = GetRotatedElem('a', 26)
	assert.Equal(t, output, b)

	b = GetRotatedElem('a', -26)
	assert.Equal(t, output, b)

	// out of bounds rotation
	output = 'c'
	b = GetRotatedElem('a', 28)
	assert.Equal(t, output, b)

	output = 'y'
	b = GetRotatedElem('a', -28)
	assert.Equal(t, output, b)
}

func TestRotate1DArrayInRightDirection(t *testing.T) {
	s := Rotate1DArray([]byte("abc"), 1)
	assert.Equal(t, []byte("bcd"), s)
}

func TestRotate1DArrayInLeftDirection(t *testing.T) {
	s := Rotate1DArray([]byte("abc"), -1)
	assert.Equal(t, []byte("zab"), s)
}

func TestRotate1DArraySameOutput(t *testing.T) {
	// Rotating 26 times will get us the same string
	s := Rotate1DArray([]byte("abc"), 26)
	assert.Equal(t, []byte("abc"), s)
}
