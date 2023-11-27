package main

import (
	"github.com/cloudy-11/backend/api/route"
	"github.com/cloudy-11/backend/bootstrap"
	"github.com/gin-contrib/cors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	// CORS
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(env.Origins, ","),
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	route.Setup(env, timeout, db, gin)
	// For heroku deploy
	port := ":" + os.Getenv("PORT")
	// For server deploy
	//gin.Run(env.ServerAddress)

	gin.Run(port)
}
