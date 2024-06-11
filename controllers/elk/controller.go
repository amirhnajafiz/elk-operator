package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	monitoringamirhnajafizgithubcomv1beta1 "github.com/amirhnajafiz/elk-operator/api/v1beta1"
)

// ELKReconciler reconciles a ELK object
type ELKReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	logger    logr.Logger
	instance  *monitoringamirhnajafizgithubcomv1beta1.ELK
	namespace string
}

//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.amirhnajafiz.github.com,resources=elks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ELKReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// set logger
	r.logger = log.FromContext(ctx)
	r.namespace = req.Namespace

	// trying to get ELK instance
	err := r.Get(ctx, req.NamespacedName, r.instance)
	if err != nil {
		r.logger.Error(err, "")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ELKReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringamirhnajafizgithubcomv1beta1.ELK{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
