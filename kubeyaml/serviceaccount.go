package kubeyaml

import (
	"io/ioutil"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdateServiceAccount(yamlPath string)error{

	accountYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	account := core_v1.ServiceAccount{}

	accountJson , err := yaml.ToJSON(accountYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(accountJson,&account); err != nil {
		return err
	}
	//////// get
	if _,err := k.Client.CoreV1().ServiceAccounts(account.Namespace).Get(account.Name,meta_v1.GetOptions{});err!=nil {
		if errors.IsNotFound(err){
			////// create
			if _,err := k.Client.CoreV1().ServiceAccounts(account.Namespace).Create(&account);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{
		////////update
		if _,err := k.Client.CoreV1().ServiceAccounts(account.Namespace).Update(&account);err != nil {
			return err
		}else{
			return nil
		}
	}
}
