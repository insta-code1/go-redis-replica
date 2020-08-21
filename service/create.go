package service

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment creates a deployment
func (s *Service) CreateDeployment(deployment *appsv1.Deployment) {
	fmt.Println("Creating deployment...")
	deploymentsClient := s.Dao.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

// CreateService creats a service
func (s *Service) CreateService(service *apiv1.Service) {
	// deploymentsClient := s.Dao.ClientSet.AppsV1().Ser(apiv1.NamespaceDefault)
	serviceClient := s.Dao.ClientSet.CoreV1().Services(apiv1.NamespaceDefault)
	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}
