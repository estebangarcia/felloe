package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
)

type Namespace struct {
	*v12.Namespace
}

func namespaceConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewNamespace(call.Arguments...)).ToObject(rt)
	}
}

func NewNamespace(args ...goja.Value) *v12.Namespace {
	namespaceName := ""

	if len(args) < 1 {
		panic("expected namespace name")
	}

	namespaceName = args[0].String()

	namespace := newNamespace()
	namespace.Name = namespaceName

	return namespace
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Namespace) String() string {
	return n.String()
}