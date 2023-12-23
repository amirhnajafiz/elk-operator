package elkcluster

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Provision is the handler for create and update requests
func (r *Reconciler) Provision(ctx context.Context) (ctrl.Result, error) {
	if r.cluster.Status.Created { // update
		return r.Modify(ctx)
	}

	// create deployment
	deployment := r.getElkDeployment()

	if err := controllerutil.SetControllerReference(r.cluster, deployment, r.scheme); err != nil {
		r.logger.Error(err, "failed to own the deployment")

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	if err := r.Create(ctx, deployment); err != nil {
		r.logger.Error(err, "failed to create a deployment")

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// update status
	r.cluster.Status.Created = true

	if err := r.Update(ctx, r.cluster); err != nil {
		r.logger.Error(err, "failed to update the resource")

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}

// get the deployment file of elk cluster
func (r *Reconciler) getElkDeployment() *v1.Deployment {
	// cluster replicas
	replicas := int32(r.cluster.Spec.Replicas)

	// cluster ports
	ports := []core.ContainerPort{
		{
			HostPort:      9200,
			ContainerPort: 9200,
		},
		{
			HostPort:      5044,
			ContainerPort: 5044,
		},
		{
			HostPort:      9300,
			ContainerPort: 9300,
		},
		{
			HostPort:      9600,
			ContainerPort: 9600,
		},
	}

	// Kibana dashboard expose
	if r.cluster.Spec.Dashboard {
		ports = append(ports, core.ContainerPort{
			HostPort:      5601,
			ContainerPort: 5601,
		})
	}

	return &v1.Deployment{
		ObjectMeta: ctrl.ObjectMeta{
			Name:        fmt.Sprintf("%s-deployment", r.cluster.Name),
			Namespace:   r.cluster.Namespace,
			Labels:      r.cluster.Labels,
			Annotations: r.cluster.Annotations,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Template: core.PodTemplateSpec{
				Spec: core.PodSpec{
					Containers: []core.Container{
						{
							Image: "sebp/elk:latest",
							Ports: ports,
						},
					},
				},
			},
		},
	}
}
