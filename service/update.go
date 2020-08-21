package service

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

// UpdateDeployment updates the specified deployment.
func (s *Service) UpdateDeployment() {
	deploymentsClient := s.Dao.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Updating deployment...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		// result.Spec.Replicas = int32Ptr(1)                             // reduce replica count
		containers := result.Spec.Template.Spec.Containers
		for idx := range containers {
			result.Spec.Template.Spec.Containers[idx].Image = "postgres:9.5"
		}

		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}
