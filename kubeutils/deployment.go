package kubeutils

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubefabric/utils"
)
func (k *KubeClient)CreateDeployment()(*appsv1.Deployment,error){

	deploymentsClient := k.Client.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
			Labels: map[string]string{
				"app":"demo",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
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
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		return nil,err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return result,nil
}

func (k *KubeClient)ListDeployment()([]appsv1.Deployment, error){
	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	deploymentsClient := k.Client.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
	return list.Items,nil
}

func (k *KubeClient)UpdateDeployment(deploymentName string)error{
	deploymentsClient := k.Client.AppsV1().Deployments(apiv1.NamespaceDefault)

	result, getErr := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
	if getErr != nil {
		return getErr
	}
	result.Spec.Replicas = utils.Int32Ptr(1)                           // reduce replica count
	result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
	_, updateErr := deploymentsClient.Update(result)

	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (k *KubeClient)DeleteDeployment(deploymentName string)error{
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := k.Client.AppsV1().Deployments(apiv1.NamespaceDefault)
	if err := deploymentsClient.Delete(deploymentName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	return nil
}