package service

import (
	"abs/model"
	"abs/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func AddPaymentMethod(groupId string, paymentMethod *model.PaymentMethod) (*model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethod.GroupId = util.ConvertStringToObjectId(groupId)
	err := paymentMethodColl.Create(paymentMethod)
	return paymentMethod, err
}

func FindPaymentMethodByGroupId(groupId string) (*[]model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethods := &[]model.PaymentMethod{}
	err := paymentMethodColl.SimpleFind(paymentMethods, bson.M{"groupId": util.ConvertStringToObjectId(groupId)})
	return paymentMethods, err
}

func UpdatePaymentMethod(paymentMethodId string, paymentMethodUpdate *model.PaymentMethodUpdate) (*model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethod := &model.PaymentMethod{}

	err := paymentMethodColl.FindByID(paymentMethodId, paymentMethod)
	if err != nil {
		return nil, err
	}

	paymentMethod.Name = paymentMethodUpdate.Name
	paymentMethod.Default = paymentMethodUpdate.Default

	err = paymentMethodColl.Update(paymentMethod)
	return paymentMethod, err
}
