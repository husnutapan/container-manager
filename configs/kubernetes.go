package configs

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type ClientConfig struct {
	kubeConfigUrl *string
	clientSet     *kubernetes.Clientset
}

func (conf *ClientConfig) GetClientSet() *kubernetes.Clientset {
	return conf.clientSet
}

func (conf *ClientConfig) CreateClient() {
	config, configUrl, err := createConfig()
	if err != nil {
		panic(err.Error())
	}
	conf.kubeConfigUrl = configUrl
	conf.clientSet, err = kubernetes.NewForConfig(config)
}

func createConfig() (*rest.Config, *string, error) {
	var configUrl *string
	if dir := homeDirectory(); dir != "" {
		configUrl = flag.String("kubeconfig", filepath.Join(dir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		configUrl = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *configUrl)
	return config, configUrl, err
}

func homeDirectory() string {
	var directory string
	if home := os.Getenv("HOME"); home != "" {
		directory = home
	} else {
		directory = os.Getenv("USERPROFILE")
	}
	return directory
}
