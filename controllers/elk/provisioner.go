package controllers

import "github.com/amirhnajafiz/elk-operator/controllers/elk/resources"

func (r *Reconciler) ProvideElasticsearch() error {
	resources.ElasticsearchConfigmapResource(elasticsearchConfigmap, r.namespace)
	resources.ElasticsearchStatefulsetResource(r.namespace)
	resources.ElasticsearchSVCResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideLogstash() error {
	resources.LogstashConfigmapResource(logstashConfigmap, r.namespace)
	resources.LogstashDeploymentResource(r.namespace)
	resources.LogstashSVCResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideFilebeat() error {
	resources.FilebeatConfigmapResource(filebeatConfigmap, r.namespace)
	resources.FilebeatRoleResource(r.namespace)
	resources.FilebeatRoleBindingResource(r.namespace)
	resources.FilebeatServiceAccountResource(r.namespace)
	resources.FilebeatDeamonsetResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideKibana() error {
	resources.KibanaConfigmapResource(kibanaConfigmap, r.namespace)
	resources.KibanaReplicasetResource(r.namespace)
	resources.KibanaSVCResource(r.namespace)

	return nil
}
