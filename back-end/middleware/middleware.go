package middleware

import (
	"abs/util"
	"github.com/gin-gonic/gin"
)

var allowOrigin = util.GodotEnv("ALLOW_ORIGIN")

func SetHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", allowOrigin)
	c.Header("Cache-Control", "no-store")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
