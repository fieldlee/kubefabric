package kubeyaml

import (
	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	apps_v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)StatefulByYaml(yamlPath string)error{
	stateYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	state := apps_v1.StatefulSet{}

	stateJson , err := yaml.ToJSON(stateYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(stateJson,&state); err != nil {
		return err
	}
	////// 查询
	if _,err := k.Client.AppsV1().StatefulSets(state.Namespace).Get(state.Name,meta_v1.GetOptions{});err !=nil{
		if errors.IsNotFound(err){ ////// 创建
			if _,err := k.Client.AppsV1().StatefulSets(state.Namespace).Create(&state);err != nil {
				return err
			}else{
				return nil
			}

		}
		return err
	}else{ ///// 更新
		if _,err := k.Client.AppsV1().StatefulSets(state.Namespace).Update(&state);err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)DelStatefulByYaml(yamlPath string)error{
	stateYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	state := apps_v1.StatefulSet{}

	stateJson , err := yaml.ToJSON(stateYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(stateJson,&state); err != nil {
		return err
	}
	////// 查询
	if _,err := k.Client.AppsV1().StatefulSets(state.Namespace).Get(state.Name,meta_v1.GetOptions{});err !=nil{
		return err
	}else{ ///// 删除
		if err := k.Client.AppsV1().StatefulSets(state.Namespace).Delete(state.Name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)DelStateful(namespace,name string)error{
	if err := k.Client.AppsV1().StatefulSets(namespace).Delete(name,&meta_v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}


func (k KubeClient)GetStateful(namespace,name string)(*apps_v1.StatefulSet,error){
	if state,err := k.Client.AppsV1().StatefulSets(namespace).Get(name,meta_v1.GetOptions{}); err != nil {
		return nil,err
	}else{
		return state, nil
	}
}

func (k KubeClient)ListStateful(namespace string)([]apps_v1.StatefulSet,error){
	if state,err := k.Client.AppsV1().StatefulSets(namespace).List(meta_v1.ListOptions{}); err != nil {
		return nil,err
	}else{
		return state.Items, nil
	}
}