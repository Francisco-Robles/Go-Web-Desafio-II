package web

import (
	"net/http"

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

func NewNotFoundApiError(message string) error {
	return &ErrorApi{http.StatusNotFound, "not_found", message}
}

func NewBadRequestApiError(message string) error {
	return &ErrorApi{http.StatusBadRequest, "bad_request", message}
}

func NewConflictApiError(message string) error {
	return &ErrorApi{http.StatusConflict, "Conflict", message}
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Data: data,
	})
}
