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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/inftyai/llmaz/api/core/v1alpha1"
)

// ModelReferApplyConfiguration represents a declarative configuration of the ModelRefer type for use
// with apply.
type ModelReferApplyConfiguration struct {
	Name *corev1alpha1.ModelName `json:"name,omitempty"`
	Role *corev1alpha1.ModelRole `json:"role,omitempty"`
}

// ModelReferApplyConfiguration constructs a declarative configuration of the ModelRefer type for use with
// apply.
func ModelRefer() *ModelReferApplyConfiguration {
	return &ModelReferApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ModelReferApplyConfiguration) WithName(value corev1alpha1.ModelName) *ModelReferApplyConfiguration {
	b.Name = &value
	return b
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *ModelReferApplyConfiguration) WithRole(value corev1alpha1.ModelRole) *ModelReferApplyConfiguration {
	b.Role = &value
	return b
}
