package kubeyaml

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubefabric/utils"
)

type KubeClient struct {
	Client *kubernetes.Clientset
}

func InitClient()(KubeClient){
	configPath,err := utils.GetCurrentPath()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(configPath)
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return KubeClient{
		Client:clientset,
	}
}

