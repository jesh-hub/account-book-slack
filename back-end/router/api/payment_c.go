package api

import (
	"abs/model"
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

// AddPayment
// @Summary Add payment
// @Description 결제내역 등록
// @Tags payment
// @Accept json
// @Produce json
// @Param groupId path string true "Group ID"
// @Param paymentAdd body model.PaymentAdd true "PaymentAdd"
// @Success 200 {object} model.Payment
// @Failure 400 {object} util.AppError
// @Failure 500 {object} util.AppError
// @Router /v1/group/{groupId}/payment [post]
func AddPayment(c *gin.Context) {
	groupId := c.Param("groupId")
	paymentAdd := &model.PaymentAdd{}
	if err := c.ShouldBindJSON(paymentAdd); err != nil {
		c.JSON(http.StatusBadRequest, util.NewAppError(err))
		return
	}

	payment, err := service.AddPayment(groupId, paymentAdd)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
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
// @Param groupId path string true "Group ID"
// @Param dateFrom query string false "Format like 2006-01"
// @Param dateTo query string false "Format like 2006-01"
// @Success 200 {array} model.Payment
// @Failure 500 {object} util.AppError
// @Router /v1/group/{groupId}/payment [get]
func FindPayment(c *gin.Context) {
	groupId := c.Param("groupId")
	paymentFind := model.PaymentFind{
		DateFrom: c.Query("dateFrom"),
		DateTo:   c.Query("dateTo"),
	}

	payments, err := service.FindPayment(groupId, paymentFind)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
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
// @Param paymentId path string true "Payment ID"
// @Param paymentUpdate body model.PaymentUpdate true "PaymentUpdate"
// @Success 200 {object} model.Payment
// @Failure 400 {object} util.AppError
// @Failure 500 {object} util.AppError
// @Router /v1/payment/{paymentId} [put]
func UpdatePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")
	paymentUpdate := &model.PaymentUpdate{}

	if err := c.ShouldBindJSON(paymentUpdate); err != nil {
		c.JSON(http.StatusBadRequest, util.NewAppError(err))
		return
	}

	payment, err := service.UpdatePayment(paymentId, paymentUpdate)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
		return
	}
	c.JSON(http.StatusOK, payment)
}
