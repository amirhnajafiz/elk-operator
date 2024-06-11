package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	monitoringamirhnajafizgithubcomv1beta1 "github.com/amirhnajafiz/elk-operator/api/v1beta1"
)

// ELKReconciler reconciles a ELK object
type ELKReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ELKReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ELKReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringamirhnajafizgithubcomv1beta1.ELK{}).
		Complete(r)
}
