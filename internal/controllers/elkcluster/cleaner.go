package elkcluster

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/amirhnajafiz/elk-operator/api/v1alpha1"
)

func (r *Reconciler) Cleanup(ctx context.Context) (ctrl.Result, error) {
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
	}

	// get users to remove them
	var (
		users v1alpha1.ElkUserList
		opts  = []client.ListOption{
			client.MatchingLabels{
				"cluster": r.cluster.Name,
			},
		}
	)

	if err := r.List(ctx, &users, opts...); err != nil {
		r.logger.Error(err, "failed to list users")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	for _, item := range users.Items {
		if err := r.Delete(ctx, &item); err != nil {
			r.logger.Error(err, "failed to remove user")
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
	}

	// delete deployment object
	if err := r.Delete(ctx, deployment); err != nil {
		r.logger.Error(err, "failed to delete deployment")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
