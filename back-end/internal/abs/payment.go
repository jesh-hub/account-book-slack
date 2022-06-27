package abs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var paymentCollection = GetCollection(DB, "payment")
var paymentMethodCollection = GetCollection(DB, "paymentMethod")

type Payment struct {
	Id                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Date               primitive.DateTime `json:"date" bson:"date" binding:"required"`
	Name               string             `json:"name" bson:"name" binding:"required"`
	Category           string             `json:"category" bson:"category"`
	Method             PaymentMethod      `json:"method" bson:"method" binding:"required"`
	Price              int                `json:"price" bson:"price" binding:"required"`
	MonthlyInstallment int                `json:"monthlyInstallment" bson:"monthlyInstallment"`
	RegUserId          string             `json:"regUserId" json:"regUserId"`
	RegDate            primitive.DateTime `json:"regDate" bson:"regDate"`
	ModUserId          string             `json:"modUserId" bson:"modUserId"`
	ModDate            primitive.DateTime `json:"modDate" bson:"modDate"`
	GroupId            primitive.ObjectID `json:"group" bson:"group" binding:"required"`
}

type PaymentMethod struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name" binding:"required"`
	Default bool               `json:"default" bson:"default"`
	GroupId primitive.ObjectID `json:"groupId" bson:"groupId" binding:"required"`
}

func newPaymentMethod(name string, groupId primitive.ObjectID) PaymentMethod {
	return PaymentMethod{
		Name:    name,
		Default: false,
		GroupId: groupId,
	}
}

func AddPaymentMethod(c *gin.Context) {
	var paymentMethod PaymentMethod
	if err := c.ShouldBindJSON(&paymentMethod); err != nil {
		errorHandler(c, 400, err)
		return
	}

	result, err := insertOne(paymentMethodCollection, paymentMethod)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func FindPaymentMethodByGroupId(c *gin.Context) {
	id := c.Query("groupId")
	if len(id) == 0 {
		errorHandler(c, 400, errors.New("groupId is required in parameter"))
		return
	}
	groupId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	var paymentMethods []PaymentMethod
	findOptions := FindOptions{
		Filter: bson.M{"groupId": groupId},
	}
	err = findMany(paymentMethodCollection, findOptions, &paymentMethods)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, paymentMethods)

}
