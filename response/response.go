package response

import (
	"github.com/gin-gonic/gin"
)

type metaResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type paginationMeta struct {
	metaResponse
	CurrentPage int         `json:"current_page"`
	NextPage    interface{} `json:"next_page"`
	PrevPage    interface{} `json:"prev_page"`
	PerPage     int         `json:"per_page"`
	PageCount   int         `json:"page_count"`
	TotalCount  int         `json:"total_count"`
}

type modelResponse struct {
	Meta metaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type ModelPaginationResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

func isSuccessCode(statusCode int) bool {
	return statusCode > 100 && statusCode < 399
}

func NewResponse(ctx *gin.Context, statusCode int, msg string, content interface{}) {
	isSuccess := isSuccessCode(statusCode)

	respData := modelResponse{
		Meta: metaResponse{
			Status:  isSuccess,
			Code:    statusCode,
			Message: msg,
		},
		Data: content,
	}

	ctx.JSON(statusCode, respData)
}
