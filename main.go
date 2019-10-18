package main

import (
	"fmt"
	"github.com/urfave/cli"
	"kubefabric/kubeutils"
	"time"

	//"k8s.io/client-go/discovery"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "action",
			Value: "show",
			Usage: "action for kubernetes",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "show"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		switch name {
		case "show":
			ShowKube()
			break
		case "create":
			//CreateKubeDeployment()
			ListKubeDeployment()
			break
		case "update":
			UpdateKubeDeployment()
			ListKubeDeployment()
			break
		case "delete":
			DeleteKubeDeployment()
			ListKubeDeployment()
			break
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	//CreateNS()
	//CreatePv()
	//CreatePvc()
	namespace := "shared-services"
	servicename := "rabbitmq-nfs-poc-svc"
	pvname := "rabbitmq-nfs-pv"
	pvcname := "rabbitmq-nfs-pvc"
	deployname := "rabbitmq-depl"
	//DeleteNS(namespace)
	//DeletePv(namespace,pvname)
	//DeletePvc(namespace,pvcname)
	//DeleteService(namespace,servicename)
	//DeleteDeployment(namespace,deployname)

	CreateNS(namespace)
	time.Sleep(20*time.Second)
	CreatePv(namespace,pvname,"192.168.1.100","/opt/nfs/data/rabbitmq/")
	time.Sleep(20*time.Second)
	CreatePvc(namespace,pvname,pvcname)
	time.Sleep(20*time.Second)
	CreateService(namespace,servicename,deployname,5672,30672)
	time.Sleep(20*time.Second)
	CreateDeployment(namespace,deployname,"rabbitmq","rabbitmq-mnt","/opt/nfs/data/rabbitmq/",pvcname,5672)
}

func CreateNS(namespace string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.CreateNamespace(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}

	result ,err := kubeClient.WatchNamespace(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch result {
	case 1 :
		fmt.Println("创建完成")
	case -1 :
		fmt.Println("删除完成")
	case 0:
		fmt.Println("错误")
	default:
		fmt.Println("=====")
	}
}

func DeleteNS(namespace string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DelelteNamespace(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreatePv(namespace,pvname,serverip,path string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePv(namespace,pvname,serverip,path)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(pv)
}

func DeletePv(namespace,pvname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeletePv(namespace,pvname)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreatePvc(namespace,pvname,pvcname string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePVC(namespace,pvname, map[string]string{"app":pvcname})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(pv)

}

func DeletePvc(namespace,pvcname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DelectPvc(namespace,pvcname)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateService(namespace,svcname,appname string,port,nodeport int){
	kubeClient := kubeutils.InitClient()
	selector := map[string]string{
		"app":appname,
	}
	svc,err := kubeClient.CreateServiceByPort(namespace,svcname,selector,port,nodeport)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(svc)
}

func DeleteService(namespace,svcname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteService(namespace,svcname)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateDeployment(namespace ,deployname,imagename,volumnname,volumnpath,pvcname string,port int){
	kubeClient := kubeutils.InitClient()
	vdeploy,err := kubeClient.CreateDeployment(namespace,deployname,imagename,volumnname,volumnpath,pvcname,1,port)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(vdeploy)
}

func DeleteDeployment(namespace ,deployname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteDeployment(namespace ,deployname)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ShowKube(){
	kubeClient := kubeutils.InitClient()
	podList := kubeClient.GetPod("")

	for _,pod := range podList{
		fmt.Printf("pod name :%s pod type:%s namespace:%s\n",pod.Name,pod.Kind,pod.Namespace)
	}

	nodeList := kubeClient.GetNode()

	fmt.Println("===========================================================")
	for _,node := range nodeList{
		fmt.Printf("node name :%s node type:%s \n",node.Name,node.Kind)
	}
}
//func CreateKubeDeployment(){
//	kubeClient := kubeutils.InitClient()
//	deployment,err := kubeClient.CreateDeployment()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(deployment.Namespace,deployment.Name)
//}
func UpdateKubeDeployment(){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.UpdateDeployment("demo-deployment")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func ListKubeDeployment(){
	kubeClient := kubeutils.InitClient()
	listDeployment,err := kubeClient.ListDeployment()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _,deploy:=range listDeployment{
		fmt.Println(deploy.Namespace,deploy.Name)
	}
}
func DeleteKubeDeployment(){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteDeployment("demo-deployment","")
	if err != nil {
		fmt.Println(err.Error())
	}
}
