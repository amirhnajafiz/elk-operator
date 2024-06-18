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

	cmp.Data["elasticsearch.yml"] = `
	cluster:
      name: ${CLUSTER_NAME}
    node:
      master: ${NODE_MASTER}
      data: ${NODE_DATA}
      name: ${NODE_NAME}
      ingest: ${NODE_INGEST}
      max_local_storage_nodes: 1
      attr.box_type: hot

    processors: ${PROCESSORS:1}

    network.host: ${NETWORK_HOST}

    path:
      data: /usr/share/elasticsearch/data
      logs: /usr/share/elasticsearch/logs

    http:
      compression: true

    discovery:
      zen:
        ping.unicast.hosts: ${DISCOVERY_SERVICE}
        minimum_master_nodes: ${NUMBER_OF_MASTERS}
	`

	return cmp
}

func LogstashConfigmapResource(name, namespace string) *corev1.ConfigMap {
	cmp := newConfigmapResource(name, namespace)

	return cmp
}

func FilebeatConfigmapResource(name, namespace string) *corev1.ConfigMap {
	cmp := newConfigmapResource(name, namespace)

	return cmp
}

func KibanaConfigmapResource(name, namespace string) *corev1.ConfigMap {
	cmp := newConfigmapResource(name, namespace)

	return cmp
}
