package response

import (
	"markman-server/tools/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON .
// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func JSON(c *gin.Context, err int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	if err != e.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err,
			"msg":  msg,
			"data": responseData,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err":  err,
			"msg":  msg,
			"data": responseData,
		})
	}
}
