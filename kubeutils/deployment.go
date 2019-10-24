package kubeutils

import (
	"errors"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"kubefabric/utils"
)

type DeploymentInfo struct {
	Namespace string
	DeploymentName string
	ImageName string
	VolumnName string
	VolumnPath string
	PVCName string
	ReplicaNum int
	Port int
	Command []string
}

func (k *KubeClient)CreateDeployment(deployment DeploymentInfo,envVarList []apiv1.EnvVar)(*appsv1.Deployment,error){

	//hostPathType := apiv1.HostPathFile

	deploymentsClient := k.Client.AppsV1().Deployments(deployment.Namespace)

	deploymentObj := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deployment.DeploymentName,
			Labels: map[string]string{
				"name":deployment.DeploymentName,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(int32(deployment.ReplicaNum)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"name": deployment.DeploymentName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": deployment.DeploymentName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  	deployment.DeploymentName,
							Image: 	deployment.ImageName,
							Command:deployment.Command,
							Env:	envVarList,
							Ports:  []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: int32(deployment.Port),
								},
							},
							VolumeMounts:[]apiv1.VolumeMount{
									{
										Name:deployment.VolumnName,
										MountPath:deployment.VolumnPath,
									},
							},
						},
					},
					Volumes:[]apiv1.Volume{
						{
							Name:deployment.VolumnName,
							VolumeSource:apiv1.VolumeSource{
								//HostPath:&apiv1.HostPathVolumeSource{
								//	Path:"",
								//	Type: &hostPathType,
								//},
								PersistentVolumeClaim:&apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName:deployment.PVCName,
								},
							},

						},
					},
				},

			},
		},
	}

	result, err := deploymentsClient.Create(deploymentObj)
	if err != nil {
		return nil,err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return result,nil
}

func (k *KubeClient)WatchDeploy(namespace string)(int,error){
	wdeploy := k.Client.AppsV1().Deployments(namespace)
	winter,err := wdeploy.Watch(metav1.ListOptions{})
	if err != nil {
		return 0,err
	}
	select {
	case wr := <- winter.ResultChan():
		switch wr.Type {
		case watch.Added:
			fmt.Println(wr.Object)
			return 1,nil
		case watch.Error:
			fmt.Println(wr.Object)
			return 0,errors.New("create deployment err")
		case watch.Deleted:
			fmt.Println(wr.Object)
			return -1,nil
		case watch.Modified:
			fmt.Println(wr.Object)
			return 1,nil
		}
	}
	return 0,nil
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

func (k *KubeClient)DeleteDeployment(namespace,deploymentName string)error{
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := k.Client.AppsV1().Deployments(namespace)
	if err := deploymentsClient.Delete(deploymentName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	return nil
}