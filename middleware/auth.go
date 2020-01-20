package middleware

import (
	"markman-server/tools/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"markman-server/tools/response"
)

type Auth struct {

}

//CheckToken ..
func (c *Auth) CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = response.SUCCESS
		token := c.Request.Header.Get("Authorization")
		var claims *common.Claims
		if token == "" {
			code = response.ErrorAuth
		} else {
			// fix variable is nil 必须先声明err,否则claims将被推断为局部变量
			var err error
			claims, err = common.ParseToken(token)
			if err != nil {
				//logs.Log("token鉴权失败：",err)
				code = response.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.ErrorAuthCheckTokenTimeout
			}
		}

		if code != response.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  response.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Set("token",token)
		c.Set("uid", claims.UID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
