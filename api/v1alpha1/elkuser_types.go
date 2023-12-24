/*
Copyright 2023.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ElkUserSpec defines the desired state of ElkUser
type ElkUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Username of the Elk cluster user
	// +kubebuilder:validation:MaxLength:=16
	Username string `json:"username,omitempty"`

	// Password of the Elk cluster user
	// +kubebuilder:validation:MaxLength:=16
	// +kubebuilder:validation:MinLength:=8
	Password string `json:"password,omitempty"`

	// Roles of the Elk cluster user
	Roles []string `json:"roles,omitempty"`
}

// ElkUserStatus defines the observed state of ElkUser
type ElkUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ElkUser is the Schema for the elkusers API
type ElkUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElkUserSpec   `json:"spec,omitempty"`
	Status ElkUserStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ElkUserList contains a list of ElkUser
type ElkUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElkUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElkUser{}, &ElkUserList{})
}
