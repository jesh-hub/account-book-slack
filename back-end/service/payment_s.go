package service

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	mgm.DefaultModel   `bson:",inline"`
	Date               string             `json:"date"`
	Name               string             `json:"name"`
	Category           string             `json:"category"`
	Method             string             `json:"method"`
	Price              int                `json:"price"`
	MonthlyInstallment int                `json:"monthlyInstallment"`
	PaymentMethodId    primitive.ObjectID `json:"paymentMethodId" bson:"paymentMethodId"`
	GroupId            primitive.ObjectID `json:"groupId" bson:"groupId"`
	RegUserId          string             `json:"regUserId" bson:"regUserId"`
	ModUserId          string             `json:"modUserId" bson:"modUserId"`
}

func AddPayment(payment *Payment) (*Payment, error) {
	paymentColl := mgm.Coll(&Payment{})
	err := paymentColl.Create(payment)
	return payment, err
}

func ()  {
	
}
