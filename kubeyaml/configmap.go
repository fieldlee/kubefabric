package kubeyaml

import (
	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)ConfigMapByYaml(yamlPath string) error  {
	configYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	config := core_v1.ConfigMap{}
	configJson , err := yaml.ToJSON(configYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(configJson,&config); err != nil {
		return err
	}
	/////////////////////Get config map
	if _,err := k.Client.CoreV1().ConfigMaps(config.Namespace).Get(config.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){
			///////////////////创建config map
			if _,err := k.Client.CoreV1().ConfigMaps(config.Namespace).Create(&config);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{
		if _,err :=k.Client.CoreV1().ConfigMaps(config.Namespace).Update(&config); err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)GetConfigMap(namespace,name string)(*core_v1.ConfigMap,error){
	if configmap, err := k.Client.CoreV1().ConfigMaps(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return nil,err
	}else{
		return configmap,nil
	}
}

func (k KubeClient)DeleteConfigMapByYaml(yamlPath string)error{
	configYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	config := core_v1.ConfigMap{}
	configJson , err := yaml.ToJSON(configYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(configJson,&config); err != nil {
		return err
	}
	if err := k.Client.CoreV1().ConfigMaps(config.Namespace).Delete(config.Name,&meta_v1.DeleteOptions{});err!=nil{
		return err
	}
	return nil
}

func (k KubeClient)DeleteConfigmap(namespace,name string)error{
	if err := k.Client.CoreV1().ConfigMaps(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
		return err
	}
	return nil
}

func (k KubeClient)ListConfigmap(namespace string)([]core_v1.ConfigMap,error){
	if list,err := k.Client.CoreV1().ConfigMaps(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil,err
	}else{
		return list.Items,nil
	}
}