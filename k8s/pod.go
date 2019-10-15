package k8s

import (

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func (k *KubeClient)GetPod(namespace string)[]v1.Pod{
	pods, err := k.Client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//for _,item := range pods.Items{
	//	fmt.Printf("itemName:%s  itemType:%s \n", item.Name,item.Kind)
	//}
	return pods.Items
}
