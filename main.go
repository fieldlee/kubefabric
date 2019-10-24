package main

import (
	"fmt"


	//"time"

	//"github.com/urfave/cli"
	"kubefabric/env"
	"kubefabric/kubeutils"
	//"time"
	apiv1 "k8s.io/api/core/v1"
	//"k8s.io/client-go/discovery"

)

func main() {


/****
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

 */


	DeployFabric()
}

func DeployFabric(){
	namespace := "blockchain-fabric"
	//pvname := "fabric-pv"
	//pvcname := "fabric-pvc"
	//labelname := "fabricfiles"
	//
	//CreateNS(namespace)
	//time.Sleep(10 * time.Second)
	//
	//CreatePv(pvname,labelname,"192.168.1.100","/opt/nfs/fabric")
	//
	//time.Sleep(10 * time.Second)
	//
	//CreatePvc(namespace,pvcname,labelname)
	//time.Sleep(10 * time.Second)

	/////CREATE FABRIC TOOL
	//CreatePod(namespace,"fabric-tools","hyperledger/fabric-tools:latest","fabric-pvc",labelname)

	//CreateCa(namespace,"blockchain-ca")
	//CreateService(namespace,"blockchain-ca","blockchain-ca",7054,30054)

	//CreateOrder(namespace,"blockchain-orderer","hyperledger/fabric-orderer:latest","fabricfiles","/fabric","fabric-pvc",7050)
	///Delete Pod
	DeleteDeployment(namespace,"blockchain-orderer")
	DeleteDeployment(namespace,"blockchain-ca")
	DeletePod(namespace,"fabric-tools")
	DeleteService(namespace,"blockchain-ca")
	//DeletePvc(namespace,pvcname)
	//DeletePv(namespace,pvname)
	//DeleteNS(namespace)
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

func CreatePv(pvname,labelname,server,path string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePv(pvname,labelname,server,path)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(pv)

	result ,err := kubeClient.WatchPv(pvname)
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
	err := kubeClient.DeletePv(pvname)
	if err != nil {
		fmt.Println(err.Error())
	}
	result ,err := kubeClient.WatchPv(pvname)
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

func CreatePvc(namespace,pvcname ,labelname string){
	kubeClient := kubeutils.InitClient()
	pv,err := kubeClient.CreatePVC(namespace,pvcname, map[string]string{"name":labelname})
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

func CreatePod(namespace,podname,imagename,pvcname,labelname string){
	kubeClient := kubeutils.InitClient()
	VolList := make([]apiv1.Volume,0)
	fabricVolume := apiv1.Volume{
		Name:labelname,
		VolumeSource:apiv1.VolumeSource{
			PersistentVolumeClaim:&apiv1.PersistentVolumeClaimVolumeSource{
				ClaimName:pvcname,
			},
		},
	}

	dockersocket := apiv1.Volume{
		Name:"dockersocket",
		VolumeSource:apiv1.VolumeSource{
			HostPath:&apiv1.HostPathVolumeSource{
				Path:"/var/run/docker.sock",
			},
		},
	}
	VolList = append(VolList,fabricVolume,dockersocket)

	EnvList := make([]apiv1.EnvVar,0)

	path := apiv1.EnvVar{
		Name:"FABRIC_CFG_PATH",
		Value:"/fabric",
	}
	EnvList = append(EnvList,path)

	VolMounts := make([]apiv1.VolumeMount,0)

	volmount1 := apiv1.VolumeMount{
		Name:"fabricfiles",
		MountPath:"/fabric",
	}
	volmount2 := apiv1.VolumeMount{
		Name:"dockersocket",
		MountPath:"/host/var/run/docker.sock",
	}
	VolMounts = append(VolMounts,volmount1,volmount2)

	podinfo := kubeutils.PodInfo{
		Namespace:namespace,
		PodName:podname,
		ImageName:imagename,
		Volumes:VolList,
		Command:[]string{"sh", "-c", "sleep 48h"},
		Env:EnvList,
		VolumeMounts:VolMounts,
	}

	err := kubeClient.CreatePod(podinfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Pod error: %s",err.Error()))
	}

	result,err := kubeClient.WatchPod(namespace,podname)
	if err != nil {
		fmt.Println(fmt.Sprintf("Pod Watch error: %s",err.Error()))
	}
	fmt.Println(fmt.Sprintf("Pod Watch result: %d",result))
}

func DeletePod(namespace,podname string){
	kubeClient := kubeutils.InitClient()
	err := kubeClient.DeletePod(namespace,podname)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateCa(namespace,deployname string){
	kubeClient := kubeutils.InitClient()
	envlist := env.GenerateEnv("env_ca","./env/")
	deployment := kubeutils.DeploymentInfo{
		Namespace:namespace,
		DeploymentName:deployname,
		ImageName:"hyperledger/fabric-ca:latest",
		VolumnName:"fabricfiles",
		VolumnPath:"/fabric",
		PVCName:"fabric-pvc",
		ReplicaNum:1,
		Port:7054,
		Command:[]string{"sh", "-c", "fabric-ca-server start -b admin:adminpw -d"},
	}
	deploy,err := kubeClient.CreateDeployment(deployment,envlist)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(deploy)

	result ,err := kubeClient.WatchDeploy(namespace)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch result {
	case 1 :
		fmt.Println("创建ca完成")
	case -1 :
		fmt.Println("删除ca完成")
	case 0:
		fmt.Println("错误ca")
	default:
		fmt.Println("ca=====")
	}
}


func CreateService(namespace,svcname,appname string,port,nodeport int){
	kubeClient := kubeutils.InitClient()
	selector := map[string]string{
		"name":appname,
	}
	service := kubeutils.ServiceInfo{
		Namespace:namespace,
		ServiceName:svcname,
		Selector: selector,
		Port:port,
		Nodeport:nodeport,
	}


	svc,err := kubeClient.CreateServiceByPort(service)
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

	envlist := env.GenerateEnv("env_fabric_peer","./")

	deployment := kubeutils.DeploymentInfo{
		Namespace:namespace,
		DeploymentName:deployname,
		ImageName:imagename,
		VolumnName:volumnname,
		VolumnPath:volumnpath,
		PVCName:pvcname,
		ReplicaNum:1,
		Port:port,
		Command:[]string{"sh", "-c", "peer node start"},
	}

	vdeploy,err := kubeClient.CreateDeployment(deployment,envlist)
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

func CreateOrder(namespace ,deployname,imagename,volumnname,volumnpath,pvcname string,port int){
	kubeClient := kubeutils.InitClient()

	envlist := env.GenerateEnv("env_order","./env/")

	deployment := kubeutils.DeploymentInfo{
		Namespace:namespace,
		DeploymentName:deployname,
		ImageName:imagename,
		VolumnName:volumnname,
		VolumnPath:volumnpath,
		PVCName:pvcname,
		ReplicaNum:3,
		Port:port,
		Command:[]string{"sh", "-c", "orderer"},
	}

	vdeploy,err := kubeClient.CreateDeployment(deployment,envlist)
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
