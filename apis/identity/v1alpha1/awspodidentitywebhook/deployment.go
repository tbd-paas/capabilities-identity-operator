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

package awspodidentitywebhook

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	identityv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/identity/v1alpha1"
	"github.com/tbd-paas/capabilities-certificates-operator/apis/identity/v1alpha1/awspodidentitywebhook/mutate"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceAwsPodIdentityWebhook creates the Deployment resource with name aws-pod-identity-webhook.
func CreateDeploymentNamespaceAwsPodIdentityWebhook(
	parent *identityv1alpha1.AWSPodIdentityWebhook,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
			"spec": map[string]interface{}{
				// controlled by field: aws.podIdentityWebhook.replicas
				//  Number of replicas to use for the AWS pod identity webhook deployment.
				"replicas": parent.Spec.Aws.PodIdentityWebhook.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "aws-pod-identity-webhook",
						"app.kubernetes.io/instance":  "aws-pod-identity-webhook",
						"app.kubernetes.io/component": "aws-pod-identity-webhook",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                                  "aws-pod-identity-webhook",
							"app.kubernetes.io/component":          "aws-pod-identity-webhook",
							"app.kubernetes.io/name":               "aws-pod-identity-webhook",
							"app.kubernetes.io/instance":           "aws-pod-identity-webhook",
							"capabilities.tbd.io/capability":       "identity",
							"capabilities.tbd.io/version":          "v0.0.1",
							"capabilities.tbd.io/platform-version": "unstable",
							"app.kubernetes.io/version":            "v0.5.3",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/managed-by":         "identity-operator",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "aws-pod-identity-webhook",
						"containers": []interface{}{
							map[string]interface{}{
								"name": "aws-pod-identity-webhook",
								// controlled by field: aws.podIdentityWebhook.image
								//  Image to use for AWS pod identity webhook deployment.
								"image":           parent.Spec.Aws.PodIdentityWebhook.Image,
								"imagePullPolicy": "Always",
								"command": []interface{}{
									"/webhook",
									"--in-cluster=false",
									"--namespace=" + parent.Spec.Namespace + "", //  controlled by field: namespace
									"--service-name=aws-pod-identity-webhook",
									"--annotation-prefix=eks.amazonaws.com",
									"--token-audience=kubernetes.default.svc",
									"--logtostderr",
									"--port=9443",
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "cert",
										"mountPath": "/etc/webhook/certs",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: aws.podIdentityWebhook.resources.requests.cpu
										//  CPU requests to use for AWS pod identity webhook deployment.
										"cpu": parent.Spec.Aws.PodIdentityWebhook.Resources.Requests.Cpu,
										// controlled by field: aws.podIdentityWebhook.resources.requests.memory
										//  Memory requests to use for AWS pod identity webhook deployment.
										"memory": parent.Spec.Aws.PodIdentityWebhook.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: aws.podIdentityWebhook.resources.limits.memory
										//  Memory limits to use for AWS pod identity webhook deployment.
										"memory": parent.Spec.Aws.PodIdentityWebhook.Resources.Limits.Memory,
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "cert",
								"secret": map[string]interface{}{
									"secretName": "aws-pod-identity-webhook-cert",
								},
							},
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"aws-pod-identity-webhook",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
							"kubernetes.io/arch": "arm64",
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}
