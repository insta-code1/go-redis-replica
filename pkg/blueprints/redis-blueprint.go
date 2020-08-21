package blueprints

import (
	"github.com/insta-code1/go-redis-replica/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedisReplicaDeployment redis deployment blueprint.
func (*Blueprints) RedisReplicaDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-slave",
			Labels: map[string]string{
				"app": "redis",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "redis",
					"role": "slave",
					"tier": "backend",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "redis",
						"role": "slave",
						"tier": "backend",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "slave",
							Image: "gcr.io/google_samples/gb-redisslave:v3",
							Ports: []apiv1.ContainerPort{
								{
									// Name:          "http",
									// Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 6379,
								},
							},
							Resources: apiv1.ResourceRequirements{
								// Requests: apiv1.ResourceList{
								// 	"cpu":    "100m",
								// 	"memory": "100Mi",
								// },
							},
							ImagePullPolicy: apiv1.PullIfNotPresent,
							Env: []apiv1.EnvVar{
								{
									Name:  "GET_HOSTS_FROM",
									Value: "dns",
								},
							},
						},
					},
				},
			},
		},
	}
}

// RedisReplicaService redis service blueprint.
func (*Blueprints) RedisReplicaService() *apiv1.Service {
	return &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-slave",
			Labels: map[string]string{
				"app":  "redis",
				"role": "slave",
				"tier": "backend",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 6379,
				},
			},
			Selector: map[string]string{
				"app":  "redis",
				"role": "slave",
				"tier": "backend",
			},
		},
	}
}
