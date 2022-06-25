package main

import (
	"abs/internal/abs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	abs.NewAbsRouterV1(r)
	if err := r.Run(":8080"); err != nil {
		return
	}
}
