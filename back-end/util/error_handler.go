package util

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrorHandlerInternal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrorHandler(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"message": err.Error(),
	})
}
