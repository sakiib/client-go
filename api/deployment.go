package api

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment() {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Println(err)
		return
	}

	deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apiserver",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "apiserver",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "apiserver",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "apiserver",
							Image: "sakibalamin/apiserver:1.0.1",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Deployment created: %q.\n", result.GetObjectMeta().GetName())
}

func GetDeployment() {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Println(err)
		return
	}

	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range list.Items {
		fmt.Printf("%s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}