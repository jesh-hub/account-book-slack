package abs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestInsertGroup(t *testing.T) {
	users := []string{"jee824k@gmail.com"}
	group := newGroup("test", users, "jee824k@gmail.com")

	result, err := insertOne(groupCollection, group)
	errorHandlerInternal(err)
	fmt.Println(prettyPrint(result))

	paymentMethod := newPaymentMethod("현대신용", result.(primitive.ObjectID))

	results, err := insertOne(paymentMethodCollection, paymentMethod)
	errorHandlerInternal(err)

	fmt.Println(prettyPrint(results))
	assert.NotEmpty(t, result)
}

func TestUpdateGroup(t *testing.T) {
	var group Group

	findOptions := FindOptions{
		Filter: bson.M{"name": "test"},
	}
	err := findOne(groupCollection, findOptions, &group)
	errorHandlerInternal(err)

	group.ModUserId = "test@test.com"
	group.ModDate = primitive.NewDateTimeFromTime(time.Now())

	groupId, err := primitive.ObjectIDFromHex(group.Id.Hex())

	fmt.Println(groupId)
	result, err := replaceOne(groupCollection, groupId, group)
	if err != nil {
		errorHandlerInternal(err)
	}

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)
}
