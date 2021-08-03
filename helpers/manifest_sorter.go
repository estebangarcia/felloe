package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
)

var KindOrder = []string{
	"Namespace",
	"NetworkPolicy",
	"ResourceQuota",
	"LimitRange",
	"PodSecurityPolicy",
	"PodDisruptionBudget",
	"ServiceAccount",
	"Secret",
	"SecretList",
	"ConfigMap",
	"StorageClass",
	"PersistentVolume",
	"PersistentVolumeClaim",
	"CustomResourceDefinition",
	"ClusterRole",
	"ClusterRoleList",
	"ClusterRoleBinding",
	"ClusterRoleBindingList",
	"Role",
	"RoleList",
	"RoleBinding",
	"RoleBindingList",
	"Service",
	"DaemonSet",
	"Pod",
	"ReplicationController",
	"ReplicaSet",
	"Deployment",
	"HorizontalPodAutoscaler",
	"StatefulSet",
	"Job",
	"CronJob",
	"Ingress",
	"APIService",
}

type GenericManifest struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec interface{} `json:"spec,omitempty"`
}

func SortManifests(manifests []GenericManifest) []GenericManifest {
	m := manifests
	sort.SliceStable(m, func(i, j int) bool {
		return lessByKind(m[i].Kind, m[j].Kind)
	})

	return m
}

func lessByKind(kindA string, kindB string) bool {
	ordering := make(map[string]int, len(KindOrder))
	for v, k := range KindOrder {
		ordering[k] = v
	}

	first := ordering[kindA]
	second := ordering[kindB]

	return first < second

}