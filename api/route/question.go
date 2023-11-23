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

func NewQuestionRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, public *gin.RouterGroup, private *gin.RouterGroup) {
	qr := repository.NewQuestionRepository(db, domain.CollectionQuestion)
	cr := repository.NewCategoryRepository(db, domain.CollectionCategory)
	qc := controller.QuestionController{
		QuestionUseCase: usecase.NewQuestionUseCase(qr, cr, timeout),
	}

	private.POST("/question", qc.Create)
	public.GET("/question/:id", qc.FetchById)
	public.GET("/question", qc.Fetch)
}
