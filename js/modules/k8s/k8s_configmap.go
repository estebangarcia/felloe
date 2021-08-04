package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
)

type ConfigMap struct {
	*v12.ConfigMap
}

func configMapConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewConfigMap(call.Arguments...)).ToObject(rt)
	}
}

func NewConfigMap(args ...goja.Value) *v12.ConfigMap {
	if len(args) == 0 {
		return newConfigMap()
	}

	data := make(map[string]string)

	if len(args) < 2 {
		panic("expected key and value")
	}

	data[args[0].String()] = args[1].String()

	configMap := newConfigMap()
	configMap.Data = data

	return configMap
}

func (c *ConfigMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ConfigMap) String() string {
	return c.String()
}