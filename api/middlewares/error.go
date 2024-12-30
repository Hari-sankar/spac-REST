package middleware

import (
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			// Handle other errors
			ctx.JSON(err.Err.(utils.ErrorStruct).Code, err.Err)
			ctx.Abort()
		}
	}
}
