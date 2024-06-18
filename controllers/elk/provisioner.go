package controllers

import (
	"context"

	"github.com/amirhnajafiz/elk-operator/controllers/elk/resources"
)

func (r *Reconciler) ProvideElasticsearch(ctx context.Context) error {
	if err := r.Create(ctx, resources.ElasticsearchConfigmapResource(elasticsearchConfigmap, r.namespace)); err != nil {
		return err
	}

	resources.ElasticsearchStatefulsetResource(r.namespace)
	resources.ElasticsearchSVCResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideLogstash(ctx context.Context) error {
	if err := r.Create(ctx, resources.LogstashConfigmapResource(logstashConfigmap, r.namespace)); err != nil {
		return err
	}

	resources.LogstashDeploymentResource(r.namespace)
	resources.LogstashSVCResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideFilebeat(ctx context.Context) error {
	if err := r.Create(ctx, resources.FilebeatConfigmapResource(filebeatConfigmap, r.namespace)); err != nil {
		return err
	}

	resources.FilebeatRoleResource(r.namespace)
	resources.FilebeatRoleBindingResource(r.namespace)
	resources.FilebeatServiceAccountResource(r.namespace)
	resources.FilebeatDeamonsetResource(r.namespace)

	return nil
}

func (r *Reconciler) ProvideKibana(ctx context.Context) error {
	if err := r.Create(ctx, resources.KibanaConfigmapResource(kibanaConfigmap, r.namespace)); err != nil {
		return err
	}

	resources.KibanaReplicasetResource(r.namespace)
	resources.KibanaSVCResource(r.namespace)

	return nil
}
