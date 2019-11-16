package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"markman-server/tools/e"
	"markman-server/tools/jwt"
)

//CheckToken ..
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		claims := jwt.Claims{}
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Set("uid", claims.UID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
