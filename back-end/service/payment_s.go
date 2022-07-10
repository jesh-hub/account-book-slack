package service

import (
	"abs/util"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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

type FindPaymentParam struct {
	DateFrom  string         `json:"dateFrom"`
	DateTo    string         `json:"dateTo"`
	GroupId   string         `json:"groupId" binding:"required"`
	OrderBy   map[string]int `json:"orderBy"`
	PriceFrom int            `json:"priceFrom"`
	PriceTo   int            `json:"priceTo"`
}

type UpdatePaymentParam struct {
	Id      string
	Payment *Payment
}

func AddPayment(payment *Payment) (*Payment, error) {
	paymentColl := mgm.Coll(&Payment{})
	err := paymentColl.Create(payment)
	return payment, err
}

func FindPayment(param FindPaymentParam) (*[]Payment, error) {
	paymentColl := mgm.Coll(&Payment{})
	payments := &[]Payment{}

	q := bson.M{
		"groupId": util.ConvertStringToObjectId(param.GroupId),
	}
	// when [start, end] parameter exist
	if len(param.DateFrom) > 0 && len(param.DateTo) > 0 {
		startTime, _ := time.Parse("2006-01", param.DateFrom)
		endTime, _ := time.Parse("2006-01", param.DateTo)
		q["date"] = bson.M{
			"$gte": primitive.NewDateTimeFromTime(startTime),
			"$lt":  primitive.NewDateTimeFromTime(endTime),
		}
	}
	// when [pricefrom, priceto] is exist
	if param.PriceFrom != 0 && param.PriceTo != 0 {
		q["price"] = bson.M{
			"$gte": param.PriceFrom,
			"$lt":  param.PriceTo,
		}
	}
	// -1: desc, 1: asc
	opts := options.Find()
	sort := bson.D{{"date", -1}}
	if len(param.OrderBy) != 0 {
		for k, v := range param.OrderBy {
			sort = append(sort, bson.E{Key: k, Value: v})
		}
	}
	opts.SetSort(sort)

	paymentMethodColl := mgm.Coll(&PaymentMethod{}).Name()
	err := paymentColl.SimpleAggregate(
		payments,
		builder.Lookup(paymentMethodColl, "paymentMethodId", "_id", "paymentMethods"),
		bson.M{operator.Match: q},
	)
	return payments, err
}

func UpdatePayment(param UpdatePaymentParam) (*Payment, error) {
	paymentColl := mgm.Coll(&Payment{})
	payment := &Payment{}

	err := paymentColl.FindByID(param.Id, payment)
	if err != nil {
		return nil, err
	}

	err = paymentColl.Update(param.Payment)
	return param.Payment, err
}
