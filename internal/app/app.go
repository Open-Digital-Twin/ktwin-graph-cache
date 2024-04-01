package app

import (
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/config"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cluster"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/server"
)

func StartApp() {
	config.Load()
	cacheConnection := cache.NewCacheConnection()

	clusterListeners := cluster.NewClusterListener(cacheConnection)
	clusterListeners.Listen()

	httpServer := server.NewHttpServer(cacheConnection)
	httpServer.Configure()
	httpServer.Start()
}
