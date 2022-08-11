package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name" binding:"required"`
	Default          bool               `json:"default" bson:"default"`
	GroupId          primitive.ObjectID `json:"groupId" bson:"groupId" binding:"required"`
}

type PaymentMethodUpdate struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Default bool   `json:"default" bson:"default"`
}
