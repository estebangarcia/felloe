package loader

import (
	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"testing"
)

func TestRequireNativeModule(t *testing.T) {
	const SCRIPT = `
	var m = require("test/m");
	var k = require("test/k");
	m.test() + 	k.test();
	`

	vm := js.New()

	registry := new(require.Registry)
	registry.Enable(vm)

	registry.RegisterNativeModule("test/m", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("test", func(call js.FunctionCall) js.Value {
			return runtime.ToValue(1)
		})
	})

	registry.RegisterNativeModule("test/k", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("test", func(call js.FunctionCall) js.Value {
			return runtime.ToValue(2)
		})
	})

	v, err := vm.RunString(SCRIPT)
	if err != nil {
		t.Fatal(err)
	}

	if !v.StrictEquals(vm.ToValue(3)) {
		t.Fatalf("Unexpected result: %v", v)
	}
}
