package server

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, appController Controller) {
	v1 := r.Group("/api/v1")
	{
		tg := v1.Group("/twin-graph")
		{
			tg.GET("", appController.GetTwinGraph)
			tg.GET("/:interfaceId", appController.GetTwinGraphByTwinInstance)
		}
	}
	r.Group("/health").GET("", GetHealth)
}
