package service

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeleteDeployment deletes the specified deployment.
func (s *Service) DeleteDeployment(deploymentName string) (string, error) {
	fmt.Println("Deleting deployment...")
	deploymentsClient := s.Dao.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return "", fmt.Errorf("Delete failed: %v", err)
	}
	return fmt.Sprint("Deleted deployment."), nil
}
