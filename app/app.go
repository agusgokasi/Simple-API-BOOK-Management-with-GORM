package app

import (
	"fmt"
	"ninth-learn/config"
	"ninth-learn/repository"
	"ninth-learn/route"
	"ninth-learn/service"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
