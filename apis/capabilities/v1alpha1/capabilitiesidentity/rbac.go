/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package capabilitiesidentity

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	capabilitiesv1alpha1 "github.com/tbd-paas/capabilities-identity-operator/apis/capabilities/v1alpha1"
	"github.com/tbd-paas/capabilities-identity-operator/apis/capabilities/v1alpha1/capabilitiesidentity/mutate"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceAwsPodIdentityWebhook creates the ServiceAccount resource with name aws-pod-identity-webhook.
func CreateServiceAccountNamespaceAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "aws-pod-identity-webhook",
					"app.kubernetes.io/name":               "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
					"app.kubernetes.io/component":          "aws-pod-identity-webhook",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "v0.5.3",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "identity-operator",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;get;update;patch

// CreateRoleNamespaceAwsPodIdentityWebhook creates the Role resource with name aws-pod-identity-webhook.
func CreateRoleNamespaceAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "aws-pod-identity-webhook",
					"app.kubernetes.io/name":               "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
					"app.kubernetes.io/component":          "aws-pod-identity-webhook",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "v0.5.3",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "identity-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"update",
						"patch",
					},
					"resourceNames": []interface{}{
						"pod-identity-webhook",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceAwsPodIdentityWebhook creates the RoleBinding resource with name aws-pod-identity-webhook.
func CreateRoleBindingNamespaceAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "aws-pod-identity-webhook",
					"app.kubernetes.io/name":               "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
					"app.kubernetes.io/component":          "aws-pod-identity-webhook",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "v0.5.3",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "identity-operator",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "aws-pod-identity-webhook",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "aws-pod-identity-webhook",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;watch;list
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=create;get;list;watch

// CreateClusterRoleAwsPodIdentityWebhook creates the ClusterRole resource with name aws-pod-identity-webhook.
func CreateClusterRoleAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "aws-pod-identity-webhook",
				"labels": map[string]interface{}{
					"app":                                  "aws-pod-identity-webhook",
					"app.kubernetes.io/name":               "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
					"app.kubernetes.io/component":          "aws-pod-identity-webhook",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "v0.5.3",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "identity-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"certificatesigningrequests",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"list",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingAwsPodIdentityWebhook creates the ClusterRoleBinding resource with name aws-pod-identity-webhook.
func CreateClusterRoleBindingAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "aws-pod-identity-webhook",
				"labels": map[string]interface{}{
					"app":                                  "aws-pod-identity-webhook",
					"app.kubernetes.io/name":               "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
					"app.kubernetes.io/component":          "aws-pod-identity-webhook",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "v0.5.3",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "identity-operator",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "aws-pod-identity-webhook",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "aws-pod-identity-webhook",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}
