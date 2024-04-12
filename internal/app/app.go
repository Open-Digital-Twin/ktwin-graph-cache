package app

import (
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/pkg/log"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/config"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cluster"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/server"
)

func StartApp() {
	config.Load()
	cacheConnection := cache.NewCacheConnection()
	logger := log.NewLogger()

	listenTime := config.GetConfigInt("LISTEN_TIME", 60)
	clusterListeners := cluster.NewClusterListener(cacheConnection, logger)
	clusterListeners.Listen(listenTime)

	httpServer := server.NewHttpServer(cacheConnection)
	httpServer.Configure()
	httpServer.Start()
}
