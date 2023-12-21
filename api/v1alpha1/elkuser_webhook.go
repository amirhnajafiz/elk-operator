/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ElkUser) ValidateUpdate(old runtime.Object) error {
	elkuserlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ElkUser) ValidateDelete() error {
	elkuserlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
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
