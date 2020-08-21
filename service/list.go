package service

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListDeployments list deployments
func (s *Service) ListDeployments() ([]string, error) {
	deploymentList := []string{}
	deploymentsClient := s.Dao.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return deploymentList, err
	}
	for _, d := range list.Items {
		msg := fmt.Sprintf("Deployment * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		fmt.Println(msg)
		deploymentList = append(deploymentList, msg)
	}
	return deploymentList, nil
}
