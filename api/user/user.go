package user

import (
	"markman-server/service/user"
	"markman-server/tools/e"
	"markman-server/tools/response"

	"github.com/gin-gonic/gin"
)

// Info .
func Info(c *gin.Context) {
	uid := c.GetInt("uid")

	u := user.GetUserInfo(uid)

	response.JSON(c, e.SUCCESS, e.GetMsg(e.SUCCESS), u)
}
