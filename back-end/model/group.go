package model

import "github.com/kamva/mgm/v3"

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string           `json:"name" bson:"name" binding:"required"`
	Users            []string         `json:"users" bson:"users" binding:"required"`
	RegUserId        string           `json:"regUserId" bson:"regUserId"`
	ModUserId        string           `json:"modUserId" bson:"modUserId"`
	PaymentMethods   *[]PaymentMethod `json:"paymentMethods,omitempty" bson:"paymentMethods,omitempty"`
}

type GroupFind struct {
	Id                string
	Email             string
	WithPaymentMethod bool
}

type GroupUpdate struct {
	Name      string   `json:"name" bson:"name" binding:"required"`
	Users     []string `json:"users" bson:"users" binding:"required"`
	ModUserId string   `json:"modUserId" bson:"modUserId" binding:"required"`
}
