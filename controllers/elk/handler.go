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
		// update status
		// save it
		// requeue
	case r.instance.Status.Elasticsearch:
		// create configmaps
		// update status
		// save it
		// requeue
	case r.instance.Status.Logstash:
		// create configmaps
		// update status
		// save it
		// requeue
	case r.instance.Status.Filebeats:
		// create configmaps
		// update status
		// save it
		// requeue
	case r.instance.Status.Kibana:
		// create configmaps
		// update status
		// save it
		// requeue
	case r.instance.Status.SVC:
		// create configmaps
		// update status
		// save it
		// requeue
	default:
		// update status
		// save it
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
