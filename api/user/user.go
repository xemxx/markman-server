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
	// log.Println(uid)
	u, err := user.GetUserInfo(uid)
	if err != nil {
		response.JSON(c, e.ERROR_NOT_EXIST_USER_NAME, e.GetMsg(e.ERROR_NOT_EXIST_USER_NAME))
		return
	}
	// log.Println(u)
	response.JSON(c, e.SUCCESS, e.GetMsg(e.SUCCESS), u)
}
