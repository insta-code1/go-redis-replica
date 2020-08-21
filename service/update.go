package service

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

// UpdateDeployment updates the specified deployment.
func (s *Service) UpdateDeployment(deploymentName, image string) (string, error) {
	deploymentsClient := s.Dao.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Updating deployment...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

		result, getErr := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		containers := result.Spec.Template.Spec.Containers
		for idx := range containers {
			result.Spec.Template.Spec.Containers[idx].Image = image
		}

		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return "", fmt.Errorf("Update failed: %v", retryErr)
	}
	fmt.Println("Updated deployment...")
	return "Updated deployment...", nil
}
