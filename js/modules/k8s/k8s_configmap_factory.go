package k8s

import (
	"felloe/js/modules"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMapFactory struct {
	configMap *v12.ConfigMap
}

func newConfigMap() *v12.ConfigMap {
	return &v12.ConfigMap{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{},
	}
}

func NewConfigMapFactory() modules.Module {
	factory := &ConfigMapFactory{
		configMap: newConfigMap(),
	}

	return modules.Module{
		Name:           "k8s/configmap",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)
			exports.Set("ConfigMapFactory", factory)
		},
	}
}

func (cf *ConfigMapFactory) ApiVersion(apiVersion string) *ConfigMapFactory {
	cf.configMap.APIVersion = apiVersion
	return cf
}

func (cf *ConfigMapFactory) Name(name string) *ConfigMapFactory {
	cf.configMap.Name = name
	return cf
}

func (cf *ConfigMapFactory) Namespace(namespace string) *ConfigMapFactory {
	cf.configMap.Namespace = namespace
	return cf
}

func (cf *ConfigMapFactory) AddAnnotation(key string, value string) *ConfigMapFactory {
	cf.configMap.Annotations[key] = value
	return cf
}

func (cf *ConfigMapFactory) Annotations(annotations map[string]string) *ConfigMapFactory {
	cf.configMap.Annotations = annotations
	return cf
}

func (cf *ConfigMapFactory) Labels(labels map[string]string) *ConfigMapFactory {
	cf.configMap.Labels = labels
	return cf
}

func (cf *ConfigMapFactory) AddData(key string, value string) *ConfigMapFactory {
	cf.configMap.Data[key] = value
	return cf
}

func (cf *ConfigMapFactory) Data(data map[string]string) *ConfigMapFactory {
	cf.configMap.Data = data
	return cf
}

func (cf *ConfigMapFactory) Build() *v12.ConfigMap {
	c := cf.configMap
	cf.configMap = newConfigMap()
	return c
}