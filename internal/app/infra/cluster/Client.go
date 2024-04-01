package cluster

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewClusterClient() ClusterClient {
	return &clusterClient{}
}

type ClusterClient interface {
	GetResources(absPath, namespace, resource string) (rest.Result, error)
}

type clusterClient struct {
	isConnected bool
	clientSet   *kubernetes.Clientset
}

func (c *clusterClient) localConnect() error {
	if c.isConnected {
		return nil
	}

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}

	// create the client set
	c.clientSet, err = kubernetes.NewForConfig(config)

	if err != nil {
		return err
	}

	return nil
}

func (c *clusterClient) connect() error {
	if c.isConnected {
		return nil
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// creates the clientset
	c.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

func (c *clusterClient) GetResources(absPath, namespace, resource string) (rest.Result, error) {
	var err error

	if os.Getenv("LOCAL") == "true" {
		err = c.localConnect()
	} else {
		err = c.connect()
	}

	if err != nil {
		return rest.Result{}, err
	}

	result := c.clientSet.RESTClient().Get().AbsPath(absPath).Namespace(namespace).Resource(resource).Do(context.TODO())

	if result.Error() != nil {
		return rest.Result{}, result.Error()
	}

	return result, nil
}
