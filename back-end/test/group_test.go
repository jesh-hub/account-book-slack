package test

import (
	"abs/model"
	"abs/service"
	"fmt"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	guestEmail = "guest@wejesh.com"
	groupId    string
	groupFind  = model.GroupFind{
		Email: guestEmail,
	}
)

func TestGetGroup(t *testing.T) {
	groups, err := service.FindGroupByEmail(groupFind)
	fmt.Printf("%# v\n", pretty.Formatter(groups))
	assert.Equal(t, err, nil)
}

func TestAddPaymentMethod(t *testing.T) {
	groups, err := service.FindGroupByEmail(groupFind)
	assert.Equal(t, err, nil)
	assert.Greater(t, len(groups), 0)
	groupId = groups[0].ID.Hex()

	paymentMethodAdd := &model.PaymentMethodAdd{
		Name: "신용카드",
	}

	paymentMethod, err := service.AddPaymentMethod(groupId, paymentMethodAdd)
	fmt.Printf("%# v\n", pretty.Formatter(paymentMethod))
	assert.Equal(t, err, nil)

	err = service.DeletePaymentMethodMany(groupId)
	assert.Equal(t, err, nil)
}
