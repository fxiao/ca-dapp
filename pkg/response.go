package pkg

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type defaultResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (g *Gin) Response(args ...interface{}) {

	statusCode := 200
	if code, ok := args[0].(int); ok {
		statusCode = code
	}

	message := StatusCodeMags[statusCode]

	if len(args) == 2 {

		if msg, ok := args[1].(string); ok {
			message = msg
		} else {
			g.C.JSON(statusCode, args[1])
			return
		}
	}

	g.C.JSON(statusCode, defaultResponse{
		StatusCode: statusCode,
		Message:    message,
	})

	return
}
