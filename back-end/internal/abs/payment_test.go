package abs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewPaymentMethod(t *testing.T) {
	var group Group

	findOptions := FindOptions{
		Filter: bson.M{"name": "test"},
	}
	err := findOne(groupCollection, findOptions, &group)

	paymentMethod := newPaymentMethod("카카오뱅크", group.Id)
	result, err := insertOne(paymentMethodCollection, paymentMethod)
	errorHandlerInternal(err)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)
}

func TestInsertManyPaymentMethod(t *testing.T) {
	var group Group

	findOptions := FindOptions{
		Filter: bson.M{"name": "test"},
	}
	err := findOne(groupCollection, findOptions, &group)

	paymentMethods := []PaymentMethod{newPaymentMethod("카카오뱅크", group.Id), newPaymentMethod("케케케", group.Id)}

	params := make([]interface{}, len(paymentMethods))
	for i, v := range paymentMethods {
		params[i] = v
	}
	result, err := insertMany(paymentMethodCollection, params)
	errorHandlerInternal(err)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)

}
