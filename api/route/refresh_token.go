package route

import (
	"github.com/cloudy-11/backend/api/controller"
	"github.com/cloudy-11/backend/bootstrap"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"github.com/cloudy-11/backend/repository"
	"github.com/cloudy-11/backend/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUseCase: usecase.NewRefreshTokenUseCase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
