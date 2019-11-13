package kubeyaml

import (
	"io/ioutil"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	batch_v1 "k8s.io/api/batch/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdateJob(yamlPath string)error{
	jobYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	job := batch_v1.Job{}

	jobJson , err := yaml.ToJSON(jobYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jobJson,&job); err != nil {
		return err
	}
	/////////////////////get
	if _,err := k.Client.BatchV1().Jobs(job.Namespace).Get(job.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){
			/////////////////////create
			if _,err := k.Client.BatchV1().Jobs(job.Namespace).Create(&job);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{
		////////////////////////update
		if _,err := k.Client.BatchV1().Jobs(job.Namespace).Update(&job);err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)GetJob(namespace,name string)(*batch_v1.Job,error){
	if job,err := k.Client.BatchV1().Jobs(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return nil, err
	}else{
		return job,nil
	}
}

func (k KubeClient)DeleteJob(namespace,name string) error {
	if _,err := k.Client.BatchV1().Jobs(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		////////////////////////delete
		if err := k.Client.BatchV1().Jobs(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)DeleteJobByYaml(yamlPath string)error{
	jobYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	job := batch_v1.Job{}

	jobJson , err := yaml.ToJSON(jobYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jobJson,&job); err != nil {
		return err
	}
	if _,err := k.Client.BatchV1().Jobs(job.Namespace).Get(job.Name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		////////////////////////delete
		if err := k.Client.BatchV1().Jobs(job.Namespace).Delete(job.Name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)ListJobs(namespace string)([]batch_v1.Job,error){
	if list,err := k.Client.BatchV1().Jobs(namespace).List(meta_v1.ListOptions{});err != nil {
		return nil,err
	}else{
		return list.Items,nil
	}
}