package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBabelCompile(t *testing.T) {
	babelProg, err := compileBabel()

	assert.NoError(t, err)
	assert.NotNil(t, babelProg)
}

func TestNewBabel(t *testing.T) {
	babel, err := getBabel()

	assert.NoError(t, err)
	assert.NotNil(t, babel.vm)
	assert.NotNil(t, babel.transform)
}