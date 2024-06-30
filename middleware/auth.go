package middleware

import (
	"log"
	"markman-server/tools/common"
	"net/http"
	"time"

	"markman-server/tools/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CheckToken ..
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = response.SUCCESS
		token := c.Request.Header.Get("Authorization")
		log.Println(token)
		var claims *common.Claims
		if token == "" {
			code = response.ErrorAuthToken
		} else {
			// fix variable is nil 必须先声明err,否则claims将被推断为局部变量
			var err error
			claims, err = common.ParseToken(token)
			if err != nil {
				//logs.Log("token鉴权失败：",err)
				code = response.ErrorAuthCheckTokenFail
			} else {
				t, err := claims.GetExpirationTime()
				if err != nil {
					slog.Error("can not parse exo time from token, ", "err", err)
				}
				if time.Now().Unix() > t.Unix() {
					code = response.ErrorAuthCheckTokenTimeout
				}
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
		c.Set("token", token)
		c.Set("uid", claims.UID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
