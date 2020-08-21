package blueprints

import (
	"github.com/insta-code1/go-redis-replica/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deploymentSchema(objectMeta metav1.ObjectMeta, selector *metav1.LabelSelector) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: objectMeta,
		// ObjectMeta: metav1.ObjectMeta{
		// 	Name: "demo-deployment",
		// },
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(5),
			// Selector: selector,
			// Selector: &metav1.LabelSelector{
			// 	MatchLabels: map[string]string{
			// 		"app": "demo",
			// 	},
			// },
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "postgres:9.0",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 5432,
								},
							},
						},
					},
				},
			},
		},
	}
}
