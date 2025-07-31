package response

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, statusCode int, success bool, message string, data interface{}, meta *Meta) {
	c.JSON(statusCode, Response{
		Success: success,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func NewMeta(total, page, limit int) *Meta {
	if limit <= 0 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	pageCount := (total + limit - 1) / limit

	return &Meta{
		Total:     total,
		Page:      page,
		PageCount: pageCount,
		HasNext:   page < pageCount,
		HasPrev:   page > 1,
	}
}
