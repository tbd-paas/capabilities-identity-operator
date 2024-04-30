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

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceAwsPodIdentityWebhook creates the Service resource with name aws-pod-identity-webhook.
func CreateServiceNamespaceAwsPodIdentityWebhook(
	parent *capabilitiesv1alpha1.IdentityCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"annotations": map[string]interface{}{
					"prometheus.io/port":   "9443",
					"prometheus.io/scheme": "https",
					"prometheus.io/scrape": "true",
				},
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
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port":       443,
						"targetPort": 9443,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":      "aws-pod-identity-webhook",
					"app.kubernetes.io/instance":  "aws-pod-identity-webhook",
					"app.kubernetes.io/component": "aws-pod-identity-webhook",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}
