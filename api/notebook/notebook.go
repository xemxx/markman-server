package notebook

import (
	"github.com/gin-gonic/gin"
	"log"
	"markman-server/service/notebook"
	"markman-server/tools/response"
	"strconv"
)

func GetSync(c *gin.Context) {
	afterSC, _ := strconv.Atoi(c.DefaultQuery("afterSC", "0"))
	maxCount, _ := strconv.Atoi(c.DefaultQuery("maxCount", "10"))
	uid := c.GetInt("uid")

	code, data := response.SUCCESS, make(map[string]interface{})
	notebooks, err := notebook.GetSync(uid, afterSC, maxCount)
	if err != nil {
		log.Println(err)
		code = response.ERROR
	} else {
		data["notebooks"] = notebooks
	}
	response.JSON(c, code, response.GetMsg(code), data)
}
