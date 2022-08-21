package service

import (
	"abs/model"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

func NewFindGroupParam() model.GroupFind {
	return model.GroupFind{
		Id:                "",
		Email:             "",
		WithPaymentMethod: false,
	}
}

func AddGroup(groupAdd *model.GroupAdd) (*model.Group, error) {
	groupColl := mgm.Coll(&model.Group{})

	// Add group
	group := groupAdd.ToEntity()
	err := groupColl.Create(group)
	if err != nil {
		return nil, err
	}

	// Add paymentMethods
	if groupAdd.PaymentMethodAdd != nil {
		for _, paymentMethodAdd := range *groupAdd.PaymentMethodAdd {
			_, err := AddPaymentMethod(group.ID.Hex(), &paymentMethodAdd)
			if err != nil {
				log.Info(err)
			}
		}
	}

	return group, err
}

func DeleteGroup(groupId string) error {
	groupColl := mgm.Coll(&model.Group{})
	group, err := FindGroupById(model.GroupFind{Id: groupId})
	if err != nil {
		return err
	}

	// 그룹 삭제
	err = groupColl.Delete(group)
	if err != nil {
		return err
	}
	return nil
}

func FindGroupById(groupFind model.GroupFind) (*model.Group, error) {
	groupColl := mgm.Coll(&model.Group{})
	group := &model.Group{}
	err := groupColl.FindByID(groupFind.Id, group)

	if groupFind.WithPaymentMethod {
		paymentMethods, _ := FindPaymentMethodByGroupId(groupFind.Id)
		group.PaymentMethods = paymentMethods
	}
	return group, err
}

func FindGroupByEmail(param model.GroupFind) (*[]model.Group, error) {
	groupColl := mgm.Coll(&model.Group{})
	groups := &[]model.Group{}

	q := bson.M{"users": param.Email}

	var err error
	if param.WithPaymentMethod {
		paymentMethodColl := mgm.Coll(&model.PaymentMethod{}).Name()
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

func UpdateGroup(id string, groupUpdate *model.GroupUpdate) (*model.Group, error) {
	groupColl := mgm.Coll(&model.Group{})
	group := &model.Group{}

	err := groupColl.FindByID(id, group)
	if err != nil {
		return nil, err
	}

	groupUpdate.ToEntity(group)
	err = groupColl.Update(group)
	return group, err
}
