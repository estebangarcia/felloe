package js

import (
	"felloe/js/loader"
	"felloe/js/modules"
	"felloe/js/modules/k8s"
	"felloe/logger"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"sync"
)

var onceVM sync.Once
var vm *goja.Runtime

func InitNativeModules(registry *require.Registry) {
	nativeModules := []modules.Module{
		k8s.New(),
		k8s.NewDeploymentFactory(),
		k8s.NewNamespaceFactory(),
		k8s.NewPodFactory(),
		k8s.NewStatefulSetFactory(),
		k8s.NewConfigMapFactory(),
		k8s.NewServiceFactory(),
	}

	for _, module := range nativeModules {
		registry.RegisterNativeModule("felloe/" + module.Name, module.ModuleLoader)
	}
}

func GetRuntime() *goja.Runtime {
	onceVM.Do(func() {
		registry := require.NewRegistryWithLoader(loader.Load)
		InitNativeModules(registry)

		vm = goja.New()
		vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		vm.Set("console", modules.NewJSConsole(logger.GetLogger()))
		registry.Enable(vm)

		exports := vm.NewObject()
		_ = vm.Set("exports", exports)
	})

	return vm
}