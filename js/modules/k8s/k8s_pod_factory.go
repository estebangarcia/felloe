package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodFactory struct {
	pod *v12.Pod
}

func newPod() *v12.Pod {
	return &v12.Pod{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		Spec: v12.PodSpec{
		},
	}
}

func NewPodFactory() modules.Module {
	factory := &PodFactory{
		pod: newPod(),
	}

	return modules.Module{
		Name:           "k8s/pod",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("PodFactory", factory)
		},
	}
}

func (pf *PodFactory) ApiVersion(apiVersion string) *PodFactory {
	pf.pod.APIVersion = apiVersion
	return pf
}

func (pf *PodFactory) Name(name string) *PodFactory {
	pf.pod.Name = name
	return pf
}

func (pf *PodFactory) Namespace(namespace string) *PodFactory {
	pf.pod.Namespace = namespace
	return pf
}

func (pf *PodFactory) AddAnnotation(key string, value string) *PodFactory {
	pf.pod.Annotations[key] = value
	return pf
}

func (pf *PodFactory) Annotations(annotations map[string]string) *PodFactory {
	pf.pod.Annotations = annotations
	return pf
}

func (pf *PodFactory) Labels(labels map[string]string) *PodFactory {
	pf.pod.Labels = labels
	return pf
}

func (pf *PodFactory) Volumes(volumes []v12.Volume) *PodFactory {
	pf.pod.Spec.Volumes = volumes
	return pf
}

func (pf *PodFactory) AddVolume(volume v12.Volume) *PodFactory {
	pf.pod.Spec.Volumes = append(pf.pod.Spec.Volumes, volume)
	return pf
}

func (pf *PodFactory) InitContainers(initContainers []v12.Container) *PodFactory {
	pf.pod.Spec.InitContainers = initContainers
	return pf
}

func (pf *PodFactory) AddInitContainer(initContainer v12.Container) *PodFactory {
	pf.pod.Spec.InitContainers = append(pf.pod.Spec.InitContainers, initContainer)
	return pf
}

func (pf *PodFactory) Containers(containers []v12.Container) *PodFactory {
	pf.pod.Spec.Containers = containers
	return pf
}

func (pf *PodFactory) AddContainer(container v12.Container) *PodFactory {
	pf.pod.Spec.Containers = append(pf.pod.Spec.Containers, container)
	return pf
}

func (pf *PodFactory) Build() *v12.Pod {
	p := pf.pod
	pf.pod = newPod()
	return p
}