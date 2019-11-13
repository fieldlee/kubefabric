package kubeyaml

import (
	"io/ioutil"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	ext_v1 "k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdateIngress(yamlPath string)error{

	ingressYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	ingress := ext_v1.Ingress{}

	ingressJson , err := yaml.ToJSON(ingressYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(ingressJson,&ingress); err != nil {
		return err
	}
	////////get
	if _,err := k.Client.ExtensionsV1beta1().Ingresses(ingress.Namespace).Get(ingress.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err) {
			if _, err := k.Client.ExtensionsV1beta1().Ingresses(ingress.Namespace).Create(&ingress);err != nil {
				return  err
			}else {
				return nil
			}
		}
		return err
	}else{
		if _, err := k.Client.ExtensionsV1beta1().Ingresses(ingress.Namespace).Update(&ingress);err != nil {
			return  err
		}else {
			return nil
		}
	}
}
