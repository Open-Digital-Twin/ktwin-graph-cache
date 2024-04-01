package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCors(r *gin.Engine) {
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))
}
