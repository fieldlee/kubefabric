package main

import (
	"client-go/k8s"
	"fmt"
)

func main() {

	kubeClient := k8s.InitClient()

	podList := kubeClient.GetPod("")

	for _,pod := range podList{
		fmt.Printf("pod name :%s pod type:%s namespace:%s\n",pod.Name,pod.Kind,pod.Namespace)
	}

	nodeList := kubeClient.GetNode()

	fmt.Println()
	for _,node := range nodeList{
		fmt.Printf("node name :%s node type:%s \n",node.Name,node.Kind)
	}
	// Deployment
	deployment,err := kubeClient.CreateDeployment()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(deployment.Namespace,deployment.Name)

	// Update Deployment
	//err := kubeClient.UpdateDeployment("demo-deployment")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	// List Deployment

	listDeployment,err := kubeClient.ListDeployment()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _,deploy:=range listDeployment{
		fmt.Println(deploy.Namespace,deploy.Name)
	}
	// Delete Deployment

	//err = kubeClient.DeleteDeployment("demo-deployment")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}
