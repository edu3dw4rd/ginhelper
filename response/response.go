package response

import (
	"math"
	"strconv"

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
	Meta paginationMeta `json:"meta"`
	Data interface{}    `json:"data"`
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

func PaginationResponse(code, total int, page, perPage string, data interface{}) *ModelPaginationResponse {
	res := new(ModelPaginationResponse)
	convPage, _ := strconv.Atoi(page)
	convPerPage, _ := strconv.Atoi(perPage)
	page_count := int(math.Ceil(float64(total) / float64(convPerPage)))
	hasNext := false

	if float64(convPage) < float64(page_count) {
		hasNext = true
	}

	meta := paginationMeta{
		metaResponse: metaResponse{
			Message: "success",
			Code:    code,
			Status:  true,
		},
		CurrentPage: convPage,
		NextPage:    hasNext,
		PrevPage:    convPage > 1,
		PerPage:     convPerPage,
		PageCount:   page_count,
		TotalCount:  total,
	}

	res.Meta = meta
	res.Data = data

	return res
}
