package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware 用于跨域请求
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		//var filterHost = [...]string{"http://*.hfjy.com"}
		// filterHost 做过滤器，防止不合法的域名访问
		// var isAccess = false
		// for _, v := range filterHost {
		// 	match, _ := regexp.MatchString(v, origin)
		// 	if match {
		// 		isAccess = true
		// 	}
		// }
		if true {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
