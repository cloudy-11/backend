package controller

import (
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/internal"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type SubmissionController struct {
	SubmissionUseCase domain.SubmissionUseCase
}

func (s *SubmissionController) Create(c *gin.Context) {
	// binding body
	userID := c.GetString(internal.X_USER_ID)
	var request domain.Submission
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	request.ID = primitive.NewObjectID()
	request.UserId = userID
	request.CreatedAt = time.Now()

	err = s.SubmissionUseCase.Create(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse[domain.Submission]{
		Message: "Created Successful",
		Data:    request,
	})
}

func (s *SubmissionController) FetchById(c *gin.Context) {
	submission, err := s.SubmissionUseCase.FetchById(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if submission == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse[domain.Submission]{
		Message: "Successfully",
		Data:    *submission,
	})
}

func (s *SubmissionController) Fetch(c *gin.Context) {
	var request domain.SubmissionQuery
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	submissions, err := s.SubmissionUseCase.Fetch(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if submissions == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Record not found",
		})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse[[]domain.Submission]{
		Message: "success",
		Data:    submissions,
	})
}
