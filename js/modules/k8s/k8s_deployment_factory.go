package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

type DeploymentFactory struct {
	deployment *v1.Deployment
}

func newDeployment() *v1.Deployment {
	return &v1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		Spec: v1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(1),
		},
	}
}

func NewDeploymentFactory() modules.Module {
	factory := &DeploymentFactory{
		deployment: newDeployment(),
	}

	return modules.Module{
		Name:           "k8s/deployment",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("DeploymentFactory", factory)
		},
	}
}

func (df *DeploymentFactory) ApiVersion(apiVersion string) *DeploymentFactory {
	df.deployment.APIVersion = apiVersion
	return df
}

func (df *DeploymentFactory) Name(name string) *DeploymentFactory {
	df.deployment.Name = name
	return df
}

func (df *DeploymentFactory) Namespace(namespace string) *DeploymentFactory {
	df.deployment.Namespace = namespace
	return df
}

func (df *DeploymentFactory) AddAnnotation(key string, value string) *DeploymentFactory {
	df.deployment.Annotations[key] = value
	return df
}

func (df *DeploymentFactory) Annotations(annotations map[string]string) *DeploymentFactory {
	df.deployment.Annotations = annotations
	return df
}

func (df *DeploymentFactory) PodLabels(labels map[string]string) *DeploymentFactory {
	df.deployment.Spec.Template.ObjectMeta.Labels = labels
	return df
}

func (df *DeploymentFactory) Replicas(replicas int32) *DeploymentFactory {
	df.deployment.Spec.Replicas = pointer.Int32Ptr(replicas)
	return df
}

func (df *DeploymentFactory) Volumes(volumes []v12.Volume) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.Volumes = volumes
	return df
}

func (df *DeploymentFactory) Selector(selector *metav1.LabelSelector) *DeploymentFactory {
	df.deployment.Spec.Selector = selector
	return df
}

func (df *DeploymentFactory) AddVolume(volume v12.Volume) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.Volumes = append(df.deployment.Spec.Template.Spec.Volumes, volume)
	return df
}

func (df *DeploymentFactory) InitContainers(initContainers []v12.Container) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.InitContainers = initContainers
	return df
}

func (df *DeploymentFactory) AddInitContainer(initContainer v12.Container) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.InitContainers = append(df.deployment.Spec.Template.Spec.Containers, initContainer)
	return df
}

func (df *DeploymentFactory) Containers(containers []v12.Container) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.Containers = containers
	return df
}

func (df *DeploymentFactory) AddContainer(container v12.Container) *DeploymentFactory {
	df.deployment.Spec.Template.Spec.Containers = append(df.deployment.Spec.Template.Spec.Containers, container)
	return df
}

func (df *DeploymentFactory) Build() *v1.Deployment{
	d := df.deployment
	df.deployment = newDeployment()
	return d
}