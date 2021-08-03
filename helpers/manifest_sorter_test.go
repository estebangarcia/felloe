package helpers

import (
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestManifestSorting(t *testing.T) {
	manifests := []GenericManifest{
		{
			TypeMeta:   metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "Deployment",
			},
		},
		{
			TypeMeta:   metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "StatefulSet",
			},
		},
		{
			TypeMeta:   metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "Namespace",
			},
		},
		{
			TypeMeta:   metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "Ingress",
			},
		},
		{
			TypeMeta:   metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "CronJob",
			},
		},
	}

	sortedManifests := SortManifests(manifests)

	var kinds []string

	for i := range sortedManifests {
		kinds = append(kinds, sortedManifests[i].Kind)
	}

	assert.Equal(t, kinds, []string{"Namespace", "Deployment", "StatefulSet", "CronJob", "Ingress"})

}