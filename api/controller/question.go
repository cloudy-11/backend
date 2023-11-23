package controller

import (
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/internal"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type QuestionController struct {
	QuestionUseCase domain.QuestionUseCase
}

func (q *QuestionController) Create(c *gin.Context) {
	// Check role
	role := c.GetString(internal.X_USER_ROLE)
	if role != domain.ADMIN_ROLE {
		c.JSON(http.StatusForbidden, domain.ErrorResponse{
			Message: "Not permission",
		})
		return
	}

	// binding body
	var request domain.Question
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	request.ID = primitive.NewObjectID()
	request.Slug = slug.Make(request.Title)
	request.IsLock = true

	err = q.QuestionUseCase.Create(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse[domain.Question]{
		Message: "Created Successful",
		Data:    request,
	})
}

func (q *QuestionController) Fetch(c *gin.Context) {
	var request domain.QuestionSearch
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	questions, err := q.QuestionUseCase.Fetch(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if questions == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Record not found",
		})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse[[]domain.Question]{
		Message: "success",
		Data:    questions,
	})
}

func (q *QuestionController) FetchById(c *gin.Context) {
	question, err := q.QuestionUseCase.FetchById(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if question == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse[domain.Question]{
		Message: "Successfully",
		Data:    *question,
	})
}
