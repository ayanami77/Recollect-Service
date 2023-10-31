package myerror

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	InvalidRequest      = errors.New("invalid request")
	InternalServerError = errors.New("internal server error")
)

func HandleError(c *gin.Context, err error) {
	Message := err.Error()
	statusCode := getStatusCode(err)

	c.JSON(statusCode, gin.H{"error": Message})
}

func getStatusCode(err error) int {
	var errorStatusMap = map[error]int{
		InvalidRequest:      400,
		InternalServerError: 500,
	}

	if code, exists := errorStatusMap[err]; exists {
		return code
	}
	return 500 // デフォルトのステータスコード
}
