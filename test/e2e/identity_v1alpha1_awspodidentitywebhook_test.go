//go:build e2e_test
// +build e2e_test

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

package e2e_test

import (
	"fmt"
	"os"

	"github.com/stretchr/testify/require"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	identityv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/identity/v1alpha1"
	"github.com/tbd-paas/capabilities-certificates-operator/apis/identity/v1alpha1/awspodidentitywebhook"
)

// identityv1alpha1AWSPodIdentityWebhook tests
func identityv1alpha1AWSPodIdentityWebhookChildrenFuncs(tester *E2ETest) error {
	// TODO: need to run r.GetResources(request) on the reconciler to get the mutated resources
	if len(awspodidentitywebhook.CreateFuncs) == 0 {
		return nil
	}

	workload, err := awspodidentitywebhook.ConvertWorkload(tester.workload)
	if err != nil {
		return fmt.Errorf("error in workload conversion; %w", err)
	}

	resourceObjects, err := awspodidentitywebhook.Generate(*workload, nil, nil)
	if err != nil {
		return fmt.Errorf("unable to create objects in memory; %w", err)
	}

	tester.children = resourceObjects

	return nil
}

func identityv1alpha1AWSPodIdentityWebhookNewHarness(namespace string) *E2ETest {
	return &E2ETest{
		namespace:          namespace,
		unstructured:       &unstructured.Unstructured{},
		workload:           &identityv1alpha1.AWSPodIdentityWebhook{},
		sampleManifestFile: "../../config/samples/identity_v1alpha1_awspodidentitywebhook.yaml",
		getChildrenFunc:    identityv1alpha1AWSPodIdentityWebhookChildrenFuncs,
		logSyntax:          "controllers.identity.AWSPodIdentityWebhook",
	}
}

func (tester *E2ETest) identityv1alpha1AWSPodIdentityWebhookTest(testSuite *E2EComponentTestSuite) {
	testSuite.suiteConfig.tests = append(testSuite.suiteConfig.tests, tester)
	tester.suiteConfig = &testSuite.suiteConfig
	require.NoErrorf(testSuite.T(), tester.setup(), "failed to setup test")

	// create the custom resource
	require.NoErrorf(testSuite.T(), testCreateCustomResource(tester), "failed to create custom resource")

	// test the deletion of a child object
	require.NoErrorf(testSuite.T(), testDeleteChildResource(tester), "failed to reconcile deletion of a child resource")

	// test the update of a child object
	// TODO: need immutable fields so that we can predict which managed fields we can modify to test reconciliation
	// see https://github.com/nukleros/operator-builder/issues/24

	// test the update of a parent object
	// TODO: need immutable fields so that we can predict which managed fields we can modify to test reconciliation
	// see https://github.com/nukleros/operator-builder/issues/24

	// test that controller logs do not contain errors
	if os.Getenv("DEPLOY_IN_CLUSTER") == "true" {
		require.NoErrorf(testSuite.T(), testControllerLogsNoErrors(tester.suiteConfig, tester.logSyntax), "found errors in controller logs")
	}
}

func (testSuite *E2EComponentTestSuite) Test_identityv1alpha1AWSPodIdentityWebhook() {
	tester := identityv1alpha1AWSPodIdentityWebhookNewHarness("")
	tester.identityv1alpha1AWSPodIdentityWebhookTest(testSuite)
}
