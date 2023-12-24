package v1alpha1

import (
	"database/sql"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var elkuserlog = logf.Log.WithName("elkuser-resource")

var (
	dbConnection *sql.DB
)

func (r *ElkUser) SetupWebhookWithManager(mgr ctrl.Manager, db *sql.DB) error {
	dbConnection = db

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-monitoring-snappcload-io-v1alpha1-elkuser,mutating=true,failurePolicy=fail,sideEffects=None,groups=monitoring.snappcload.io,resources=elkusers,verbs=create;update,versions=v1alpha1,name=melkuser.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ElkUser{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ElkUser) Default() {
	elkuserlog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-monitoring-snappcload-io-v1alpha1-elkuser,mutating=false,failurePolicy=fail,sideEffects=None,groups=monitoring.snappcload.io,resources=elkusers,verbs=create;update,versions=v1alpha1,name=velkuser.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ElkUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ElkUser) ValidateCreate() error {
	elkuserlog.Info("validate create", "name", r.Name)

	if flag, err := r.userExists(); err != nil {
		return err
	} else if flag {
		return ErrUsernameExists
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ElkUser) ValidateUpdate(old runtime.Object) error {
	elkuserlog.Info("validate update", "name", r.Name)

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ElkUser) ValidateDelete() error {
	elkuserlog.Info("validate delete", "name", r.Name)

	if flag, err := r.userExists(); err != nil {
		return err
	} else if !flag {
		return ErrUserNotFound
	}

	return nil
}

func (r *ElkUser) userExists() (bool, error) {
	var (
		username = r.Spec.Username
		count    int
	)

	if row, err := dbConnection.Query("select count(*) from users where username = ?", username); err == nil {
		if er := row.Scan(&count); er != nil {
			return false, fmt.Errorf("failed to parse query: %w", er)
		}

		return count > 0, nil
	}

	return false, fmt.Errorf("failed to execute query")
}
