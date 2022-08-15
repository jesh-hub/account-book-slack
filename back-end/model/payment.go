package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	mgm.DefaultModel   `bson:",inline"`
	Date               primitive.DateTime `json:"date" binding:"required"`
	Name               string             `json:"name" binding:"required"`
	Category           string             `json:"category"`
	Price              int                `json:"price" binding:"required"`
	MonthlyInstallment int                `json:"monthlyInstallment"`
	PaymentMethodId    primitive.ObjectID `json:"paymentMethodId" bson:"paymentMethodId" binding:"required"`
	GroupId            primitive.ObjectID `json:"groupId" bson:"groupId" binding:"required"`
	RegUserId          string             `json:"regUserId" bson:"regUserId" binding:"required"`
	ModUserId          string             `json:"modUserId" bson:"modUserId"`
	PaymentMethods     *[]PaymentMethod   `json:"paymentMethods"`
}

type PaymentFind struct {
	DateFrom  string         `json:"dateFrom"`
	DateTo    string         `json:"dateTo"`
	GroupId   string         `json:"groupId" binding:"required"`
	OrderBy   map[string]int `json:"orderBy"`
	PriceFrom int            `json:"priceFrom"`
	PriceTo   int            `json:"priceTo"`
}

type PaymentUpdate struct {
	Date               primitive.DateTime `json:"date" binding:"required"`
	Name               string             `json:"name" binding:"required"`
	Category           string             `json:"category" binding:"required"`
	Price              int                `json:"price" binding:"required"`
	MonthlyInstallment int                `json:"monthlyInstallment"`
	PaymentMethodId    primitive.ObjectID `json:"paymentMethodId" binding:"required"`
	ModUserId          string             `json:"modUserId" binding:"required"`
}
