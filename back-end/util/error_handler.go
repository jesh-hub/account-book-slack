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
	log.Printf("URI: %s, Message: %s\n", c.Request.RequestURI, err)
	c.JSON(code, gin.H{
		"message": err.Error(),
	})
}
