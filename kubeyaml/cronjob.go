package kubeyaml


import (
	"io/ioutil"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	batch_v1 "k8s.io/api/batch/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k KubeClient)CreateUpdateCronJob(yamlPath string)error{
	cronjobYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	cronjob := batch_v1.CronJob{}

	cronjobJson , err := yaml.ToJSON(cronjobYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(cronjobJson,&cronjob); err != nil {
		return err
	}
	/////////////////////get
	if _,err := k.Client.BatchV1beta1().CronJobs(cronjob.Namespace).Get(cronjob.Name,meta_v1.GetOptions{});err != nil {
		if errors.IsNotFound(err){
			/////////////////////create
			if _,err := k.Client.BatchV1beta1().CronJobs(cronjob.Namespace).Create(&cronjob);err != nil {
				return err
			}else{
				return nil
			}
		}
		return err
	}else{
		////////////////////////update
		if _,err := k.Client.BatchV1beta1().CronJobs(cronjob.Namespace).Update(&cronjob);err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)GetCronJob(namespace,name string)(*batch_v1.CronJob,error){
	if job,err := k.Client.BatchV1beta1().CronJobs(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return nil, err
	}else{
		return job,nil
	}
}

func (k KubeClient)DeleteCronJob(namespace,name string) error {
	if _,err := k.Client.BatchV1beta1().CronJobs(namespace).Get(name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		////////////////////////delete
		if err := k.Client.BatchV1beta1().CronJobs(namespace).Delete(name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}

func (k KubeClient)DeleteCronJobByYaml(yamlPath string)error{
	cronjobYaml,err  := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	cronjob := batch_v1.CronJob{}

	cronjobJson , err := yaml.ToJSON(cronjobYaml)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(cronjobJson,&cronjob); err != nil {
		return err
	}
	if _,err := k.Client.BatchV1beta1().CronJobs(cronjob.Namespace).Get(cronjob.Name,meta_v1.GetOptions{});err != nil {
		return err
	}else{
		////////////////////////delete
		if err := k.Client.BatchV1beta1().CronJobs(cronjob.Namespace).Delete(cronjob.Name,&meta_v1.DeleteOptions{});err != nil {
			return err
		}else{
			return nil
		}
	}
}
