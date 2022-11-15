package trycatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	is := assert.New(t)
	var result *Result[string, error]

	is.Nil(result, "Expected result to be nil")
	is.Empty(result, "Expected result to be empty")
	result = Init(result)
	is.NotNil(result, "Expected result to not be empty")
}

// func TestTry(t *testing.T) {
// 	is := assert.New(t)
// 	var result *Result[string, error]

// }

// func TestCatch(t *testing.T) {
// 	is := assert.New(t)
// 	var result *Result[string, error]

// }

// func TestTryCatch(t *testing.T) {
// 	is := assert.New(t)
// 	var result *Result[string, error]

// }
