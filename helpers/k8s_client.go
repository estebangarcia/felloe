package helpers

import (
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

var onceClient sync.Once
var client *kubernetes.Clientset

func GetK8sClient() *kubernetes.Clientset {
	onceClient.Do(func() {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		configOverrides := &clientcmd.ConfigOverrides{}

		namespace := viper.GetString("namespace")
		kubeConfigContext := viper.GetString("kube-context")

		if namespace != "" {
			configOverrides.Context.Namespace = namespace
		}

		if kubeConfigContext != "" {
			configOverrides.CurrentContext = kubeConfigContext
		}

		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

		config, err := kubeConfig.ClientConfig()
		if err != nil {
			panic(err)
		}
		client, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}
	})

	return client
}
