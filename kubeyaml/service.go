package kubeyaml

import (
	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)ServiceByYaml(yamlPath string)error{
	serviceYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	service := core_v1.Service{}

	serviceJson , err := yaml.ToJSON(serviceYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(serviceJson,&service); err != nil {
		return err
	}
	////// 查询

	if _,err := k.Client.CoreV1().Services(service.Namespace).Get(service.Name,meta_v1.GetOptions{});err !=nil{
		if errors.IsNotFound(err){ ////// 创建
			if _,err := k.Client.CoreV1().Services(service.Namespace).Create(&service);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{ ///// 更新
		if _,err := k.Client.CoreV1().Services(service.Namespace).Update(&service);err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)DeleteServiceByYaml(yamlPath string)error{
	serviceYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	service := core_v1.Service{}

	serviceJson , err := yaml.ToJSON(serviceYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(serviceJson,&service); err != nil {
		return err
	}
	
	if err := k.Client.CoreV1().Services(service.Namespace).Delete(service.Name,&meta_v1.DeleteOptions{});err != nil {
		return err
	}
	return nil
}


func (k KubeClient)DeleteService(namespace,name string)error{
	if err := k.Client.CoreV1().Services(namespace).Delete(name,&meta_v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}



func (k KubeClient)GetService( namespace,name string)(*core_v1.Service,error){
	if srv,err := k.Client.CoreV1().Services(namespace).Get(name,meta_v1.GetOptions{}); err != nil {
		return nil,err
	}else{
		return srv,nil
	}
}

func (k KubeClient)ListServices(namespace string)([]core_v1.Service,error){
	if list,err := k.Client.CoreV1().Services(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil,err
	}else{
		return list.Items,nil
	}
}