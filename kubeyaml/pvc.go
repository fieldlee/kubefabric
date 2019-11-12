package kubeyaml

import (
	"encoding/json"
	"io/ioutil"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	)

func (k KubeClient)CreatePvByYaml(yamlPath string)error{
	pvYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	pv := core_v1.PersistentVolume{}

	pvJson ,err := yaml.ToJSON(pvYaml)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(pvJson,&pv);err != nil {
		return err
	}
	//////////////////////查询
	if _, err := k.Client.CoreV1().PersistentVolumes().Get(pv.Name,meta_v1.GetOptions{}); err != nil {
		if errors.IsNotFound(err){
			if _, err := k.Client.CoreV1().PersistentVolumes().Create(&pv);err != nil {
				return err
			}
		}
		return err
	}else{  /////////////// 如果存在，更新
		if _,err := k.Client.CoreV1().PersistentVolumes().Update(&pv); err != nil {
			return err
		}
		return nil
	}
}

func (k KubeClient)CreatePvcByYaml(yamlPath string)error{
	pvcYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	pvc := core_v1.PersistentVolumeClaim{}

	pvcJson ,err := yaml.ToJSON(pvcYaml)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(pvcJson,&pvc);err != nil {
		return err
	}
	//////////////////////查询
	if _, err := k.Client.CoreV1().PersistentVolumeClaims(pvc.Namespace).Get(pvc.Name,meta_v1.GetOptions{}); err != nil {
		if errors.IsNotFound(err){
			if _, err := k.Client.CoreV1().PersistentVolumeClaims(pvc.Namespace).Create(&pvc);err != nil {
				return err
			}
		}
		return err
	}else{  /////////////// 如果存在，更新
		if _,err := k.Client.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(&pvc); err != nil {
			return err
		}
		return nil
	}
}