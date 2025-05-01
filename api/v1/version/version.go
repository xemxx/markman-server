package version

import (
	"github.com/gin-gonic/gin"

	"markman-server/tools/config"
	"markman-server/tools/response"
)

// @Summary	Get Server Version
// @Schemes
// @Description	Get server version information
// @Tags			version
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Response{data=map[string]string}	"desc"
// @Router			/version [get]
func GetVersion(c *gin.Context) {
	code, data := response.SUCCESS, make(map[string]interface{})

	// 获取服务端版本信息
	data["version"] = config.Cfg.App.Version
	data["minClientVersion"] = config.Cfg.App.MinClientVersion

	response.JSON(c, code, response.GetMsg(code), data)
}
