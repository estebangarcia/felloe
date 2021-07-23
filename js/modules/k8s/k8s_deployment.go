package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
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
	deploymentName := ""
	deploymentImage := ""

	if len(args) < 2 {
		panic("")
	}

	deploymentName = args[0].String()
	deploymentImage = args[1].String()

	dep := &v1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec:       v1.DeploymentSpec{
			Replicas:                pointer.Int32Ptr(1),
			Selector:                &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"release.felloe.io/name": deploymentName,
				},
			},
			Template:                v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"release.felloe.io/name": deploymentName,
					},
				},
				Spec:       v12.PodSpec{
					Containers:                    []v12.Container{
						{
							Name: "app",
							Image: deploymentImage,
						},
					},
				},
			},
		},
	}

	return dep
}

func (d *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Deployment) String() string {
	return d.String()
}