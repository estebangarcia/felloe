package compiler

import (
	_ "embed"
	"felloe/js"
	"felloe/logger"
	"github.com/dop251/goja"
	//"felloe/js"
)

//go:embed lib/babel.min.js
var babelSrc string //nolint:gochecknoglobals

type babel struct {
	vm *goja.Runtime
	transform func(src string, opts map[string]interface{}) (goja.Value, error)
}

func compileBabel() (*goja.Program, error) {
	babelProg, err := goja.Compile("babel.min.js", babelSrc, false)
	if err != nil {
		return nil, err
	}

	return babelProg, nil
}

func newBabel() (*babel, error) {
	babelProg, err := compileBabel()

	if err != nil {
		return nil, err
	}

	vm := goja.New()
	log := logger.GetLogger()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	if err := vm.Set("console", js.NewJSConsole(log)); err != nil {
		return nil, err
	}

	if _, err := vm.RunProgram(babelProg); err != nil {
		return nil, err
	}

	b := &babel {
		vm: vm,
	}

	var transform goja.Callable
	babel := vm.Get("Babel")

	if err := vm.ExportTo(babel.ToObject(vm).Get("transform"), &transform); err != nil {
		return nil, err
	}

	b.transform = func(src string, opts map[string]interface{}) (goja.Value, error) {
		return transform(babel, vm.ToValue(src), vm.ToValue(opts))
	}

	return b, nil
}