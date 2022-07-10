package service

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string           `json:"name" bson:"name" binding:"required"`
	Users            []string         `json:"users" bson:"users" binding:"required"`
	RegUserId        string           `json:"regUserId" bson:"regUserId"`
	ModUserId        string           `json:"modUserId" bson:"modUserId"`
	PaymentMethods   *[]PaymentMethod `json:"paymentMethods"`
}

type FindGroupParam struct {
	Id                string
	Email             string
	WithPaymentMethod string
}

type UpdateGroupParam struct {
	Id    string
	Group *Group
}

func NewFindGroupParam() FindGroupParam {
	return FindGroupParam{
		Id:                "",
		Email:             "",
		WithPaymentMethod: "false",
	}
}

func AddGroup(group *Group) (*Group, error) {
	groupColl := mgm.Coll(&Group{})
	err := groupColl.Create(group)
	return group, err
}

func FindGroupById(param FindGroupParam) (*Group, error) {
	groupColl := mgm.Coll(&Group{})
	group := &Group{}
	err := groupColl.FindByID(param.Id, group)

	if strings.ToLower(param.WithPaymentMethod) == "true" {
		paymentMethodParam := FindPaymentMethodParam{
			GroupId: group.ID.Hex(),
		}
		paymentMethods, _ := FindPaymentMethodByGroupId(paymentMethodParam)
		group.PaymentMethods = paymentMethods
	}
	return group, err
}

func FindGroupByEmail(param FindGroupParam) (*[]Group, error) {
	groupColl := mgm.Coll(&Group{})
	groups := &[]Group{}

	q := bson.M{"users": param.Email}
	var err error

	if strings.ToLower(param.WithPaymentMethod) == "true" {
		paymentMethodColl := mgm.Coll(&PaymentMethod{}).Name()
		err = groupColl.SimpleAggregate(
			groups,
			builder.Lookup(paymentMethodColl, "_id", "groupId", "paymentMethods"),
			bson.M{operator.Match: q},
		)
	} else {
		err = groupColl.SimpleFind(groups, q)
	}
	return groups, err
}

func UpdateGroup(param UpdateGroupParam) (*Group, error) {
	groupColl := mgm.Coll(&Group{})
	group := &Group{}

	err := groupColl.FindByID(param.Id, group)
	if err != nil {
		return nil, err
	}

	err = groupColl.Update(param.Group)
	return param.Group, err
}
