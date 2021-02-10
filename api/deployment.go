package api

import (
	"context"
	"fmt"
	"log"

	"k8s.io/client-go/util/retry"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(replica int32, image string, name string) {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Println(err)
		return
	}

	deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(replica),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: "sakibalamin/" + image,
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

func UpdateDeployment(replica int32, image string, name string) {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Println(err.Error())
		return
	}

	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentClient.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(replica)
		result.Spec.Template.Spec.Containers[0].Image = "sakibalamin/" + image

		_, updateErr := deploymentClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		log.Println(retryErr.Error())
	}

	fmt.Println("Updated deployment...")
}

func DeleteDeployment(name string) {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Println(err.Error())
		return
	}

	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Deleted deployment.")
}
