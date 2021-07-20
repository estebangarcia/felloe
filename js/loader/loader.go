package loader

import (
	"errors"
	"felloe/compiler"
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
