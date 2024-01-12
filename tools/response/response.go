package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON .
// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码;
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func JSON(c *gin.Context, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": responseData,
	})
}
