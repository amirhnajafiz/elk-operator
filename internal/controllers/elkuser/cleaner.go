package elkuser

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) Cleanup(ctx context.Context) (ctrl.Result, error) {
	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
