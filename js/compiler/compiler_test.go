package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCompiler(t *testing.T) {
	compiler, err := New()

	assert.NoError(t, err)
	assert.NotNil(t, compiler)
}


func TestTransform(t *testing.T) {
	var script = `a => a + 100;`

	var result = `"use strict";

(function (a) {
  return a + 100;
});`

	compiler, _ := New()

	out, err := compiler.Transform(script)
	assert.NoError(t, err)
	assert.Equal(t, out, result)
}