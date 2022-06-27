package abs

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

const (
	USER_COLL = "user"
)

var userCollection = GetCollection(DB, "user")

type User struct {
	Id        string             `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Picture   string             `json:"picture"`
	RegDate   primitive.DateTime `json:"regDate" bson:"regDate" `
	ModDate   primitive.DateTime `json:"modDate" bson:"modDate"`
}

type LoginParameter struct {
	Credential string `json:"credential" binding:"required"`
}

func Login(c *gin.Context) {
	var loginParameter LoginParameter
	if err := c.BindJSON(&loginParameter); err != nil {
		errorHandler(c, 400, err)
		return
	}

	// Parse jwt token
	claims, err := JwtDecode(loginParameter.Credential)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	email := fmt.Sprintf("%v", claims["email"])
	// Check user is already exist
	user, err := findUserByEmail(email)
	if err != nil {
		// if user is not exist, add new user
		id := addUser(claims)
		user, err = findUserByEmail(id)
	}
	if err != nil {
		errorHandler(c, 500, err)
	}

	// Set value(name, picture) from google
	user = setUserinfo(user, claims)
	c.JSON(http.StatusOK, user)
}

func addUser(claims jwt.MapClaims) string {
	user := newUser(fmt.Sprintf("%v", claims["email"]))
	id, _ := insertOne(userCollection, user)
	return fmt.Sprintf("%v", id)
}

func findUserByEmail(email string) (User, error) {
	var user User

	findOptions := FindOptions{
		Filter: bson.D{{"_id", email}},
	}
	err := findOne(userCollection, findOptions, &user)
	if user == (User{}) {
		err = errors.New(fmt.Sprintf("User(%s) is not existed\n", email))
	}
	return user, err
}

func newUser(email string) User {
	return User{
		Id:      email,
		RegDate: primitive.NewDateTimeFromTime(time.Now()),
		ModDate: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func setUserinfo(user User, claims jwt.MapClaims) User {
	user.FirstName = fmt.Sprintf("%v", claims["family_name"])
	user.LastName = fmt.Sprintf("%v", claims["given_name"])
	user.Picture = fmt.Sprintf("%v", claims["picture"])
	return user
}
