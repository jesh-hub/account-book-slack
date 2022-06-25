package abs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type User struct {
	Id      string             `json:"id" bson:"id"`
	Name    string             `json:"name" bson:"name"`
	RegDate primitive.DateTime `json:"regDate" bson:"regDate" `
	ModDate primitive.DateTime `json:"modDate" bson:"modDate"`
}

func Signup(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	fmt.Println(user.Id)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
