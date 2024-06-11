package controllers

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Cleanup method is used to delete all resources that are created by ELK operator.
func (r *ELKReconciler) Cleanup(ctx context.Context) (ctrl.Result, error) {
	return subreconciler.Evaluate(subreconciler.Requeue())
}
