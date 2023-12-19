package elkuser

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Reconciler reconciles a ElkUser object
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=monitoring.snappcload.io,resources=elkusers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// Modify the Reconcile function to compare the state specified by
// the ElkUser object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	// TODO: get user
	// TODO: if deleted => remove from database
	// TODO: if existed => prevent
	// TODO: if not create a new user and insert it into db

	return ctrl.Result{}, nil
}
