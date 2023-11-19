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

func NewCategoryRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, public *gin.RouterGroup, private *gin.RouterGroup) {
	cr := repository.NewCategoryRepository(db, domain.CollectionCategory)
	cc := controller.CategoryController{
		CategoryUserCase: usecase.NewCategoryUseCase(cr, timeout),
	}

	private.POST("/category", cc.Create)
	public.GET("/category/:id", cc.FetchById)
	public.GET("/category", cc.Fetch)
}
