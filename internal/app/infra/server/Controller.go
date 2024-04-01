package server

import (
	twingraph "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph"
	cache "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"k8s.io/client-go/rest"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetTwinGraph(g *gin.Context)
	GetTwinGraphByTwinInstance(g *gin.Context)
	UpdateTwinGraph(result rest.Result)
}

func NewAppController(cacheConnection cache.CacheConnection) Controller {
	return &controller{cacheConnection: cacheConnection}
}

type controller struct {
	cacheConnection cache.CacheConnection
}

// Twin Graph

func (c *controller) GetTwinGraph(g *gin.Context) {
	container := twingraph.InitializeTwinGraphContainer(c.cacheConnection)
	container.Controller.GetTwinGraph(g)
}

func (c *controller) GetTwinGraphByTwinInstance(g *gin.Context) {
	container := twingraph.InitializeTwinGraphContainer(c.cacheConnection)
	container.Controller.GetTwinGraphByTwinInstance(g)
}

func (c *controller) UpdateTwinGraph(result rest.Result) {
	container := twingraph.InitializeTwinGraphContainer(c.cacheConnection)
	container.Controller.UpdateTwinGraph(result)
}
