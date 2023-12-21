package elkcluster

import (
	"github.com/amirhnajafiz/elk-operator/api/v1alpha1"

	v1 "k8s.io/api/apps/v1"
)

func (r *Reconciler) tearUp(cluster *v1alpha1.ElkCluster) error {
	return nil
}

func (r *Reconciler) getElasticsearchDeployment() *v1.Deployment {
	return nil
}

func (r *Reconciler) getKibanaDeployment() *v1.Deployment {
	return nil
}
