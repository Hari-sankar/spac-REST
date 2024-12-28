package controllers

import (
	"net/http"
	"spac-REST/api/schemas"
	"spac-REST/api/usecases"
	"spac-REST/api/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthController(authUseCase usecases.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

func (ctrl *AuthController) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var req schemas.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}

	response, err := ctrl.authUseCase.SignIn(ctx, req.Username, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// funtion to handle the user signup
func (ctrl *AuthController) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var err error
	var response *schemas.SignUpResponse
	var req schemas.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	if response, err = ctrl.authUseCase.SignUp(ctx, &req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}
