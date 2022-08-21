package main

import (
	"abs/database"
	_ "abs/docs"
	"abs/middleware"
	"abs/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.Init()

	r := gin.Default()

	r.Use(middleware.AllowPreflight)
	r.Use(middleware.SetHeader)

	router.NewAbsRouterV1(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":8080"); err != nil {
		return
	}
}
