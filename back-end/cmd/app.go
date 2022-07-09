package main

import (
	"abs/database"
	"abs/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	// run server
	r := gin.Default()
	router.NewAbsRouterV1(r)
	if err := r.Run(":8080"); err != nil {
		return
	}
}
