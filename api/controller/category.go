package controller

import (
	"github.com/cloudy-11/backend/api/middleware"
	"github.com/cloudy-11/backend/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CategoryController struct {
	CategoryUserCase domain.CategoryUseCase
}

func (cc *CategoryController) Create(c *gin.Context) {

	role := c.GetString(middleware.X_USER_ROLE)
	if role != domain.ADMIN_ROLE {
		c.JSON(http.StatusForbidden, domain.ErrorResponse{
			Message: "Not permission",
		})
		return
	}
	var request domain.Category
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	request.ID = primitive.NewObjectID()
	err = cc.CategoryUserCase.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse[domain.Category]{
		Message: "Created Successful",
		Data:    request,
	})
}

func (cc *CategoryController) FetchById(c *gin.Context) {
	category, err := cc.CategoryUserCase.FetchById(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse[domain.Category]{
		Message: "Successfully",
		Data:    *category,
	})
}
