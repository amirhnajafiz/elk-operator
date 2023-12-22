package elkcluster

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/amirhnajafiz/elk-operator/api/v1alpha1"
)

// Reconciler reconciles a ElkCluster object
type Reconciler struct {
	client.Client
	logger  logr.Logger
	cluster *v1alpha1.ElkCluster
	Scheme  *runtime.Scheme
}

//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.logger = log.FromContext(ctx)
	r.cluster = &v1alpha1.ElkCluster{}

	// get resource object
	switch err := r.Get(ctx, req.NamespacedName, r.cluster); {
	case apierrors.IsNotFound(err):
		// resource not found
		r.logger.Info(fmt.Sprintf("ElkCluster %s in namespace %s not found!", req.Name, req.Namespace))
		return ctrl.Result{}, nil
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default:
		// delete event
		if r.cluster.ObjectMeta.DeletionTimestamp != nil {
			return r.Cleanup(ctx)
		}
	}

	return r.Provision(ctx)
}
