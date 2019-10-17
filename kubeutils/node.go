package kubeutils

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func (k *KubeClient)GetNode()[]v1.Node{
	nodeList,err := k.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _,node := range nodeList.Items{

		fmt.Printf("nodeName:%s ", node.Name )
	}
	return nodeList.Items
}
