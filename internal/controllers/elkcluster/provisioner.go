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

	if r.cluster.Spec.Dashboard {
		ports = append(ports, core.ContainerPort{
			HostPort:      5601,
			ContainerPort: 5601,
		})
	}

	return &v1.Deployment{
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
