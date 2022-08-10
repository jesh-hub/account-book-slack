package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddPayment
// @Summary Add payment
// @Description 결제내역 등록
// @Tags payment
// @Accept json
// @Produce json
// @Param Payment body service.Payment true "Payment"
// @Success 200 {object} service.Payment
// @Router /v1/payment [post]
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

// FindPayment
// @Summary Find payment
// @Description 결제내역 조회
// @Description Date 관련 파라미터 있을 경우 : DateFrom <= 데이터 < DateTo
// @Description Date 관련 파라미터 없을 경우 : 전체 기간 조회
// @Tags payment
// @Accept json
// @Produce json
// @Param DateFrom query string false "Format like 2006-01"
// @Param DateTo query string false "Format like 2006-01"
// @Param GroupId query string true "Group Id"
// @Success 200 {array} service.Payment
// @Router /v1/payment [get]
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

// UpdatePayment
// @Summary Update payment
// @Description 결제내역 수정
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body service.Payment true "Payment"
// @Success 200 {array} service.Payment
// @Router /v1/payment/{id} [put]
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
