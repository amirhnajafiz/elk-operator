package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	monitoringamirhnajafizgithubcomv1beta1 "github.com/amirhnajafiz/elk-operator/api/v1beta1"
)

// Reconciler reconciles a ELK object
type Reconciler struct {
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
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// set logger, instance, and variables
	r.logger = log.FromContext(ctx)
	r.instance = &monitoringamirhnajafizgithubcomv1beta1.ELK{}
	r.initVars(req)

	// trying to get ELK instance
	switch err := r.Get(ctx, req.NamespacedName, r.instance); {
	case apierrors.IsNotFound(err): // if not found, cleanup the resources
		return r.Cleanup(ctx)
	case err != nil:
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default: // check deletion timestamp to cleanup the resources
		if r.instance.ObjectMeta.DeletionTimestamp != nil {
			return r.Cleanup(ctx)
		}
	}

	return r.Handler(ctx)
}

func (r *Reconciler) initVars(req ctrl.Request) {
	r.namespace = req.Namespace
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringamirhnajafizgithubcomv1beta1.ELK{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
