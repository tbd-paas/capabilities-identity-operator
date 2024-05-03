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
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	identityv1alpha1 "github.com/tbd-paas/capabilities-identity-operator/apis/identity/v1alpha1"
)

// sampleAWSPodIdentityWebhook is a sample containing all fields
const sampleAWSPodIdentityWebhook = `apiVersion: identity.platform.tbd.io/v1alpha1
kind: AWSPodIdentityWebhook
metadata:
  name: awspodidentitywebhook-sample
spec:
  namespace: "tbd-identity-system"
  replicas: 2
  image: "amazon/amazon-eks-pod-identity-webhook:v0.5.3"
  resources:
    requests:
      cpu: "25m"
      memory: "32Mi"
    limits:
      memory: "64Mi"
`

// sampleAWSPodIdentityWebhookRequired is a sample containing only required fields
const sampleAWSPodIdentityWebhookRequired = `apiVersion: identity.platform.tbd.io/v1alpha1
kind: AWSPodIdentityWebhook
metadata:
  name: awspodidentitywebhook-sample
spec:
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleAWSPodIdentityWebhookRequired
	}

	return sampleAWSPodIdentityWebhook
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj identityv1alpha1.AWSPodIdentityWebhook,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte) ([]client.Object, error) {
	var workloadObj identityv1alpha1.AWSPodIdentityWebhook
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	return Generate(workloadObj, nil, nil)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*identityv1alpha1.AWSPodIdentityWebhook,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateServiceAccountNamespaceAwsPodIdentityWebhook,
	CreateRoleNamespaceAwsPodIdentityWebhook,
	CreateRoleBindingNamespaceAwsPodIdentityWebhook,
	CreateClusterRoleAwsPodIdentityWebhook,
	CreateClusterRoleBindingAwsPodIdentityWebhook,
	CreateCertNamespaceAwsPodIdentityWebhook,
	CreateDeploymentNamespaceAwsPodIdentityWebhook,
	CreateServiceNamespaceAwsPodIdentityWebhook,
	CreateMutatingWebhookNamespaceAwsPodIdentityWebhook,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*identityv1alpha1.AWSPodIdentityWebhook,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){}

func ConvertWorkload(component workload.Workload) (*identityv1alpha1.AWSPodIdentityWebhook, error) {
	p, ok := component.(*identityv1alpha1.AWSPodIdentityWebhook)
	if !ok {
		return nil, identityv1alpha1.ErrUnableToConvertAWSPodIdentityWebhook
	}

	return p, nil
}
