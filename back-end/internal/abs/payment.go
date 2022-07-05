package abs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Payment struct {
	mgm.DefaultModel   `bson:",inline"`
	Date               primitive.DateTime `json:"date" bson:"date" binding:"required"`
	Name               string             `json:"name" bson:"name" binding:"required"`
	Category           string             `json:"category" bson:"category"`
	Method             PaymentMethod      `json:"method" bson:"method" binding:"required"`
	Price              int                `json:"price" bson:"price" binding:"required"`
	MonthlyInstallment int                `json:"monthlyInstallment" bson:"monthlyInstallment"`
	RegUserId          string             `json:"regUserId" json:"regUserId"`
	ModUserId          string             `json:"modUserId" bson:"modUserId"`
	GroupId            primitive.ObjectID `json:"group" bson:"group" binding:"required"`
}

type PaymentMethod struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name" binding:"required"`
	Default          bool               `json:"default" bson:"default"`
	GroupId          primitive.ObjectID `json:"groupId" bson:"groupId" binding:"required"`
}

func AddPaymentMethod(c *gin.Context) {
	paymentMethod := &PaymentMethod{}
	if err := c.ShouldBindJSON(paymentMethod); err != nil {
		errorHandler(c, 400, err)
		return
	}

	paymentMethodColl := mgm.Coll(&PaymentMethod{})
	err := paymentMethodColl.Create(paymentMethod)
	if err != nil {
		errorHandler(c, 500, err)
		return
	}

	c.JSON(http.StatusOK, paymentMethod)
}

func FindPaymentMethodByGroupId(c *gin.Context) {
	groupId := c.Query("groupId")
	if len(groupId) == 0 {
		errorHandler(c, 400, errors.New("groupId is required in parameter"))
		return
	}

	paymentMethodColl := mgm.Coll(&PaymentMethod{})
	paymentMethods := &[]PaymentMethod{}
	err := paymentMethodColl.SimpleFind(paymentMethods, bson.M{"groupId": groupId})

	if err != nil {
		errorHandler(c, 500, err)
		return
	}

	c.JSON(http.StatusOK, paymentMethods)
}
