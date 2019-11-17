package auth

import (
	"net/http"
	"time"

	"markman-server/tools/e"
	"markman-server/tools/jwt"

	"github.com/gin-gonic/gin"
)

//CheckToken ..
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		var claims *jwt.Claims
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			// TODO: fix parseToken not return claims
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
		//log.Println("first:", claims)
		c.Set("username", claims.Username)
		c.Next()
	}
}
