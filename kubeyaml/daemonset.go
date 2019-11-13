package kubeyaml

import (

	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	apps_v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdateDaemonSet(yamlPath string)error{
	daemonYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	daemon := apps_v1.DaemonSet{}
	daemonJson , err := yaml.ToJSON(daemonYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(daemonJson,&daemon); err != nil {
		return err
	}

	if _,err :=  k.Client.AppsV1().DaemonSets(daemon.Namespace).Get(daemon.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err) {
			if _,err := k.Client.AppsV1().DaemonSets(daemon.Namespace).Create(&daemon);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{
		if _,err :=  k.Client.AppsV1().DaemonSets(daemon.Namespace).Update(&daemon);err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)GetDaemonSet(namespace,name string)(*apps_v1.DaemonSet,error){
	if daemon,err := k.Client.AppsV1().DaemonSets(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return nil,err
	}else{
		return daemon,nil
	}
}

func (k KubeClient)DeleteDaemonSet(namespace,name string)error{
	if _,err :=  k.Client.AppsV1().DaemonSets(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		if err :=  k.Client.AppsV1().DaemonSets(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)DeleteDaemonSetByYaml(yamlPath string)error{
	daemonYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	daemon := apps_v1.DaemonSet{}
	daemonJson , err := yaml.ToJSON(daemonYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(daemonJson,&daemon); err != nil {
		return err
	}
	if _,err :=  k.Client.AppsV1().DaemonSets(daemon.Namespace).Get(daemon.Name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		if err :=  k.Client.AppsV1().DaemonSets(daemon.Namespace).Delete(daemon.Name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)ListDaemonSet(namespace string)([]apps_v1.DaemonSet,error){
	if list,err :=  k.Client.AppsV1().DaemonSets(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil,err
	}else{
		return list.Items,nil
	}
}