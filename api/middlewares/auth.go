package middleware

import (
	"fmt"
	"net/http"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateRequest(ctx *gin.Context) {

	//Get the token from context
	clientToken, err := utils.ExtractBearerToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": utils.NewErrorStruct(http.StatusUnauthorized, fmt.Sprintf("unauthorized access %v,", err.Error())),
		})
		ctx.Abort()
		return
	}

	if clientToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Unauthorised user",
		})
		ctx.Abort()

		return
	}

	claims, err := utils.ValidateToken(clientToken)

	if err != nil {
		if clientToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorised user",
			})
		}

		ctx.JSON(400, gin.H{
			"Message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("username", claims.UserName)
	ctx.Set("userID", claims.UserID)
	ctx.Next()

}
