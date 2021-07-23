package modules

import (
	"github.com/dop251/goja_nodejs/require"
)

type Module struct {
	Name string
	ModuleLoader require.ModuleLoader
}