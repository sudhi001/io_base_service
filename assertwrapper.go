package io_base_service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Wrapper around assert.Equal
func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}

// Wrapper around assert.NotEqual
func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, expected, actual, msgAndArgs...)
}

// Wrapper around assert.True
func True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return assert.True(t, value, msgAndArgs...)
}

// Wrapper around assert.False
func False(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return assert.False(t, value, msgAndArgs...)
}

// Wrapper around assert.Nil
func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return assert.Nil(t, object, msgAndArgs...)
}

// Wrapper around assert.NotNil
func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotNil(t, object, msgAndArgs...)
}

// NotEmpty checks that an object is not empty
func NotEmpty(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEmpty(t, object, msgAndArgs...)
}
