package model

import "github.com/kamva/mgm/v3"

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string           `json:"name" bson:"name"`
	Users            []string         `json:"users" bson:"users"`
	RegUserId        string           `json:"regUserId" bson:"regUserId"`
	ModUserId        string           `json:"modUserId" bson:"modUserId"`
	PaymentMethods   *[]PaymentMethod `json:"paymentMethods,omitempty" bson:"paymentMethods,omitempty"`
}

type GroupAdd struct {
	Name             string              `json:"name" bson:"name" binding:"required"`
	Users            []string            `json:"users" bson:"users" binding:"required"`
	RegUserId        string              `json:"regUserId" bson:"regUserId" binding:"required"`
	PaymentMethodAdd *[]PaymentMethodAdd `json:"paymentMethodAdd,omitempty" bson:"paymentMethodAdd,omitempty"`
}

func (ga *GroupAdd) ToEntity() *Group {
	return &Group{
		Name:      ga.Name,
		Users:     ga.Users,
		RegUserId: ga.RegUserId,
	}
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

func (gu *GroupUpdate) ToEntity(group *Group) {
	group.Name = gu.Name
	group.Users = gu.Users
	group.ModUserId = gu.ModUserId
}
