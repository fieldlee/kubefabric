package kubeyaml

import (
	"encoding/json"
	"io/ioutil"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func (k KubeClient)NamespaceByYaml(yamlPath string)error{

	namespaceYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	namespace := core_v1.Namespace{}

	namespaceJson , err := yaml.ToJSON(namespaceYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(namespaceJson,&namespace); err != nil {
		return err
	}

	if _,err := k.Client.CoreV1().Namespaces().Get(namespace.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){
			if _, err := k.Client.CoreV1().Namespaces().Create(&namespace);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{ ////////////////////更新
		if _, err := k.Client.CoreV1().Namespaces().Update(&namespace);err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)CreateNamespace(namespace string)error{
	n := k.Client.CoreV1().Namespaces()
	nameConfig := &core_v1.Namespace{
		ObjectMeta:meta_v1.ObjectMeta{
			Name:namespace,
		},
	}
	_,err:=n.Create(nameConfig)
	if err != nil {
		return err
	}
	return nil
}

func (k KubeClient)ListNamespace(namespace string)(*core_v1.Namespace,error){
	if namespace,err := k.Client.CoreV1().Namespaces().Get(namespace,meta_v1.GetOptions{});err != nil {
		return nil,err
	}else{
		return namespace,nil
	}
}

func (k KubeClient)DeleteNamespace(namespace string)error{
	if err := k.Client.CoreV1().Namespaces().Delete(namespace,&meta_v1.DeleteOptions{});err != nil {
		return err
	}else{
		return nil
	}
}