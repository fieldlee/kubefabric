package kubeutils

import (
	"errors"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/watch"
)

type ServiceInfo struct {
	Namespace string
	ServiceName string
	Selector map[string]string
	Port int
	Nodeport int
	BalanceIp string
}

func (k *KubeClient)CreateServiceByPort(service ServiceInfo)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(service.Namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:service.ServiceName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeNodePort,
			Selector:service.Selector,
			Ports:[]apiv1.ServicePort{
				{
					Name:       service.ServiceName,
					Port:       int32(service.Port),
					TargetPort: intstr.FromInt(service.Port),
					Protocol:   apiv1.ProtocolTCP,
					NodePort:   int32(service.Nodeport),
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

func (k *KubeClient)CreateServiceByIP(service ServiceInfo)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(service.Namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:service.ServiceName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeClusterIP,

			Selector:service.Selector,
			Ports:[]apiv1.ServicePort{
				{
					Name:       service.ServiceName,
					Port:       int32(service.Port),
					TargetPort: intstr.FromInt(service.Port),
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

func (k *KubeClient)CreateServiceLoadBalancer(service ServiceInfo)(*apiv1.Service,error){
	sinter := k.Client.CoreV1().Services(service.Namespace)
	svccon := &apiv1.Service{
		ObjectMeta:	metav1.ObjectMeta{
			Name:service.ServiceName,
		},
		Spec:apiv1.ServiceSpec{
			Type:apiv1.ServiceTypeClusterIP,
			Selector:service.Selector,
			LoadBalancerIP:service.BalanceIp,
			Ports:[]apiv1.ServicePort{
				{
					Name:       service.ServiceName,
					Port:       int32(service.Port),
					TargetPort: intstr.FromInt(service.Port),
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
func (k *KubeClient)WatchService(namespace string)(int,error){
	wservice := k.Client.CoreV1().Services(namespace)
	winter,err := wservice.Watch(metav1.ListOptions{})
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
			return 0,errors.New("create service err")
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

func (k *KubeClient)DeleteService(namespace,svcname string)error{
	s:= k.Client.CoreV1().Services(namespace)
	err := s.Delete(svcname,&metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}