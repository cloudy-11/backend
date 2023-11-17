package controller

import (
	"github.com/cloudy-11/backend/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (uc *UserController) Fetch(ctx *gin.Context) {
	userID := ctx.GetString("x-user-id")

	user, err := uc.UserUseCase.GetByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse[domain.User]{
		Message: "success",
		Data:    user,
	})
}
