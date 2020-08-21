package blueprints

import (
	"github.com/insta-code1/go-redis-replica/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// RedisMasterDeployment redis deployment blueprint.
func (*Blueprints) RedisMasterDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-master",
			Labels: map[string]string{
				"app": "redis",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "redis",
					"role": "master",
					"tier": "backend",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "redis",
						"role": "master",
						"tier": "backend",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "master",
							Image: "redis:6.0.0",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 6379,
								},
							},
							ImagePullPolicy: apiv1.PullIfNotPresent,
						},
					},
				},
			},
		},
	}
}

// RedisMasterService redis service blueprint.
func (*Blueprints) RedisMasterService() *apiv1.Service {
	return &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-master",
			Labels: map[string]string{
				"app":  "redis",
				"role": "master",
				"tier": "backend",
			},
		},
		Spec: apiv1.ServiceSpec{
			Type:         apiv1.ServiceTypeClusterIP,
			ExternalName: "redis-master",
			Ports: []apiv1.ServicePort{
				{
					Port:       6379,
					TargetPort: intstr.FromInt(6379),
				},
			},
			Selector: map[string]string{
				"app":  "redis",
				"role": "master",
				"tier": "backend",
			},
		},
	}
}

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
							Image: "redis:6.0.0",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 6379,
								},
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
					DNSPolicy: apiv1.DNSClusterFirst,
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
