package elkuser

import (
	"context"
	"database/sql"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

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

//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	// TODO: get user
	// TODO: if deleted => remove from database
	// TODO: if existed => prevent
	// TODO: if not create a new user and insert it into db

	return ctrl.Result{}, nil
}
