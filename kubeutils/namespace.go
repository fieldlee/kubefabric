package kubeutils

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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