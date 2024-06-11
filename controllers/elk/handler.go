package controllers

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Handler is the main method of our reconciler, it calls other methods
// based on the current instance status.
func (r *Reconciler) Handler(ctx context.Context) (ctrl.Result, error) {
	switch {
	case r.instance.Status.Configmap:
		// create configmaps
		if err := r.ProvideConfigmaps(); err != nil {
			r.logger.Error(err, "failed to create configmaps")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.Configmap = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	case r.instance.Status.Elasticsearch:
		// create elasticsearch
		if err := r.ProvideElasticsearch(); err != nil {
			r.logger.Error(err, "failed to create elasticsearch")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.Elasticsearch = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	case r.instance.Status.Logstash:
		// create logstash
		if err := r.ProvideLogstash(); err != nil {
			r.logger.Error(err, "failed to create logstash")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.Logstash = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	case r.instance.Status.Filebeats:
		// create filebeats
		if err := r.ProvideFilebeat(); err != nil {
			r.logger.Error(err, "failed to create filebeat")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.Filebeats = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	case r.instance.Status.Kibana:
		// create kibana
		if err := r.ProvideKibana(); err != nil {
			r.logger.Error(err, "failed to create kibana")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.Kibana = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	case r.instance.Status.SVC:
		// create services
		if err := r.ProvideServices(); err != nil {
			r.logger.Error(err, "failed to create services")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
		// update status
		r.instance.Status.SVC = true
		// save it and requeue
		return r.updateInstance(ctx, true)
	default:
		// update status
		r.instance.Status.Ready = true
		// save it
		return r.updateInstance(ctx, false)
	}
}

// updateInstance method saves the current instance, and ends the concile process
func (r *Reconciler) updateInstance(ctx context.Context, queue bool) (ctrl.Result, error) {
	if err := r.Update(ctx, r.instance); err != nil {
		r.logger.Error(err, "failed to update instance")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// requeue if flag is true
	if queue {
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
