package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS          = 1000
	ERROR            = 1001
	TokenExpired     = 1002
	PermissionDenied = 1003 // permission denied
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

func Success(c *gin.Context) {
	responseJSON(c, http.StatusOK, SUCCESS, "successful", map[string]interface{}{})
}

func SendSuccessData(c *gin.Context, message string, data interface{}) {
	responseJSON(c, http.StatusOK, SUCCESS, message, data)
}

func Error(c *gin.Context, message string) {
	responseJSON(c, http.StatusBadRequest, ERROR, message, map[string]interface{}{})
}

func SendErrorData(c *gin.Context, code int, message string) {
	responseJSON(c, http.StatusBadRequest, code, message, map[string]interface{}{})
}
