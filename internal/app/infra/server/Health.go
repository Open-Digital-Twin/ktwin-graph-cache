package server

import (
	"github.com/gin-gonic/gin"
)

func GetHealth(g *gin.Context) {
	g.JSON(200, gin.H{
		"status": "OK",
	})
}
