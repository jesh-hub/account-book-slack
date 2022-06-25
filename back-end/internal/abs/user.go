package abs

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

type LoginParameter struct {
	Credential string `json:"credential"`
}

func Login(c *gin.Context) {
	var loginParameter LoginParameter
	if err := c.BindJSON(&loginParameter); err != nil {
		errHandlerClient(c, err)
		return
	}

	if len(loginParameter.Credential) == 0 {
		errHandlerClient(c, errors.New("token is nil"))
		return
	}

	token, _ := jwt.Parse(loginParameter.Credential, nil)
	if token == nil {
		errorHandlerServer(c, errors.New("cannot parse token"))
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	c.JSON(http.StatusOK, loginParameter)
}
