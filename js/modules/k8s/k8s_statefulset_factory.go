package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetFactory struct {
	statefulSet *v1.StatefulSet
}

func newStatefulSet() *v1.StatefulSet {
	return &v1.StatefulSet{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "v1",
		},
		Spec: v1.StatefulSetSpec{
		},
	}
}

func NewStatefulSetFactory() modules.Module {
	factory := &StatefulSetFactory{
		statefulSet: newStatefulSet(),
	}

	return modules.Module{
		Name:           "k8s/statefulSet",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("StatefulSetFactory", factory)
		},
	}
}

func (sf *StatefulSetFactory) ApiVersion(apiVersion string) *StatefulSetFactory {
	sf.statefulSet.APIVersion = apiVersion
	return sf
}

func (sf *StatefulSetFactory) Name(name string) *StatefulSetFactory {
	sf.statefulSet.Name = name
	return sf
}

func (sf *StatefulSetFactory) Namespace(namespace string) *StatefulSetFactory {
	sf.statefulSet.Namespace = namespace
	return sf
}

func (sf *StatefulSetFactory) AddAnnotation(key string, value string) *StatefulSetFactory {
	sf.statefulSet.Annotations[key] = value
	return sf
}

func (sf *StatefulSetFactory) Annotations(annotations map[string]string) *StatefulSetFactory {
	sf.statefulSet.Annotations = annotations
	return sf
}

func (sf *StatefulSetFactory) Labels(labels map[string]string) *StatefulSetFactory {
	sf.statefulSet.Labels = labels
	return sf
}

func (sf *StatefulSetFactory) Volumes(volumes []v12.Volume) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.Volumes = volumes
	return sf
}

func (sf *StatefulSetFactory) AddVolume(volume v12.Volume) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.Volumes = append(sf.statefulSet.Spec.Template.Spec.Volumes, volume)
	return sf
}

func (sf *StatefulSetFactory) InitContainers(initContainers []v12.Container) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.InitContainers = initContainers
	return sf
}

func (sf *StatefulSetFactory) AddInitContainer(initContainer v12.Container) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.InitContainers = append(sf.statefulSet.Spec.Template.Spec.InitContainers, initContainer)
	return sf
}

func (sf *StatefulSetFactory) Containers(containers []v12.Container) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.Containers = containers
	return sf
}

func (sf *StatefulSetFactory) AddContainer(container v12.Container) *StatefulSetFactory {
	sf.statefulSet.Spec.Template.Spec.Containers = append(sf.statefulSet.Spec.Template.Spec.Containers, container)
	return sf
}

func (sf *StatefulSetFactory) Build() *v1.StatefulSet {
	p := sf.statefulSet
	sf.statefulSet = newStatefulSet()
	return p
}