package controllers

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Handler is the main method of our reconciler, it calls other methods
// based on the current instance status.
func (r *Reconciler) Handler(ctx context.Context) (ctrl.Result, error) {
	return subreconciler.Evaluate(subreconciler.Requeue())
}
