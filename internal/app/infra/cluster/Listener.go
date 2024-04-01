package cluster

import (
	"fmt"
	"time"

	twingraph "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
)

func NewClusterListener(cacheConnection cache.CacheConnection) ClusterListener {
	return &clusterListener{
		cacheConnection: cacheConnection,
		clusterClient:   NewClusterClient(),
	}
}

type ClusterListener interface {
	Listen()
}

type clusterListener struct {
	clusterClient   ClusterClient
	cacheConnection cache.CacheConnection
}

func (c *clusterListener) Listen() {
	go func() {
		for {
			c.listenTwinInstances()
			time.Sleep(100 * time.Second)
		}
	}()
}

func (c *clusterListener) listenTwinInstances() {
	result, err := c.clusterClient.GetResources("/apis/dtd.ktwin/v0", "ktwin", "twininstances")
	if err == nil {
		container := twingraph.InitializeTwinGraphContainer(c.cacheConnection)
		container.Controller.UpdateTwinGraph(result)
	} else {
		fmt.Errorf("Error getting resources from cluster: %v", err)
	}
}
