package istiod

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/symcn/mid-operator/pkg/controllers/resources/templates"
	util "github.com/symcn/mid-operator/pkg/utils"
)

func (r *Reconciler) serviceAccount() runtime.Object {
	return &corev1.ServiceAccount{
		ObjectMeta: templates.ObjectMeta(serviceAccountName, istiodLabels, r.Config),
	}
}

func (r *Reconciler) clusterRole() runtime.Object {
	rules := []rbacv1.PolicyRule{
		{
			APIGroups: []string{"networking.istio.io"},
			Resources: []string{"gateways"},
			Verbs:     []string{"create"},
		},
		{
			APIGroups: []string{"config.istio.io", "rbac.istio.io", "security.istio.io", "networking.istio.io", "authentication.istio.io"},
			Resources: []string{"*"},
			Verbs:     []string{"get", "watch", "list"},
		},
		{
			APIGroups: []string{"apiextensions.k8s.io"},
			Resources: []string{"customresourcedefinitions"},
			Verbs:     []string{"get", "watch", "list"},
		},
		{
			APIGroups: []string{"extensions", "apps"},
			Resources: []string{"deployments"},
			Verbs:     []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"endpoints", "pods", "services", "namespaces", "nodes"},
			Verbs:     []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"endpoints", "pods", "services", "namespaces", "nodes"},
			Verbs:     []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"extensions"},
			Resources: []string{"ingresses"},
			Verbs:     []string{"get", "watch", "list"},
		},
		{
			APIGroups: []string{"extensions"},
			Resources: []string{"ingresses/status"},
			Verbs:     []string{"*"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"configmaps"},
			Verbs:     []string{"create", "get", "list", "watch", "update"},
		},
		{
			APIGroups: []string{"certificates.k8s.io"},
			Resources: []string{"certificatesigningrequests", "certificatesigningrequests/approval", "certificatesigningrequests/status"},
			Verbs:     []string{"update", "create", "get", "delete", "watch"},
		},
		{
			APIGroups: []string{"authentication.k8s.io"},
			Resources: []string{"tokenreviews"},
			Verbs:     []string{"create"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"secrets"},
			Verbs:     []string{"create", "get", "watch", "list", "update", "delete"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"serviceaccounts"},
			Verbs:     []string{"get", "watch", "list"},
		},
	}

	if util.PointerToBool(r.Config.Spec.Istiod.Enabled) {
		rules = append(rules, rbacv1.PolicyRule{
			APIGroups: []string{"admissionregistration.k8s.io"},
			Resources: []string{"mutatingwebhookconfigurations", "validatingwebhookconfigurations"},
			Verbs:     []string{"get", "list", "watch", "patch", "update"},
		})
	}

	return &rbacv1.ClusterRole{
		ObjectMeta: templates.ObjectMetaClusterScope(clusterRoleNameIstiod, istiodLabels, r.Config),
		Rules:      rules,
	}
}

func (r *Reconciler) clusterRoleBinding() runtime.Object {
	return &rbacv1.ClusterRoleBinding{
		ObjectMeta: templates.ObjectMetaClusterScope(clusterRoleBindingNameIstiod, istiodLabels, r.Config),
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			APIGroup: "rbac.authorization.k8s.io",
			Name:     clusterRoleNameIstiod,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      serviceAccountName,
				Namespace: r.Config.Namespace,
			},
		},
	}
}
