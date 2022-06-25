package abs

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func errorHandlerServer(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func errHandlerClient(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"message": err.Error(),
	})
}
