package templates

import (
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/symcn/mid-operator/pkg/utils"
)

func ObjectMeta(name string, labels map[string]string, config runtime.Object) metav1.ObjectMeta {
	obj := config.DeepCopyObject()
	objMeta, _ := meta.Accessor(obj)
	ovk := config.GetObjectKind().GroupVersionKind()

	return metav1.ObjectMeta{
		Name:      name,
		Namespace: objMeta.GetNamespace(),
		Labels:    labels,
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion:         ovk.GroupVersion().String(),
				Kind:               ovk.Kind,
				Name:               objMeta.GetName(),
				UID:                objMeta.GetUID(),
				Controller:         utils.BoolPointer(true),
				BlockOwnerDeletion: utils.BoolPointer(true),
			},
		},
	}
}

func ObjectMetaWithAnnotations(name string, labels map[string]string, annotations map[string]string, config runtime.Object) metav1.ObjectMeta {
	o := ObjectMeta(name, labels, config)
	o.Annotations = annotations
	return o
}

func ObjectMetaClusterScope(name string, labels map[string]string, config runtime.Object) metav1.ObjectMeta {
	obj := config.DeepCopyObject()
	objMeta, _ := meta.Accessor(obj)
	ovk := config.GetObjectKind().GroupVersionKind()

	return metav1.ObjectMeta{
		Name:   name,
		Labels: labels,
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion:         ovk.GroupVersion().String(),
				Kind:               ovk.Kind,
				Name:               objMeta.GetName(),
				UID:                objMeta.GetUID(),
				Controller:         utils.BoolPointer(true),
				BlockOwnerDeletion: utils.BoolPointer(true),
			},
		},
	}
}

func ControlPlaneAuthPolicy(istiodEnabled, controlPlaneSecurityEnabled bool) string {
	if !istiodEnabled && controlPlaneSecurityEnabled {
		return "MUTUAL_TLS"
	}
	return "NONE"
}