package elkcluster

import (
	"context"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) Provision(ctx context.Context) (ctrl.Result, error) {
	if r.cluster.Status.Created {
		return r.Modify(ctx)
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}

func (r *Reconciler) getElasticsearchDeployment() *v1.Deployment {
	replicas := int32(r.cluster.Spec.Replicas)

	return &v1.Deployment{
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Template: core.PodTemplateSpec{},
		},
	}
}

func (r *Reconciler) getKibanaDeployment() *v1.Deployment {
	return nil
}
