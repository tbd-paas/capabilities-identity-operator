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

package mutate

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	identityv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/identity/v1alpha1"
)

// MutateServiceNamespaceAwsPodIdentityWebhook mutates the Service resource with name aws-pod-identity-webhook.
func MutateServiceNamespaceAwsPodIdentityWebhook(
	original client.Object,
	parent *identityv1alpha1.AWSPodIdentityWebhook,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	// mutation logic goes here

	return []client.Object{original}, nil
}
