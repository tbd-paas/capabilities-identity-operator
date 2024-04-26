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

package capabilities

import (
	v1alpha1capabilities "github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities/v1alpha1"
	v1alpha1identitycapability "github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities/v1alpha1/capabilitiesidentity"
)

// Code generated by operator-builder. DO NOT EDIT.

// IdentityCapabilityLatestGroupVersion returns the latest group version object associated with this
// particular kind.
var IdentityCapabilityLatestGroupVersion = v1alpha1capabilities.GroupVersion

// IdentityCapabilityLatestSample returns the latest sample manifest associated with this
// particular kind.
var IdentityCapabilityLatestSample = v1alpha1identitycapability.Sample(false)
