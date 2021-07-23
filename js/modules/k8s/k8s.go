package k8s

import (
	"context"
	"felloe/helpers"
	"felloe/js/modules"
	"github.com/dop251/goja"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var moduleConstructors = map[string]func(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object{
	"Deployment": deploymentConstructor,
}

type K8s struct {
	client *kubernetes.Clientset
}

func New() modules.Module {
	k8s := &K8s{
		client: helpers.GetK8sClient(),
	}

	return modules.Module{
		Name:           "k8s",
		ModuleLoader: func(runtime *goja.Runtime, mod *goja.Object) {
			exports := mod.Get("exports").(*goja.Object)

			exports.Set("Client", k8s)
			for constructorName, constructorFunc := range moduleConstructors {
				exports.Set(constructorName, constructorFunc(runtime, mod))
			}
		},
	}
}

func (k *K8s) ListNamespaces() []v1.Namespace {
	namespaces, err := k.client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return namespaces.Items
}

func (k *K8s) ListPods(namespace string) []v1.Pod {
	pods, err := k.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return pods.Items
}

func (k *K8s) ListDeployments(namespace string) []v12.Deployment {
	deployments, err := k.client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return deployments.Items
}

func (k *K8s) ListStatefulSets(namespace string) []v12.StatefulSet {
	statefulSets, err := k.client.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return statefulSets.Items
}

func (k *K8s) ListJobs(namespace string) []v13.Job {
	jobs, err := k.client.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return jobs.Items
}

func (k *K8s) ListCronJobs(namespace string) []v13.CronJob {
	jobs, err := k.client.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return jobs.Items
}

func (k *K8s) ListConfigMaps(namespace string) []v1.ConfigMap {
	configMaps, err := k.client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return configMaps.Items
}

func (k *K8s) ListSecrets(namespace string) []v1.ConfigMap {
	configMaps, err := k.client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return configMaps.Items
}