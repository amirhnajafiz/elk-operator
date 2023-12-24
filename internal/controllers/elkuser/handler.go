package elkuser

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/amirhnajafiz/elk-operator/api/v1alpha1"
)

// Reconciler reconciles a ElkUser object
type Reconciler struct {
	client.Client
	logger logr.Logger
	user   *v1alpha1.ElkUser
	scheme *runtime.Scheme
	db     *sql.DB
}

func NewReconciler(mgr manager.Manager, db *sql.DB) *Reconciler {
	return &Reconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		db:     db,
	}
}

//+kubebuilder:rbac:groups=monitoring.snappcloud.io,resources=elkusers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.snappcloud.io,resources=elkusers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.snappcloud.io,resources=elkusers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.logger = log.FromContext(ctx)
	r.user = &v1alpha1.ElkUser{}

	// get resource object
	switch err := r.Get(ctx, req.NamespacedName, r.user); {
	case apierrors.IsNotFound(err):
		// resource not found
		r.logger.Info(fmt.Sprintf("ElkUser %s in namespace %s not found!", req.Name, req.Namespace))
		return subreconciler.Evaluate(subreconciler.DoNotRequeue())
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default:
		// delete event
		if r.user.ObjectMeta.DeletionTimestamp != nil {
			return r.Cleanup(ctx)
		}
	}

	return r.Provision(ctx)
}
