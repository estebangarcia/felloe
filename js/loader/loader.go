package loader

import (
	"errors"
	"felloe/compiler"
	"felloe/js/modules"
	"felloe/js/modules/k8s"
	"github.com/dop251/goja_nodejs/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func Load(path string) ([]byte, error) {
	modulePath := path

	if !checkFileExist(modulePath) && checkFileExist(modulePath + ".js") {
		modulePath = modulePath + ".js"
	}

	c, err := compiler.New()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath.FromSlash(modulePath))
	if err != nil {
		if os.IsNotExist(err) || errors.Is(err, syscall.EISDIR) {
			err = require.ModuleFileDoesNotExistError
		}
	}

	src, err := c.Transform(string(data))
	return []byte(src), err
}

func checkFileExist(path string) bool {
	_, err := os.Stat(filepath.FromSlash(path))
	return !os.IsNotExist(err)
}

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