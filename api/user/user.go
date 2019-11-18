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
	u, err := user.GetUserInfo(uid)
	if err != nil {
		response.JSON(c, e.ERROR_NOT_EXIST_USER_NAME, e.GetMsg(e.ERROR_NOT_EXIST_USER_NAME))
		return
	}
	response.JSON(c, e.SUCCESS, e.GetMsg(e.SUCCESS), map[string]string{
		"username":    u.Username,
		"create_time": u.CreateTime.Format("2006-01-02 15:04:05"),
	})
}
