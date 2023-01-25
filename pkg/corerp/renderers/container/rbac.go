// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package container

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/kubernetes"
	"github.com/project-radius/radius/pkg/resourcekinds"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
)

func makeRBACRole(appName, name, namespace string, resource *datamodel.ContainerResource) *rpv1.OutputResource {
	labels := kubernetes.MakeDescriptiveLabels(appName, resource.Name, resource.Type)

	role := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kubernetes.NormalizeResourceName(name),
			Namespace: namespace,
			Labels:    labels,
		},
		// At this time, we support only secret rbac permission for the namespace.
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				Verbs:     []string{"get", "list"},
			},
		},
	}

	or := rpv1.NewKubernetesOutputResource(
		resourcekinds.KubernetesRole,
		rpv1.LocalIDKubernetesRole,
		role,
		role.ObjectMeta)

	return &or
}

func makeRBACRoleBinding(appName, name, saName, namespace string, resource *datamodel.ContainerResource) *rpv1.OutputResource {
	labels := kubernetes.MakeDescriptiveLabels(appName, resource.Name, resource.Type)

	bindings := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kubernetes.NormalizeResourceName(name),
			Namespace: namespace,
			Labels:    labels,
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			Name:     kubernetes.NormalizeResourceName(name),
			APIGroup: "rbac.authorization.k8s.io",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind: "ServiceAccount",
				Name: saName,
			},
		},
	}

	or := rpv1.NewKubernetesOutputResource(
		resourcekinds.KubernetesRoleBinding,
		rpv1.LocalIDKubernetesRoleBinding,
		bindings,
		bindings.ObjectMeta)

	or.Dependencies = []rpv1.Dependency{
		{
			LocalID: rpv1.LocalIDKubernetesRole,
		},
	}
	return &or
}
