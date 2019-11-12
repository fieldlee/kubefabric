package kubeyaml

import (

	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	apps_v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)DeploymentByYaml(yamlPath string)error{
	deployYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	deploy := apps_v1.Deployment{}

	deployJson , err := yaml.ToJSON(deployYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(deployJson,&deploy); err != nil {
		return err
	}

	////////////////查找Deployment 是否存在
	if _, err := k.Client.AppsV1().Deployments(deploy.Namespace).Get(deploy.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){  /// not found 创建
			if _,err := k.Client.AppsV1().Deployments(deploy.Namespace).Create(&deploy);err != nil {
				return err
			}else{
				return nil
			}
		}else{  ////其他错误返回
			return err
		}
	}else{  ////////////////已经存在，则更新Deployment
		if _,err := k.Client.AppsV1().Deployments(deploy.Namespace).Update(&deploy);err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)DeleteDeploymentByYaml(yamlPath string)error{
	deployYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	deploy := apps_v1.Deployment{}

	deployJson , err := yaml.ToJSON(deployYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(deployJson,&deploy); err != nil {
		return err
	}

	if err := k.Client.AppsV1().Deployments(deploy.Namespace).Delete(deploy.Name,&meta_v1.DeleteOptions{});err != nil {
		return err
	}
	return nil
}

func (k KubeClient)DeleteDeployment(namespace,name string)error{

	if err := k.Client.AppsV1().Deployments(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
		return err
	}
	return nil
}

func (k KubeClient)ListDeployment(namespace,name string) (*apps_v1.DeploymentList, error ){
	if list,err := k.Client.AppsV1().Deployments(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil , err
	}else{
		return list, nil
	}
}