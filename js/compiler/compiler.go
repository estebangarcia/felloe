package compiler

import (
	"fmt"
	"io/ioutil"
)

var (
	DefaultBabelOpts = map[string]interface{}{
		"plugins": []interface{}{
			[]interface{}{"transform-es2015-template-literals", map[string]interface{}{"loose": false, "spec": false}},
			"transform-es2015-literals",
			"transform-es2015-function-name",
			[]interface{}{"transform-es2015-arrow-functions", map[string]interface{}{"spec": false}},
			[]interface{}{"transform-es2015-classes", map[string]interface{}{"loose": false}},
			"transform-es2015-object-super",
			"transform-es2015-duplicate-keys",
			[]interface{}{"transform-es2015-computed-properties", map[string]interface{}{"loose": false}},
			[]interface{}{"transform-es2015-spread", map[string]interface{}{"loose": false}},
			"transform-es2015-parameters",
			[]interface{}{"transform-es2015-destructuring", map[string]interface{}{"loose": false}},
			[]interface{}{"transform-es2015-modules-commonjs", map[string]interface{}{"loose": false}},
			"transform-exponentiation-operator",
		},
	}
)

type Compiler struct {
	babel *babel
}

func New() (*Compiler, error) {
	babel, err := getBabel()
	if err != nil {
		return nil, fmt.Errorf("error running babel: %v", err.Error())
	}

	return &Compiler{
		babel: babel,
	}, nil
}

func (c Compiler) Transform(src string) (string, error) {
	v, err := c.babel.transform(src, DefaultBabelOpts)
	if err != nil {
		return "", err
	}

	babelTransformOutputObject := v.ToObject(c.babel.vm)

	var code string
	if err = c.babel.vm.ExportTo(babelTransformOutputObject.Get("code"), &code); err != nil {
		return "", err
	}

	return code, nil
}

func CompileScript(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("couldn't open %v", path)
	}

	c, err := New()
	if err != nil {
		return "", err
	}

	compiled, err := c.Transform(string(fileContent))
	if err != nil {
		return "", err
	}

	return compiled, nil
}