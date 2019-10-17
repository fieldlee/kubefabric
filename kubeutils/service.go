package kubeutils

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)
////声明service对象
//var service *v1.Service
////构造service对象
////创建service
//service, err := clientset.CoreV1().Services(<namespace>).Create(<service>)
////更新service
//service, err := clientset.CoreV1().Services(<namespace>).Update(<service>)
////删除service
//err := clientset.CoreV1().Services(<namespace>).Delete(<service.Name>, &meta_v1.DeleteOptions{})
////查询service
//service, err := clientset.CoreV1().Services(<namespace>).Get(<service.Name>, &meta_v1.GetOptions{})
////列出service
//serviceList, err := clientset.CoreV1().Services(<namespace>).List(&meta_v1.ListOptions{})
////watch service
//watchInterface, err := clientset.CoreV1().Services(<namespace>).Watch(&meta_v1.ListOptions{})

//Spec: v1.ServiceSpec{
//Type:     v1.ServiceTypeNodePort,
//Selector: pod.Labels,
//Ports: []v1.ServicePort{
//{
//Port: 8888,
//},
//},
func (k *KubeClient)CreateServiceByPort(namespace,svcName string,selector map[string]string,port,nodeport int)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:svcName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeNodePort,
			Selector:selector,
			Ports:[]apiv1.ServicePort{
				{
					Name:       svcName,
					Port:       int32(port),
					TargetPort: intstr.FromInt(port),
					Protocol:   "TCP",
					NodePort:   int32(nodeport),
				},
			},
		},
	}

	svc,err := sinter.Create(svccon)
	if err != nil {
		return nil,err
	}
	return svc, nil
}

func (k *KubeClient)CreateServiceByIP(namespace,svcName string,selector map[string]string,port int)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:svcName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeClusterIP,
			Selector:selector,
			Ports:[]apiv1.ServicePort{
				{
					Name:       svcName,
					Port:       int32(port),
					TargetPort: intstr.FromInt(port),
					Protocol:   "TCP",
					//NodePort:   int32(nodeport),
				},
			},
		},
	}

	svc,err := sinter.Create(svccon)
	if err != nil {
		return nil,err
	}
	return svc, nil
}

func (k *KubeClient)CreateServiceLoadBalancer(namespace,svcName string,selector map[string]string,port int)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:svcName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeClusterIP,
			Selector:selector,
			LoadBalancerIP:"",
			Ports:[]apiv1.ServicePort{
				{
					Name:       svcName,
					Port:       int32(port),
					TargetPort: intstr.FromInt(port),
					Protocol:   "TCP",
					//NodePort:   int32(nodeport),
				},
			},
		},
		Status:apiv1.ServiceStatus{
			LoadBalancer:apiv1.LoadBalancerStatus{
				Ingress:[]apiv1.LoadBalancerIngress{
					{
						IP:"",
						Hostname:"",
					},
				},
			},
		},
	}

	svc,err := sinter.Create(svccon)
	if err != nil {
		return nil,err
	}
	return svc, nil
}

func (k *KubeClient)GetService(namespace string)([]apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(namespace)
	svcList,err := sinter.List(metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	for _, svc := range svcList.Items{
		fmt.Printf("service name :%s namespace :%s\n",svc.Name,svc.Namespace)
	}
	return svcList.Items,nil
}

func (k *KubeClient)UpdateService(namespace , svcname string)error{
	s:= k.Client.CoreV1().Services(namespace)
	svcconfig,err := s.Get(svcname,metav1.GetOptions{})
	if err != nil {
		return err
	}

	_,err =s.Update(svcconfig)
	if err != nil {
		return err
	}
	return nil
}

func (k *KubeClient)ListService(namespace,svcname string)([]apiv1.Service,error){
	s:= k.Client.CoreV1().Services(namespace)
	svcLIst,err :=  s.List(metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	for _,svc := range svcLIst.Items{
		fmt.Printf("service name :%s namespace :%s\n",svc.Name,svc.Namespace)
	}
	return svcLIst.Items,nil
}