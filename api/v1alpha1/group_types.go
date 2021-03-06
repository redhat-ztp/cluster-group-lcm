/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RemediationStrategySpec defines the remediation policy
type RemediationStrategySpec struct {
	// Canaries defines the list of names of Site objects that should be remediated first when remediateAction is set to enforce
	Canaries       []string `json:"canaries,omitempty"`
	MaxConcurrency int      `json:"maxConcurrency,omitempty"`
}

// GroupPolicyTemplate defines the object definition of a Policy of the Group
type GroupPolicyTemplate struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	ObjectDefinition runtime.RawExtension `json:"objectDefinition,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Sites defines the list of names of Site objects of the Group.
	Sites []string `json:"sites,omitempty"`
	// GroupPolicyTemplates defines the list of Policy object definitions of the Group.
	GroupPolicyTemplates []GroupPolicyTemplate   `json:"groupPolicyTemplates,omitempty"`
	RemediationStrategy  RemediationStrategySpec `json:"remediationStrategy,omitempty"`
	OnFailureAction      string                  `json:"onFailureAction,omitempty"`
	RemediationAction    string                  `json:"remediationAction,omitempty"`
}

// PolicyStatus defines the observed state of a Policy
type PolicyStatus struct {
	Name            string `json:"name,omitempty"`
	ComplianceState string `json:"complianceState,omitempty"`
}

// GroupStatus defines the observed state of Group
type GroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PlacementBindings []string       `json:"placementBindings"`
	PlacementRules    []string       `json:"placementRules"`
	Policies          []PolicyStatus `json:"policies"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Group is the Schema for the groups API
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GroupSpec   `json:"spec,omitempty"`
	Status GroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GroupList contains a list of Group
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
