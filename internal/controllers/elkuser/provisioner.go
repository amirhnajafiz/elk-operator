package elkuser

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) Provision(ctx context.Context) (ctrl.Result, error) {
	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
