package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ELKSpec defines the desired state of ELK
type ELKSpec struct {
	// Logstash defines the desired state of Logstash deployment
	Logstash LogstashSpec `json:"logstash,omitempty"`
	// Elasticsearch defines the desired state of Elasticsearch deployment
	Elasticsearch ElasticsearchSpec `json:"elasticsearch,omitempty"`
}

// LogstashSpec defines the desired state of Logstash
type LogstashSpec struct {
	// Replicas define the number of logstash instances
	Replicas int `json:"replicas,omitempty"`
}

// ElasticsearchSpec defines the desired state of Elasticsearch
type ElasticsearchSpec struct {
	// Replicas define the number of elasticsearch instances
	Replicas int `json:"replicas,omitempty"`
}

// ELKStatus defines the observed state of ELK
type ELKStatus struct {
	// Configmap stage status (if true then it is built)
	Configmap bool `json:"configmap,omitempty"`
	// Elasticsearch stage status (if true then it is built)
	Elasticsearch bool `json:"elasticsearch,omitempty"`
	// Logstash stage status (if true then it is built)
	Logstash bool `json:"logstash,omitempty"`
	// Filebeats stage status (if true then it is built)
	Filebeats bool `json:"filebeats,omitempty"`
	// Kibana stage status (if true then it is built)
	Kibana bool `json:"kibana,omitempty"`
	// SVC stage status (if true then it is built)
	SVC bool `json:"svc,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ELK is the Schema for the elks API
type ELK struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ELKSpec   `json:"spec,omitempty"`
	Status ELKStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ELKList contains a list of ELK
type ELKList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ELK `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ELK{}, &ELKList{})
}
