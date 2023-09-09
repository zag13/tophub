package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/bootstrap"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(data ...interface{}) Response {
	response := Response{
		Code:    0,
		Message: "success",
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	return response
}

func Error(code int, errors ...any) Response {
	response := Response{
		Code:    code,
		Message: errorCodeMessageMap[ErrorCodeSystemBusy],
	}

	if msg, ok := errorCodeMessageMap[code]; ok {
		response.Message = msg
	}
	if len(errors) > 0 && bootstrap.Env.IsShowErrors {
		response.Errors = errors[0]
	}

	return response
}

func ToSuccessResponse(c *gin.Context, data ...any) {
	response := Success(data...)
	c.JSON(http.StatusOK, response)
}

func ToErrorResponse(c *gin.Context, code int, errors ...any) {
	response := Error(code, errors...)
	c.JSON(http.StatusBadRequest, response)
}

type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"pageSize"`
	TotalNumber int `json:"totalNumber,omitempty"`
	TotalPage   int `json:"totalPage,omitempty"`
}

func GetTotalPage(page, pageSize, totalNumber int) int {
	if pageSize == 0 {
		return 1
	}

	if totalNumber%pageSize == 0 {
		return page + totalNumber/pageSize - 1
	}

	return page + totalNumber/pageSize
}
