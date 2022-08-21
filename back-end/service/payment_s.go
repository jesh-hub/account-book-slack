package service

import (
	"abs/model"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func AddPayment(groupId string, paymentAdd *model.PaymentAdd) (*model.Payment, error) {
	paymentColl := mgm.Coll(&model.Payment{})

	groupObjectId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	payment := paymentAdd.ToEntity()
	payment.GroupId = groupObjectId

	err = paymentColl.Create(payment)
	return payment, err
}

func FindPayment(groupId string, paymentFind model.PaymentFind) (*[]model.Payment, error) {
	paymentColl := mgm.Coll(&model.Payment{})
	payments := &[]model.Payment{}

	groupObjectId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	q := bson.M{
		"groupId": groupObjectId,
	}
	// when [start, end] parameter id existed
	if len(paymentFind.DateFrom) > 0 && len(paymentFind.DateTo) > 0 {
		startTime, _ := time.Parse("2006-01", paymentFind.DateFrom)
		endTime, _ := time.Parse("2006-01", paymentFind.DateTo)
		q["date"] = bson.M{
			"$gte": primitive.NewDateTimeFromTime(startTime),
			"$lt":  primitive.NewDateTimeFromTime(endTime),
		}
	}
	// when [priceFrom, priceTo] is existed
	if paymentFind.PriceFrom != 0 && paymentFind.PriceTo != 0 {
		q["price"] = bson.M{
			"$gte": paymentFind.PriceFrom,
			"$lt":  paymentFind.PriceTo,
		}
	}
	// -1: desc, 1: asc
	opts := options.Find()
	sort := bson.D{{"date", -1}}
	if len(paymentFind.OrderBy) != 0 {
		for k, v := range paymentFind.OrderBy {
			sort = append(sort, bson.E{Key: k, Value: v})
		}
	}
	opts.SetSort(sort)

	paymentMethodColl := mgm.Coll(&model.PaymentMethod{}).Name()
	err = paymentColl.SimpleAggregate(
		payments,
		builder.Lookup(paymentMethodColl, "paymentMethodId", "_id", "paymentMethods"),
		bson.M{operator.Match: q},
	)
	return payments, err
}

func UpdatePayment(paymentId string, paymentUpdate *model.PaymentUpdate) (*model.Payment, error) {
	paymentColl := mgm.Coll(&model.Payment{})
	payment := &model.Payment{}

	err := paymentColl.FindByID(paymentId, payment)
	if err != nil {
		return nil, err
	}

	paymentUpdate.UpdateEntity(payment)

	err = paymentColl.Update(payment)
	return payment, err
}
