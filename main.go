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
	//pvname := "rabbitmq-nfs-pv"
	pvcname := "rabbitmq-nfs-pvc"
	deployname := "rabbitmq-depl"
	ingressname := "rabbitmq-ingress"

	//DeleteIngress(namespace,ingressname)
	//time.Sleep(10*time.Second)
	//DeleteDeployment(namespace,deployname)
	//time.Sleep(10*time.Second)
	//DeleteService(namespace,servicename)
	//time.Sleep(10*time.Second)
	//DeletePvc(namespace,pvcname)
	//time.Sleep(10*time.Second)
	//DeletePv(namespace,pvname)
	//time.Sleep(10*time.Second)
	//DeleteNS(namespace)

	//CreateNS(namespace)
	//time.Sleep(20*time.Second)
	//CreatePv(namespace,pvname,"192.168.1.100","/opt/nfs/data/rabbitmq/")
	//time.Sleep(20*time.Second)
	//CreatePvc(namespace,pvcname)
	time.Sleep(20*time.Second)
	CreateService(namespace,servicename,deployname,5672,30672)
	time.Sleep(20*time.Second)
	CreateDeployment(namespace,deployname,"rabbitmq","rabbitmq-mnt","/var/lib/rabbitmq/",pvcname,5672)
	time.Sleep(20*time.Second)
	CreateIngress(namespace,ingressname,servicename,30672)

}


func CreateIngress(namespace,ingName,svcname string,port int){
	kubeClient := kubeutils.InitClient()
	_,err := kubeClient.CreateIngress(namespace,ingName,svcname,"field.lee.com","/rabbit",port)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchIngress(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch result {
	case 1 :
		fmt.Println("创建Ingress完成")
	case -1 :
		fmt.Println("删除Ingress完成")
	case 0:
		fmt.Println("错误Ingress")
	default:
		fmt.Println("Ingress=====")
	}
}

func DeleteIngress(namespace,ingName string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteIngress(namespace,ingName)
	if err != nil {
		fmt.Println(err.Error())
	}
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
		fmt.Println("创建Namespace完成")
	case -1 :
		fmt.Println("删除Namespace完成")
	case 0:
		fmt.Println("错误Namespace")
	default:
		fmt.Println("Namespace=====")
	}
}

func DeleteNS(namespace string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DelelteNamespace(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchNamespace(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建Namespace完成")
	case -1 :
		fmt.Println("删除Namespace完成")
	case 0:
		fmt.Println("错误Namespace")
	default:
		fmt.Println("Namespace=====")
	}
}

func CreatePv(namespace,pvname,serverip,path string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePv(namespace,pvname,serverip,path)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(pv)

	result ,err := kubeClient.WatchPv(namespace,pvname)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建PV完成")
	case -1 :
		fmt.Println("删除PV完成")
	case 0:
		fmt.Println("错误PV")
	default:
		fmt.Println("PV=====")
	}
}

func DeletePv(namespace,pvname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeletePv(namespace,pvname)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchPv(namespace,pvname)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建PV完成")
	case -1 :
		fmt.Println("删除PV完成")
	case 0:
		fmt.Println("错误PV")
	default:
		fmt.Println("PV=====")
	}
}

func CreatePvc(namespace,pvcname string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePVC(namespace,pvcname, map[string]string{"app":pvcname})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(pv)
	result ,err := kubeClient.WatchPvc(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建PVC完成")
	case -1 :
		fmt.Println("删除PVC完成")
	case 0:
		fmt.Println("错误PVC")
	default:
		fmt.Println("PVC=====")
	}
}

func DeletePvc(namespace,pvcname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DelectPvc(namespace,pvcname)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchPvc(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建PVC完成")
	case -1 :
		fmt.Println("删除PVC完成")
	case 0:
		fmt.Println("错误PVC")
	default:
		fmt.Println("PVC=====")
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
	result ,err := kubeClient.WatchService(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建Service完成")
	case -1 :
		fmt.Println("删除Service完成")
	case 0:
		fmt.Println("错误Service")
	default:
		fmt.Println("Service=====")
	}
}

func DeleteService(namespace,svcname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteService(namespace,svcname)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchService(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建Service完成")
	case -1 :
		fmt.Println("删除Service完成")
	case 0:
		fmt.Println("错误Service")
	default:
		fmt.Println("Service=====")
	}
}

func CreateDeployment(namespace ,deployname,imagename,volumnname,volumnpath,pvcname string,port int){
	kubeClient := kubeutils.InitClient()
	vdeploy,err := kubeClient.CreateDeployment(namespace,deployname,imagename,volumnname,volumnpath,pvcname,1,port)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(vdeploy)
	result ,err := kubeClient.WatchDeploy(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建Deployment完成")
	case -1 :
		fmt.Println("删除Deployment完成")
	case 0:
		fmt.Println("错误Deployment")
	default:
		fmt.Println("Deployment=====")
	}
}

func DeleteDeployment(namespace ,deployname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeleteDeployment(namespace ,deployname)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchDeploy(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建Deployment完成")
	case -1 :
		fmt.Println("删除Deployment完成")
	case 0:
		fmt.Println("错误Deployment")
	default:
		fmt.Println("Deployment=====")
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
