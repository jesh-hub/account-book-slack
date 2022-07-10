package main

import (
	"abs/database"
	"abs/middleware"
	"abs/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	// run server
	r := gin.Default()
	r.Use(middleware.SetHeader)
	router.NewAbsRouterV1(r)
	if err := r.Run(":8080"); err != nil {
		return
	}
}
