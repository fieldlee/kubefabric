package kubeyaml

import (
	"io/ioutil"
	"encoding/json"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func (k KubeClient)ListNode() ([]v1.Node,error) {
	list,err := k.Client.CoreV1().Nodes().List(meta_v1.ListOptions{})
	if err != nil {
		return nil,err
	}

	return list.Items,nil
}

func (k KubeClient)GetNode(name string)(*v1.Node,error){
	node,err := k.Client.CoreV1().Nodes().Get(name,meta_v1.GetOptions{})
	if err != nil {
		return nil,err
	}

	return node,nil
}

func (k KubeClient)CreateUpdateNodeByYaml(yamlPath string)error{
	nodeYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	node := v1.Node{}

	nodeJson , err := yaml.ToJSON(nodeYaml)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(nodeJson,&node); err != nil {
		return err
	}
	//////////////////////?get
	if _,err := k.Client.CoreV1().Nodes().Get(node.Name,meta_v1.GetOptions{});err != nil{
		if errors.IsNotFound(err){ ////////////////// create node
			if _,err := k.Client.CoreV1().Nodes().Create(&node);err != nil{
				return err
			}
		}
		return  err
	}else{
		//////////////// update
		_,err := k.Client.CoreV1().Nodes().Update(&node)
		if err != nil {
			return err
		}
		return nil
	}
}