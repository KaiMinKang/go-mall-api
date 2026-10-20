package response

import (
	"mall-api/internal/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &response{
		Code:    errcode.CodeSuccess,
		Message: "success",
		Data:    data,
	})
}
func Fail(c *gin.Context, bizCode int, httpStatus int, message string) {
	c.JSON(httpStatus, &response{
		Code:    bizCode,
		Message: message,
		Data:    nil,
	})
}
