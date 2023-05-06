package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR       = 404
	SUCCESS     = 200
	ParamsError = 1002 // 参数错误
	TokenError  = 1003
	SQLError    = 1004
)

type Response struct {
	Code    int         `json:"code"`    // 自定义状态码
	Message string      `json:"message"` // 传递消息
	Data    interface{} `json:"data"`    // 传递数据
}

func responseJSON(c *gin.Context, httpCode int, code int, message string, data interface{}) {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(httpCode, response)
}

func Success(c *gin.Context, message string, data interface{}) {
	responseJSON(c, http.StatusOK, SUCCESS, message, data)
}

func Error(c *gin.Context, code int, message string) {
	responseJSON(c, http.StatusBadRequest, code, message, map[string]interface{}{})
}
