package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"markman-server/api/user"
	"markman-server/middleware"
	"markman-server/tools/config"
)

// InitRouter .
func InitRouter() *gin.Engine {
	cfg:=config.Cfg
	//不存在时也为debug模式
	gin.SetMode(cfg.GetString("app.run_mode"))
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	mCors:=middleware.Cors{}
	r.Use(mCors.CorsMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signIn", user.SignIn)
	r.POST("/signUp", user.SignUp)
	ur := r.Group("/user")
	mAuth:=middleware.Auth{}
	ur.Use(mAuth.CheckToken())
	{
		ur.POST("/info", user.Info)
		ur.POST("/flashToken", user.FlashToken)
	}

	return r
}
