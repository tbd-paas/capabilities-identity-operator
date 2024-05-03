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

package v1alpha1

import (
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertIdentityCapability = errors.New("unable to convert to IdentityCapability")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IdentityCapabilitySpec defines the desired state of IdentityCapability.
type IdentityCapabilitySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Required
	//
	//	Namespace to use where underlying identity capability will be deployed.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Aws IdentityCapabilitySpecAws `json:"aws,omitempty"`
}

type IdentityCapabilitySpecAws struct {
	// +kubebuilder:validation:Optional
	PodIdentityWebhook IdentityCapabilitySpecAwsPodIdentityWebhook `json:"podIdentityWebhook,omitempty"`
}

type IdentityCapabilitySpecAwsPodIdentityWebhook struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the AWS pod identity webhook deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="amazon/amazon-eks-pod-identity-webhook:v0.5.3"
	// +kubebuilder:validation:Optional
	// (Default: "amazon/amazon-eks-pod-identity-webhook:v0.5.3")
	//
	//	Image to use for AWS pod identity webhook deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources IdentityCapabilitySpecAwsPodIdentityWebhookResources `json:"resources,omitempty"`
}

type IdentityCapabilitySpecAwsPodIdentityWebhookResources struct {
	// +kubebuilder:validation:Optional
	Requests IdentityCapabilitySpecAwsPodIdentityWebhookResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits IdentityCapabilitySpecAwsPodIdentityWebhookResourcesLimits `json:"limits,omitempty"`
}

type IdentityCapabilitySpecAwsPodIdentityWebhookResourcesRequests struct {
	// +kubebuilder:default="25m"
	// +kubebuilder:validation:Optional
	// (Default: "25m")
	//
	//	CPU requests to use for AWS pod identity webhook deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="32Mi"
	// +kubebuilder:validation:Optional
	// (Default: "32Mi")
	//
	//	Memory requests to use for AWS pod identity webhook deployment.
	Memory string `json:"memory,omitempty"`
}

type IdentityCapabilitySpecAwsPodIdentityWebhookResourcesLimits struct {
	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory limits to use for AWS pod identity webhook deployment.
	Memory string `json:"memory,omitempty"`
}

// IdentityCapabilityStatus defines the observed state of IdentityCapability.
type IdentityCapabilityStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// IdentityCapability is the Schema for the identitycapabilities API.
type IdentityCapability struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityCapabilitySpec   `json:"spec,omitempty"`
	Status            IdentityCapabilityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityCapabilityList contains a list of IdentityCapability.
type IdentityCapabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityCapability `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *IdentityCapability) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *IdentityCapability) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *IdentityCapability) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *IdentityCapability) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *IdentityCapability) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *IdentityCapability) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *IdentityCapability) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *IdentityCapability) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*IdentityCapability) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*IdentityCapability) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("IdentityCapability")
}

func init() {
	SchemeBuilder.Register(&IdentityCapability{}, &IdentityCapabilityList{})
}
