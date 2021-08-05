package compiler

import (
	_ "embed"
	"felloe/js/modules"
	"felloe/logger"
	"github.com/dop251/goja"
	"sync"

	//"felloe/js"
)

//go:embed lib/babel.min.js
var babelSrc string //nolint:gochecknoglobals
var onceBabel sync.Once
var b *babel

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

func getBabel() (*babel, error) {
	var err error

	onceBabel.Do(func() {
		var babelProg *goja.Program
		babelProg, err = compileBabel()

		if err != nil {
			return
		}

		vm := getBabelRuntime()

		if _, err = vm.RunProgram(babelProg); err != nil {
			return
		}

		b = &babel{
			vm: vm,
		}

		var transform goja.Callable
		babel := vm.Get("Babel")

		if err = vm.ExportTo(babel.ToObject(vm).Get("transform"), &transform); err != nil {
			return
		}

		b.transform = func(src string, opts map[string]interface{}) (goja.Value, error) {
			return transform(babel, vm.ToValue(src), vm.ToValue(opts))
		}
	})

	return b, err
}

func getBabelRuntime() *goja.Runtime {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	vm.Set("console", modules.NewJSConsole(logger.GetLogger()))
	return vm
}