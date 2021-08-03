package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	namespace := &v12.Namespace{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}

	return namespace
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Namespace) String() string {
	return n.String()
}