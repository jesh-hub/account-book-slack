package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPaymentMethod(c *gin.Context) {
	paymentMethod := &service.PaymentMethod{}
	if err := c.ShouldBindJSON(paymentMethod); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	paymentMethod, err := service.AddPaymentMethod(paymentMethod)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, paymentMethod)
}

func FindPaymentMethodByGroupId(c *gin.Context) {
	param := service.NewFindPaymentMethodParam()
	param.GroupId = c.Query("groupId")

	paymentMethods, err := service.FindPaymentMethodByGroupId(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, paymentMethods)
}

func UpdatePaymentMethod(c *gin.Context) {
	id := c.Param("id")
	paymentMethod := &service.PaymentMethod{}

	if err := c.ShouldBindJSON(&paymentMethod); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	param := service.UpdatePaymentMethodParam{
		Id:            id,
		PaymentMethod: paymentMethod,
	}

	paymentMethod, err := service.UpdatePaymentMethod(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, paymentMethod)
}
