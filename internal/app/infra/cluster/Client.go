package cluster

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/pkg/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewClusterClient(logger log.Logger) ClusterClient {
	return &clusterClient{
		logger: logger,
	}
}

type ClusterClient interface {
	GetResources(absPath, namespace, resource string) (rest.Result, error)
}

type clusterClient struct {
	isConnected bool
	clientSet   *kubernetes.Clientset
	logger      log.Logger
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
		c.logger.Error(fmt.Sprintf("Error building kubeconfig: %v", err))
		return err
	}

	// create the client set
	c.clientSet, err = kubernetes.NewForConfig(config)

	if err != nil {
		c.logger.Error(fmt.Sprintf("Error creating client set: %v", err))
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
		c.logger.Error(fmt.Sprintf("Error creating in-cluster config: %v", err))
		return err
	}

	// creates the clientset
	c.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error creating client set: %v", err))
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
