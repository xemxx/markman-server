package router

import (
	"github.com/gin-gonic/gin"

	"markman-server/api/user"
	"markman-server/middleware/auth"
	"markman-server/middleware/cors"
	"markman-server/tools/config"
)

// InitRouter .
func InitRouter() *gin.Engine {
	gin.SetMode(config.Cfg.GetString("app.run_mode"))
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(cors.CorsMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signin", user.SignIn)
	r.POST("/signup", user.SignUp)
	ur := r.Group("/user")
	ur.Use(auth.CheckToken())
	{
		ur.POST("/info", user.Info)
	}

	return r
}
