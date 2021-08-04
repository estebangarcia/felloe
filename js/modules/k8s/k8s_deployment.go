package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
)

type Deployment struct {
	*v1.Deployment
}

func deploymentConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewDeployment(call.Arguments...)).ToObject(rt)
	}
}

func NewDeployment(args ...goja.Value) *v1.Deployment {
	if len(args) == 0 {
		return newDeployment()
	}

	deploymentName := ""
	deploymentImage := ""

	if len(args) < 2 {
		panic("expected name and image")
	}

	deploymentName = args[0].String()
	deploymentImage = args[1].String()

	dep := newDeployment()
	dep.Name = deploymentName
	dep.Spec.Template.Spec.Containers = []v12.Container{
		{Name: "app", Image: deploymentImage},
	}

	return dep
}

func (d *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Deployment) String() string {
	return d.String()
}