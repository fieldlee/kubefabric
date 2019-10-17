package kubeutils

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func (k *KubeClient)GetPod(namespace string)[]v1.Pod{
	pods, err := k.Client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _,item := range pods.Items{
		fmt.Printf("itemName:%s  itemType:%s \n", item.Name,item.Kind)
	}
	return pods.Items
}


////声明pod对象
//var pod *v1.Pod
////创建pod
//pod, err := clientset.CoreV1().Pods(<namespace>).Create(<pod>)
////更新pod
//pod, err := clientset.CoreV1().Pods(<namespace>).Update(<pod>)
////删除pod
//err := clientset.CoreV1().Pods(<namespace>).Delete(<pod.Name>, &meta_v1.DeleteOptions{})
////查询pod
//pod, err := clientset.CoreV1().Pods(<namespace>).Get(<pod.Name>, &meta_v1.GetOptions{})
////列出pod
//podList, err := clientset.CoreV1().Pods(<namespace>).List(&meta_v1.ListOptions{})
////watch pod
//watchInterface, err := clientset.CoreV1().Pods(<namespace>).Watch(&meta_v1.ListOptions{})

func (k *KubeClient)CreatePod(namespace,podname string)error{
	pinter:=k.Client.CoreV1().Pods(namespace)
	podv1 := &apiv1.Pod{
		ObjectMeta:metav1.ObjectMeta{
			Name:podname,
			Labels: map[string]string{

			},
			Annotations: map[string]string{
				"name":podname,
			},
		},
		Spec:apiv1.PodSpec{
			RestartPolicy:apiv1.RestartPolicyAlways,
			NodeSelector: map[string]string{

			},
			Containers:[]apiv1.Container{
				{
					Name:podname,
					ImagePullPolicy:apiv1.PullIfNotPresent,
				},
			},
		},
	}


	_,err := pinter.Create(podv1)
	if err != nil {
		return err
	}
	return nil
}