package web

import (
	"github.com/gin-gonic/gin"
)

type ErrorApi struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func (e *ErrorApi) Error() string {
	return e.Message
}

func NewApiError(c *gin.Context, status int, code string, message string) {
	c.JSON(status, ErrorApi{
		Status: status,
		Code: code,
		Message: message,
	})
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Data: data,
	})
}
