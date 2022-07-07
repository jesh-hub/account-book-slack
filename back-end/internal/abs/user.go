package abs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Picture          string `json:"picture"`
}

type LoginParameter struct {
	Credential string `json:"credential" binding:"required"`
}

func Login(c *gin.Context) {
	var loginParameter LoginParameter
	if err := c.ShouldBindJSON(&loginParameter); err != nil {
		errorHandler(c, 400, err)
		return
	}

	// Parse jwt token
	claims, err := JwtDecode(loginParameter.Credential)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	// Check user is already exist
	email := fmt.Sprintf("%v", claims["email"])
	userColl := mgm.Coll(&User{})
	user := &User{}
	err = userColl.First(bson.M{"email": email}, user)
	if err == mongo.ErrNoDocuments {
		err = userColl.Create(user)
	}
	if err != nil {
		errorHandler(c, http.StatusInternalServerError, err)
		return
	}

	// Set value(name, picture) from google
	user = setUserinfo(user, claims)
	c.JSON(http.StatusOK, user)
}

func setUserinfo(user *User, claims map[string]interface{}) *User {
	user.FirstName = fmt.Sprintf("%v", claims["family_name"])
	user.LastName = fmt.Sprintf("%v", claims["given_name"])
	user.Picture = fmt.Sprintf("%v", claims["picture"])
	return user
}
