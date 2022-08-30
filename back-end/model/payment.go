package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	mgm.DefaultModel   `bson:",inline"`
	Date               primitive.DateTime `json:"date" bson:"date"`
	Name               string             `json:"name" bson:"name"`
	Category           string             `json:"category" bson:"category"`
	Price              int                `json:"price" bson:"price"`
	MonthlyInstallment int                `json:"monthlyInstallment" bson:"monthlyInstallment"`
	PaymentMethodId    primitive.ObjectID `json:"paymentMethodId" bson:"paymentMethodId"`
	GroupId            primitive.ObjectID `json:"groupId" bson:"groupId"`
	RegUserId          string             `json:"regUserId" bson:"regUserId"`
	ModUserId          string             `json:"modUserId" bson:"modUserId"`
	PaymentMethods     *[]PaymentMethod   `json:"paymentMethods"`
}

type PaymentAdd struct {
	Date               primitive.DateTime `json:"date" binding:"required"`
	Name               string             `json:"name" binding:"required"`
	Category           string             `json:"category" binding:"required"`
	Price              int                `json:"price" binding:"required"`
	MonthlyInstallment int                `json:"monthlyInstallment"`
	PaymentMethodId    primitive.ObjectID `json:"paymentMethodId" binding:"required"`
	RegUserId          string             `json:"regUserId" bson:"regUserId" binding:"required"`
}

func (pa *PaymentAdd) ToEntity() *Payment {
	return &Payment{
		Date:               pa.Date,
		Name:               pa.Name,
		Price:              pa.Price,
		Category:           pa.Category,
		RegUserId:          pa.RegUserId,
		PaymentMethodId:    pa.PaymentMethodId,
		MonthlyInstallment: pa.MonthlyInstallment,
	}
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

func (pu *PaymentUpdate) UpdateEntity(payment *Payment) {
	payment.Date = pu.Date
	payment.Name = pu.Name
	payment.Price = pu.Price
	payment.Category = pu.Category
	payment.ModUserId = pu.ModUserId
	payment.PaymentMethodId = pu.PaymentMethodId
	if pu.MonthlyInstallment > 0 {
		payment.MonthlyInstallment = pu.MonthlyInstallment
	}
}

type PaymentStatistics struct {
	TotalIncome      int `json:"totalIncome" bson:"totalIncome"`
	TotalExpenditure int `json:"totalExpenditure" bson:"totalExpenditure"`
}
