package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddPaymentMethod
// @Summary Add paymentMethod
// @Description 결제수단 등록
// @Tags payment_method
// @Accept json
// @Produce json
// @Param paymentMethod body service.PaymentMethod true "PaymentMethod"
// @Success 200 {object} service.PaymentMethod
// @Router /v1/paymentMethod [post]
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

// FindPaymentMethodByGroupId
// @Summary Find paymentMethod by group id
// @Description group id로 결제수단 조회
// @Tags payment_method
// @Accept json
// @Produce json
// @Param groupId query string true "Group ID"
// @Success 200 {array} service.PaymentMethod
// @Router /v1/paymentMethod/{groupId} [get]
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

// UpdatePaymentMethod
// @Summary Update paymentMethod
// @Description 결제수단 수정
// @Tags payment_method
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod ID"
// @Param PaymentMethod body service.PaymentMethod true "PaymentMethod"
// @Success 200 {array} service.PaymentMethod
// @Router /v1/paymentMethod/{id} [put]
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
