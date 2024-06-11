package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ELKSpec defines the desired state of ELK
type ELKSpec struct {
	Logstash      LogstashSpec      `json:"logstash"`
	Elasticsearch ElasticsearchSpec `json:"elasticsearch"`
}

type LogstashSpec struct {
	Replicas int `json:"replicas,omitempty"`
}

type ElasticsearchSpec struct {
	Replicas int `json:"replicas,omitempty"`
}

// ELKStatus defines the observed state of ELK
type ELKStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Configmap      bool `json:"configmap"`
	Elasticcsearch bool `json:"elastic"`
	Logstash       bool `json:"logstash"`
	Filebeats      bool `json:"filebeats"`
	Kibana         bool `json:"kibana"`
	SVC            bool `json:"svc"`
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
