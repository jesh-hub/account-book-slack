package abs

import (
	"github.com/gin-gonic/gin"
	"log"
)

func errorHandlerInternal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func errorHandler(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"message": err.Error(),
	})
}
