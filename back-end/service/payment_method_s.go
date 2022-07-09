package service

import (
	"abs/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name" binding:"required"`
	Default          bool               `json:"default" bson:"default"`
	GroupId          primitive.ObjectID `json:"groupId" bson:"groupId" binding:"required"`
}

type FindPaymentMethodParam struct {
	GroupId string
}

type UpdatePaymentMethodParam struct {
	Id            string
	PaymentMethod *PaymentMethod
}

func NewFindPaymentMethodParam() FindPaymentMethodParam {
	return FindPaymentMethodParam{
		GroupId: "",
	}
}

func AddPaymentMethod(paymentMethod *PaymentMethod) (*PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&PaymentMethod{})
	err := paymentMethodColl.Create(paymentMethod)
	return paymentMethod, err
}

func FindPaymentMethodByGroupId(param FindPaymentMethodParam) (*[]PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&PaymentMethod{})
	paymentMethods := &[]PaymentMethod{}
	err := paymentMethodColl.SimpleFind(paymentMethods, bson.M{"groupId": util.ConvertStringToObjectId(param.GroupId)})
	return paymentMethods, err
}

func UpdatePaymentMethod(param UpdatePaymentMethodParam) (*PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&PaymentMethod{})
	paymentMethod := &PaymentMethod{}

	err := paymentMethodColl.FindByID(param.Id, paymentMethod)
	if err != nil {
		return nil, err
	}

	err = paymentMethodColl.Update(param.PaymentMethod)
	return param.PaymentMethod, err
}
