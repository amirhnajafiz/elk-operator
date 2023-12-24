package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var elkclusterlog = logf.Log.WithName("elkcluster-resource")

func (r *ElkCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-monitoring-snappcload-io-v1alpha1-elkcluster,mutating=true,failurePolicy=fail,sideEffects=None,groups=monitoring.snappcload.io,resources=elkclusters,verbs=create;update,versions=v1alpha1,name=melkcluster.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ElkCluster{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ElkCluster) Default() {
	elkclusterlog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-monitoring-snappcload-io-v1alpha1-elkcluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=monitoring.snappcload.io,resources=elkclusters,verbs=create;update,versions=v1alpha1,name=velkcluster.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ElkCluster{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ElkCluster) ValidateCreate() error {
	elkclusterlog.Info("validate create", "name", r.Name)

	if r.Spec.Replicas%2 != 1 {
		return ErrReplicaNumber
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ElkCluster) ValidateUpdate(old runtime.Object) error {
	elkclusterlog.Info("validate update", "name", r.Name)

	if r.Spec.Replicas%2 != 1 {
		return ErrReplicaNumber
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ElkCluster) ValidateDelete() error {
	elkclusterlog.Info("validate delete", "name", r.Name)

	return nil
}
