package elkcluster

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *Reconciler) Modify(ctx context.Context) (ctrl.Result, error) {
	deployment := &v1.Deployment{}
	name := fmt.Sprintf("%s-deployment", r.cluster.Name)

	switch err := r.Get(ctx, client.ObjectKey{Name: name}, deployment); {
	case apierrors.IsNotFound(err):
		// resource not found
		r.logger.Info(fmt.Sprintf("ElkCluster deployment %s not found!", name))
		return ctrl.Result{}, nil
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default:
		r.cluster.Status.Created = false

		// delete event (create it again)
		if r.cluster.ObjectMeta.DeletionTimestamp != nil {
			return r.Provision(ctx)
		}
	}

	// update replicas
	tmp := int32(r.cluster.Spec.Replicas)
	deployment.Spec.Replicas = &tmp

	if err := r.Update(ctx, deployment); err != nil {
		r.logger.Error(err, fmt.Sprintf("failed to update ElkCluster deployment %s !", name))

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
