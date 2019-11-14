package routers

import (
	"github.com/gin-gonic/gin"

	"markman-server/tools/config"
)

// InitRouter .
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(config.Cfg.GetString("run_mode"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
