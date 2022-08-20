package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Default          bool               `json:"default" bson:"default"`
	GroupId          primitive.ObjectID `json:"groupId" bson:"groupId"`
}

type PaymentMethodAdd struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Default bool   `json:"default" bson:"default"`
}

func (pma *PaymentMethodAdd) ToEntity() *PaymentMethod {
	return &PaymentMethod{
		Name:    pma.Name,
		Default: pma.Default,
	}
}

type PaymentMethodUpdate struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Default bool   `json:"default" bson:"default"`
}

func (pmu *PaymentMethodUpdate) ToEntity(paymentMethod *PaymentMethod) {
	paymentMethod.Name = pmu.Name
	paymentMethod.Default = pmu.Default
}
