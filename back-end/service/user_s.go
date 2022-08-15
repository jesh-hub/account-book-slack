package service

import (
	"abs/model"
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginParam struct {
	Credential string `json:"credential" binding:"required"`
}

func Login(param *LoginParam) (*model.User, error) {
	// Parse jwt token
	claims, err := JwtDecode(param.Credential)
	if err != nil {
		return nil, err
	}

	// Check user is already exist
	email := fmt.Sprintf("%v", claims["email"])
	userColl := mgm.Coll(&model.User{})
	user := &model.User{}
	err = userColl.First(bson.M{"email": email}, user)
	if err == mongo.ErrNoDocuments {
		user.Email = email
		err = userColl.Create(user)
	}

	// Set value(name, picture) from google
	user = setUserinfo(user, claims)
	return user, err
}

func setUserinfo(user *model.User, claims map[string]interface{}) *model.User {
	user.FirstName = fmt.Sprintf("%v", claims["family_name"])
	user.LastName = fmt.Sprintf("%v", claims["given_name"])
	user.Picture = fmt.Sprintf("%v", claims["picture"])
	return user
}
