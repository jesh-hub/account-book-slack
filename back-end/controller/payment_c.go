package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPayment(c *gin.Context) {
	payment := &service.Payment{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	payment, err := service.AddPayment(payment)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, payment)
}
func FindPayment(c *gin.Context) {
	param := service.FindPaymentParam{
		DateFrom: c.Query("dateFrom"),
		DateTo:   c.Query("dateTo"),
		GroupId:  c.Query("groupId"),
	}

	payments, err := service.FindPayment(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, payments)
}
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	payment := &service.Payment{}

	if err := c.ShouldBindJSON(&payment); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	param := service.UpdatePaymentParam{
		Id:      id,
		Payment: payment,
	}

	payment, err := service.UpdatePayment(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, payment)
}
