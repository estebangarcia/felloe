package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
)

type StatefulSet struct {
	*v1.StatefulSet
}

func statefulSetConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewStatefulSet(call.Arguments...)).ToObject(rt)
	}
}

func NewStatefulSet(args ...goja.Value) *v1.StatefulSet {
	if len(args) == 0 {
		return newStatefulSet()
	}

	statefulSetName := ""
	statefulSetImage := ""

	if len(args) < 2 {
		panic("expected pod name and image")
	}

	statefulSetName = args[0].String()
	statefulSetImage = args[1].String()

	statefulSet := newStatefulSet()
	statefulSet.Name = statefulSetName
	statefulSet.Spec.Template.Spec.Containers = []v12.Container{
		{Name: "app", Image: statefulSetImage},
	}

	return statefulSet
}

func (s *StatefulSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *StatefulSet) String() string {
	return s.String()
}