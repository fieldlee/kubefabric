package kubeutils

import (

	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

////声明ingress对象
//var ingress *v1beta1.Ingress
////构造ingress对象
////创建ingress
//ingress, err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).Create(<ingress>)
////更新ingress
//ingress, err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).Update(<ingress>)
////删除ingress
//err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).Delete(<ingress.Name>, &meta_v1.DeleteOptions{})
////查询ingress
//ingress, err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).Get(<ingress.Name>, &meta_v1.GetOptions{})
////列出ingress
//ingressList, err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).List(&meta_v1.ListOptions{})
////watch ingress
//watchInterface, err := clientset.ExtensionsV1beta1().Ingresses(<namespace>).Watch(&meta_v1.ListOptions{})


func (k *KubeClient)CreateIngress(namespace,ingname string)(*v1beta1.Ingress,error){
	svcIng:=k.Client.ExtensionsV1beta1().Ingresses(namespace)
	svcCon := &v1beta1.Ingress{
		ObjectMeta:metav1.ObjectMeta{
			Name:ingname,
		},
		Spec:v1beta1.IngressSpec{
			Backend:&v1beta1.IngressBackend{
				ServiceName:"",
				ServicePort:intstr.FromInt(100),
			},
			TLS:[]v1beta1.IngressTLS{
				{

				},
			},
			Rules:[]v1beta1.IngressRule{
				{
					Host:"",
					IngressRuleValue:v1beta1.IngressRuleValue{
						HTTP:&v1beta1.HTTPIngressRuleValue{
							Paths:[]v1beta1.HTTPIngressPath{
								{
									Path:"",
									Backend:v1beta1.IngressBackend{
										ServiceName:"",
										ServicePort:intstr.FromInt(100),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	svcAccount,err := svcIng.Create(svcCon)

	if err != nil {
		return nil,err
	}
	return svcAccount,nil
}

func (k *KubeClient)DeleteIngress(namespace,name string)error{
	err := k.Client.ExtensionsV1beta1().Ingresses(namespace).Delete(name,&metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}


func (k *KubeClient)GetIngress(namespace,name string)(*v1beta1.Ingress,error){
	v1Ingress,err := k.Client.ExtensionsV1beta1().Ingresses(namespace).Get(name,metav1.GetOptions{})
	if err != nil {
		return v1Ingress,err
	}
	return v1Ingress,nil
}

func (k *KubeClient)ListIngress(namespace,name string)([]v1beta1.Ingress,error){
	ingress,err := k.Client.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	return ingress.Items,nil
}