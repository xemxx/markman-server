package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"markman-server/api/v1/node"
	"markman-server/api/v1/note"
	"markman-server/api/v1/notebook"
	"markman-server/api/v1/user"
	"markman-server/api/v1/version"
	"markman-server/middleware"
	"markman-server/tools/config"
)

// InitRouter .
func InitRouter() *gin.Engine {
	cfg := config.Cfg
	//不存在时也为debug模式
	gin.SetMode(cfg.App.RunMode)
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(middleware.CorsMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/version", version.GetVersion)
	r.POST("/signIn", user.SignIn)
	r.POST("/signUp", user.SignUp)

	// user
	ur := r.Group("/user")
	ur.Use(middleware.Auth())
	{
		ur.POST("/flashToken", user.FlashToken)
		ur.GET("/getLastSyncCount", user.GetLastSyncCount)
	}

	// notebook
	nbr := r.Group("/notebook")
	nbr.Use(middleware.Auth())
	{
		nbr.GET("/getSync", notebook.GetSync)
		nbr.POST("/create", notebook.Create)
		nbr.POST("/delete", notebook.Delete)
		nbr.POST("/update", notebook.Update)
	}

	// note
	nr := r.Group("/note")
	nr.Use(middleware.Auth())
	{
		nr.GET("/getSync", note.GetSync)
		nr.POST("/create", note.Create)
		nr.POST("/delete", note.Delete)
		nr.POST("/update", note.Update)
	}

	// node (统一的节点API)
	noder := r.Group("/node")
	noder.Use(middleware.Auth())
	{
		noder.GET("/getSync", node.GetSync)
		noder.POST("/create", node.Create)
		noder.POST("/delete", node.Delete)
		noder.POST("/update", node.Update)
	}

	return r
}
