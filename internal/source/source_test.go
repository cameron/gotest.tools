package source_test

// using a separate package for test to avoid circular imports with the assert
// package

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	"gotest.tools/internal/source"
)

func TestFormattedCallExprArg_SingleLine(t *testing.T) {
	msg, err := shim("not", "this", "this text")
	assert.NilError(t, err)
	assert.Equal(t, `"this text"`, msg)
}

func TestFormattedCallExprArg_MultiLine(t *testing.T) {
	msg, err := shim(
		"first",
		"second",
		"this text",
	)
	assert.NilError(t, err)
	assert.Equal(t, `"this text"`, msg)
}

func TestFormattedCallExprArg_IfStatement(t *testing.T) {
	if msg, err := shim(
		"first",
		"second",
		"this text",
	); true {
		assert.NilError(t, err)
		assert.Equal(t, `"this text"`, msg)
	}
}

func shim(_, _, _ string) (string, error) {
	return source.FormattedCallExprArg(1, 2)
}

func TestFormattedCallExprArg_InDefer(t *testing.T) {
	cap := &capture{}
	func() {
		fmt.Println()
		defer cap.shim("first", "second")
	}()

	assert.NilError(t, cap.err)
	assert.Equal(t, cap.value, `"second"`)
}

type capture struct {
	value string
	err   error
}

func (c *capture) shim(_, _ string) {
	c.value, c.err = source.FormattedCallExprArg(1, 1)
}

func TestFormattedCallExprArg_InAnonymousDefer(t *testing.T) {
	cap := &capture{}
	func() {
		fmt.Println()
		defer fmt.Println()
		defer func() { cap.shim("first", "second") }()
	}()

	assert.NilError(t, cap.err)
	assert.Equal(t, cap.value, `"second"`)
}
