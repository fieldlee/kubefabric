package kubeyaml

import (
	"io/ioutil"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	apps_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdatePod(yamlPath string)error{
	podYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	pod := apps_v1.Pod{}

	podJson , err := yaml.ToJSON(podYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(podJson,&pod); err != nil {
		return err
	}

	////////////////查找Pod 是否存在
	if _, err := k.Client.CoreV1().Pods(pod.Namespace).Get(pod.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){  /// not found 创建
			if _,err := k.Client.CoreV1().Pods(pod.Namespace).Create(&pod);err != nil {
				return err
			}else{
				return nil
			}
		}else{  ////其他错误返回
			return err
		}
	}else{  ////////////////已经存在，则更新Pod
		if _,err := k.Client.CoreV1().Pods(pod.Namespace).Update(&pod);err != nil {
			return err
		}else{
			return nil
		}
	}
}


func (k KubeClient)GetPod(namespace,name string)(*apps_v1.Pod,error){
	if pod, err := k.Client.CoreV1().Pods(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return nil,err
	}else{
		return pod,nil
	}
}

func (k KubeClient)DeletePod(namespace,name string)error{
	if _, err := k.Client.CoreV1().Pods(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return err
	}else{  ////////////////已经存在，delete
		if err := k.Client.CoreV1().Pods(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)DeletePodByYaml(yamlPath string)error{
	podYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	pod := apps_v1.Pod{}
	podJson , err := yaml.ToJSON(podYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(podJson,&pod); err != nil {
		return err
	}

	if _, err := k.Client.CoreV1().Pods(pod.Namespace).Get(pod.Name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		if err := k.Client.CoreV1().Pods(pod.Namespace).Delete(pod.Name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)ListPods(namespace string)([]apps_v1.Pod,error){
	if list,err := k.Client.CoreV1().Pods(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil,err
	}else{
		return list.Items,nil
	}
}