package elkuser

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/amirhnajafiz/elk-operator/pkg/crypto"
	"github.com/amirhnajafiz/elk-operator/pkg/query"
)

func (r *Reconciler) Provision(ctx context.Context) (ctrl.Result, error) {
	// create a user query instance
	user := &query.UsersQuery{
		Username: r.user.Spec.Username,
		Password: crypto.Hash(r.user.Spec.Password),
		Roles:    r.user.Spec.Roles,
		Clusters: r.user.Spec.Clusters,
	}

	r.user.Status.Created = true
	r.user.Spec.Password = user.Password

	tx, _ := r.db.Begin()

	// insert into db
	if _, err := tx.Exec(user.QueryInsert()); err != nil {
		r.logger.Error(err, "failed to register user")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// update user resource
	if err := r.Update(ctx, r.user); err != nil {
		r.logger.Error(err, "failed to update user resource")

		_ = tx.Rollback()

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	_ = tx.Commit()

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
