package user

import (
	"markman-server/service/user"
	"markman-server/tools/common"
	"markman-server/tools/response"

	"github.com/gin-gonic/gin"
)

// FlashToken 每次登录新发放新token
func FlashToken(c *gin.Context) {
	//解析登录参数
	code := response.SUCCESS
	data := make(map[string]interface{})
	uid := c.GetInt("uid")
	username := c.GetString("username")
	//生成新token并且返回
	newToken, err := common.GenerateToken(username, uid)
	if err != nil {
		code = response.ErrorAuthToken
	} else {
		//保存用户新token
		_ = user.SaveToken(uid, newToken)
		data["token"] = newToken
	}
	response.JSON(c, code, response.GetMsg(code), data)
}

// 获取最新同步标志
func GetLastSyncCount(c *gin.Context) {
	code, data := response.SUCCESS, make(map[string]interface{})
	uid := c.GetInt("uid")

	SC, err := user.GetLastSC(uid)
	if err != nil {
		code = response.ERROR
	} else {
		data["SC"] = SC
	}
	response.JSON(c, code, response.GetMsg(code), data)
}
