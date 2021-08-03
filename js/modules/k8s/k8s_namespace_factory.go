package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceFactory struct {
	namespace *v12.Namespace
}

func newNamespace() *v12.Namespace {
	return &v12.Namespace{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},

	}
}

func NewNamespaceFactory() modules.Module {
	factory := &NamespaceFactory{
		namespace: newNamespace(),
	}

	return modules.Module{
		Name:           "k8s/namespace",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("NamespaceFactory", factory)
		},
	}
}

func (nf *NamespaceFactory) Name(name string) *NamespaceFactory {
	nf.namespace.Name = name
	return nf
}

func (nf *NamespaceFactory) AddAnnotation(key string, value string) *NamespaceFactory {
	nf.namespace.Annotations[key] = value
	return nf
}

func (nf *NamespaceFactory) Annotations(annotations map[string]string) *NamespaceFactory {
	nf.namespace.Annotations = annotations
	return nf
}


func (nf *NamespaceFactory) Build() *v12.Namespace{
	n := nf.namespace
	nf.namespace = newNamespace()
	return n
}