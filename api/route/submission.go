package route

import (
	"github.com/cloudy-11/backend/api/controller"
	"github.com/cloudy-11/backend/bootstrap"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"github.com/cloudy-11/backend/repository"
	"github.com/cloudy-11/backend/usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewSubmissionRote(env *bootstrap.Env, timeout time.Duration, db mongo.Database, public *gin.RouterGroup, private *gin.RouterGroup) {
	sr := repository.NewSubmissionRepository(db, domain.CollectionSubmission)
	qr := repository.NewQuestionRepository(db, domain.CollectionQuestion)
	sc := controller.SubmissionController{
		SubmissionUseCase: usecase.NewSubmissionUseCase(sr, qr, timeout),
	}

	private.POST("/submission", sc.Create)
	public.GET("/submission/:id", sc.FetchById)
	public.GET("/submission", sc.Fetch)
}
