package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
)

type Pod struct {
	*v12.Pod
}

func podConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewPod(call.Arguments...)).ToObject(rt)
	}
}

func NewPod(args ...goja.Value) *v12.Pod {
	if len(args) == 0 {
		return newPod()
	}

	podName := ""
	image := ""

	if len(args) < 2 {
		panic("expected pod name and image")
	}

	podName = args[0].String()
	image = args[1].String()

	pod := newPod()
	pod.Name = podName
	pod.Spec.Containers = []v12.Container{{Name: "app", Image: image}}

	return pod
}

func (p *Pod) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Pod) String() string {
	return p.String()
}