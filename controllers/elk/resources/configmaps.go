package resources

import (
	corev1 "k8s.io/api/core/v1"
)

// newConfigmapResource returns a default configmap resource
func newConfigmapResource(name, namespace string) *corev1.ConfigMap {
	cmp := &corev1.ConfigMap{}

	cmp.Namespace = namespace
	cmp.Name = name

	return cmp
}

func ElasticsearchConfigmapResource(name, namespace string) *corev1.ConfigMap {
	cmp := newConfigmapResource(name, namespace)

	return cmp
}

func LogstashConfigmapResource(namespace string) {

}

func FilebeatConfigmapResource(namespace string) {

}

func KibanaConfigmapResource(namespace string) {

}
