package elkuser

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/amirhnajafiz/elk-operator/pkg/query"
)

func (r *Reconciler) Cleanup(ctx context.Context) (ctrl.Result, error) {
	user := &query.UsersQuery{
		Username: r.user.Spec.Username,
	}

	// delete user from db
	if _, err := r.db.ExecContext(ctx, user.QueryDelete()); err != nil {
		r.logger.Error(err, "failed to delete user")

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
