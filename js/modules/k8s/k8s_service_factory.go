package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceFactory struct {
	service *v12.Service
}

func newService() *v12.Service {
	return &v12.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		Spec: v12.ServiceSpec{},
	}
}

func NewServiceFactory() modules.Module {
	factory := &ServiceFactory{
		service: newService(),
	}

	return modules.Module{
		Name:           "k8s/service",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("ServiceFactory", factory)
		},
	}
}

func (sf *ServiceFactory) ApiVersion(apiVersion string) *ServiceFactory {
	sf.service.APIVersion = apiVersion
	return sf
}

func (sf *ServiceFactory) Name(name string) *ServiceFactory {
	sf.service.Name = name
	return sf
}

func (sf *ServiceFactory) Namespace(namespace string) *ServiceFactory {
	sf.service.Namespace = namespace
	return sf
}

func (sf *ServiceFactory) AddAnnotation(key string, value string) *ServiceFactory {
	sf.service.Annotations[key] = value
	return sf
}

func (sf *ServiceFactory) Annotations(annotations map[string]string) *ServiceFactory {
	sf.service.Annotations = annotations
	return sf
}

func (sf *ServiceFactory) Labels(labels map[string]string) *ServiceFactory {
	sf.service.Labels = labels
	return sf
}

func (sf *ServiceFactory) AddPort(port v12.ServicePort) *ServiceFactory {
	sf.service.Spec.Ports = append(sf.service.Spec.Ports, port)
	return sf
}

func (sf *ServiceFactory) Ports(ports []v12.ServicePort) *ServiceFactory {
	sf.service.Spec.Ports = ports
	return sf
}

func (sf *ServiceFactory) SessionAffinity(affinity v12.ServiceAffinity) *ServiceFactory {
	sf.service.Spec.SessionAffinity = affinity
	return sf
}

func (sf *ServiceFactory) SessionAffinityConfig(affinityConfig *v12.SessionAffinityConfig) *ServiceFactory {
	sf.service.Spec.SessionAffinityConfig = affinityConfig
	return sf
}

func (sf *ServiceFactory) ClusterIP(clusterIp string) *ServiceFactory {
	sf.service.Spec.ClusterIP = clusterIp
	return sf
}

func (sf *ServiceFactory) Type(typ v12.ServiceType) *ServiceFactory {
	sf.service.Spec.Type = typ
	return sf
}

func (sf *ServiceFactory) Selector(selector map[string]string) *ServiceFactory {
	sf.service.Spec.Selector = selector
	return sf
}

func (sf *ServiceFactory) AddSelector(key string, value string) *ServiceFactory {
	sf.service.Spec.Selector[key] = value
	return sf
}

func (sf *ServiceFactory) Build() *v12.Service {
	s := sf.service
	sf.service = newService()
	return s
}