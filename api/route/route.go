package route

import (
	"github.com/cloudy-11/backend/api/middleware"
	"github.com/cloudy-11/backend/bootstrap"
	"github.com/cloudy-11/backend/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	// Public APIs
	publicRouter := gin.Group("api/v1")
	NewLoginRouter(env, timeout, db, publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// Private APIs
	protectedRouter := gin.Group("api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewUserRouter(env, timeout, db, protectedRouter)

	NewCategoryRoute(env, timeout, db, publicRouter, protectedRouter)
}
