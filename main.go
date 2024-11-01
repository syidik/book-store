package main

import (
	"fmt"

	"book-store/internal/config"
	"book-store/internal/datasource"
	"book-store/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.InitAppConfig(); err != nil {
		fmt.Print(err.Error())
	}
}

func main() {
	dbClient, err := datasource.SetupDB(
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresPort,
		config.AppConfig.PostgresDB,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPwd,
	)

	if err != nil {
		// return nil, err
	}

	//setup router
	var mode = gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	router := gin.New()
	// router.Use(middlewares.CORSMiddleware())
	router.Use(gin.Recovery())

	//setup api
	api := router.Group("api")
	routes.NewRoute(api, dbClient).Routes()
	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
}
