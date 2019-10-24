package kubeutils

import (
	"errors"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
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

type PodInfo struct {
	Namespace string
	PodName string
	Volumes []apiv1.Volume
	ImageName string
	Command []string
	Env 	[]apiv1.EnvVar
	VolumeMounts []apiv1.VolumeMount
}

func (k *KubeClient)CreatePod(pod PodInfo)error{
	pinter:=k.Client.CoreV1().Pods(pod.Namespace)
	podv1 := &apiv1.Pod{
		ObjectMeta:metav1.ObjectMeta{
			Name:	pod.PodName,
			Annotations: map[string]string{
				"name": pod.PodName,
			},
		},
		Spec:apiv1.PodSpec{
			Volumes: pod.Volumes,
			RestartPolicy:apiv1.RestartPolicyAlways,
			Containers:[]apiv1.Container{
				{
					Name: pod.PodName,
					Image:pod.ImageName,
					ImagePullPolicy:apiv1.PullIfNotPresent,
					Env:	pod.Env,
					VolumeMounts:pod.VolumeMounts,
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

func  (k *KubeClient) DeletePod(namespace,podname string)error{
	pinter:=k.Client.CoreV1().Pods(namespace)
	err := pinter.Delete(podname,&metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(fmt.Errorf("delete pod error : %v",err))
		return err
	}
	return nil
}

func (k *KubeClient) WatchPod(namespace,podname string)(int, error) {
	pinter:=k.Client.CoreV1().Pods(namespace)
	watchinter,err := pinter.Watch(metav1.ListOptions{})
	if err != nil {
		fmt.Println(fmt.Errorf("delete pod error : %v",err))
		return 0 ,err
	}
	select {
	case wr := <- watchinter.ResultChan():
		switch wr.Type {
		case watch.Added:
			fmt.Println(wr.Object)
			return 1,nil
		case watch.Error:
			fmt.Println(wr.Object)
			return 0,errors.New("create Pod err")
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