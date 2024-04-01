package server

import (
	cache "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/middleware"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Configure()
	Start()
}

type httpServer struct {
	engine        *gin.Engine
	appController Controller
}

func NewHttpServer(cacheConnection cache.CacheConnection) HttpServer {
	appController := NewAppController(cacheConnection)
	return &httpServer{
		engine:        gin.Default(),
		appController: appController,
	}
}

func (s *httpServer) Configure() {
	middleware.UseCors(s.engine)
	ConfigureRoutes(s.engine, s.appController)
	ConfigureSwagger(s.engine)
}

func (s *httpServer) Start() {
	s.engine.Run(":8080")
}
