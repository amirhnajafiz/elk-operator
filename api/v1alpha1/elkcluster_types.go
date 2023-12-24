package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ElkClusterSpec defines the desired state of ElkCluster
type ElkClusterSpec struct {
	// Replicas is the number of Elasticsearch clusters
	// +kubebuilder:validation:Maximum:=3
	// +kubebuilder:validation:Minimum:=1
	Replicas int `json:"replica,omitempty"`

	// Dashboard is a flag to create a Kibana dashboard
	// +kubebuilder:validation:Optional
	Dashboard bool `json:"dashboard,omitempty"`

	// Ingress is the address of Elk cluster
	Ingress string `json:"ingress,omitempty"`
}

// ElkClusterStatus defines the observed state of ElkCluster
type ElkClusterStatus struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	Created bool `json:"created,omitempty"`
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
