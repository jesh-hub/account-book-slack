package service

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Picture          string `json:"picture"`
}

type LoginParam struct {
	Credential string `json:"credential" binding:"required"`
}

func Login(param LoginParam) (*User, error) {
	// Parse jwt token
	claims, err := JwtDecode(param.Credential)
	if err != nil {
		return nil, err
	}

	// Check user is already exist
	email := fmt.Sprintf("%v", claims["email"])
	userColl := mgm.Coll(&User{})
	user := &User{}
	err = userColl.First(bson.M{"email": email}, user)
	if err == mongo.ErrNoDocuments {
		err = userColl.Create(user)
	}

	// Set value(name, picture) from google
	user = setUserinfo(user, claims)
	return user, err
}

func setUserinfo(user *User, claims map[string]interface{}) *User {
	user.FirstName = fmt.Sprintf("%v", claims["family_name"])
	user.LastName = fmt.Sprintf("%v", claims["given_name"])
	user.Picture = fmt.Sprintf("%v", claims["picture"])
	return user
}
