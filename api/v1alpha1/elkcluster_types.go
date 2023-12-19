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

// ElkClusterSpec defines the desired state of ElkCluster
type ElkClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Replicas is the number of Elasticsearch clusters
	// +kubebuilder:validation:Maximum:=3
	// +kubebuilder:validation:Minimum:=1
	Replicas int `json:"replica,omitempty"`

	// Dashboard is a flag to create a Kibana dashboard
	// +kubebuilder:validation:Optional
	Dashboard bool `json:"dashboard,omitempty"`
}

// ElkClusterStatus defines the observed state of ElkCluster
type ElkClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ElkCluster is the Schema for the elkclusters API
type ElkCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElkClusterSpec   `json:"spec,omitempty"`
	Status ElkClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ElkClusterList contains a list of ElkCluster
type ElkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElkCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElkCluster{}, &ElkClusterList{})
}
