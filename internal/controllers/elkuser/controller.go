package elkuser

import (
	ctrl "sigs.k8s.io/controller-runtime"

	monitoringv1alpha1 "github.com/amirhnajafiz/elk-operator/api/v1alpha1"
)

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringv1alpha1.ElkUser{}).
		Complete(r)
}
