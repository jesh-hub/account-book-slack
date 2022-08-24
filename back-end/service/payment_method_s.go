package service

import (
	"abs/model"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPaymentMethod(groupId string, paymentMethodAdd *model.PaymentMethodAdd) (*model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethod := paymentMethodAdd.ToEntity()

	groupObjectId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}
	paymentMethod.GroupId = groupObjectId

	err = paymentMethodColl.Create(paymentMethod)
	return paymentMethod, err
}

func DeletePaymentMethodMany(groupId string) error {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	ctx := mgm.Ctx()

	groupObjectId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}

	results, err := paymentMethodColl.DeleteMany(ctx, bson.M{"groupId": groupObjectId}, nil)
	log.Printf("Delete paymentMethods : %d", results.DeletedCount)
	if err != nil {
		return err
	}
	return nil
}

func FindPaymentMethodByGroupId(groupId string) (*[]model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethods := &[]model.PaymentMethod{}

	groupObjectId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	err = paymentMethodColl.SimpleFind(paymentMethods, bson.M{"groupId": groupObjectId})
	return paymentMethods, err
}

func UpdatePaymentMethod(paymentMethodId string, paymentMethodUpdate *model.PaymentMethodUpdate) (*model.PaymentMethod, error) {
	paymentMethodColl := mgm.Coll(&model.PaymentMethod{})
	paymentMethod := &model.PaymentMethod{}

	err := paymentMethodColl.FindByID(paymentMethodId, paymentMethod)
	if err != nil {
		return nil, err
	}

	paymentMethodUpdate.UpdateEntity(paymentMethod)
	err = paymentMethodColl.Update(paymentMethod)
	return paymentMethod, err
}
