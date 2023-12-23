package elkcluster

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Provision is the handler for create and update requests
func (r *Reconciler) Provision(ctx context.Context) (ctrl.Result, error) {
	if r.cluster.Status.Created {
		return r.Modify(ctx)
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
