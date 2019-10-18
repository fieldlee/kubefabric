package kubeutils

import (
	"errors"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)
func (k *KubeClient)CreateNamespace(namespace string)error{
	n := k.Client.CoreV1().Namespaces()
	nameConfig := &apiv1.Namespace{
		ObjectMeta:metav1.ObjectMeta{
			Name:namespace,
		},
		Spec:apiv1.NamespaceSpec{

		},
	}
	_,err:=n.Create(nameConfig)
	if err != nil {
		return err
	}
	return nil
}

func (k *KubeClient)ListNamespace()([]apiv1.Namespace,error){
	n := k.Client.CoreV1().Namespaces()
	nameList,err := n.List(metav1.ListOptions{})
	if err != nil {
		return nil , err
	}
	return nameList.Items,nil
}

func (k *KubeClient)DelelteNamespace(namespace string)error{
	n := k.Client.CoreV1().Namespaces()
	err := n.Delete(namespace,&metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *KubeClient)WatchNamespace(namespace string)(int,error){
	n := k.Client.CoreV1().Namespaces()
	watchInter,err  := n.Watch(metav1.ListOptions{})
	if err != nil {
		return 0, err
	}
	select {
		case wr := <- watchInter.ResultChan():
			switch wr.Type {
			case watch.Added:
				fmt.Println(wr.Object)
				return 1,nil
			case watch.Error:
				fmt.Println(wr.Object)
				return 0,errors.New("create namespace err")
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