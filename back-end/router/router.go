package router

import (
	"abs/controller"
	"github.com/gin-gonic/gin"
)

func NewAbsRouterV1(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.POST("/login", controller.Login)
		// group
		v1.GET("/group/:groupId", controller.FindGroupById)
		v1.GET("/group", controller.FindGroup)
		v1.POST("/group", controller.AddGroup)
		v1.PUT("/group/:groupId", controller.UpdateGroup)
		// payment method
		v1.GET("/group/:groupId/paymentMethod", controller.FindPaymentMethodByGroupId)
		v1.POST("/group/:groupId/paymentMethod", controller.AddPaymentMethod)
		v1.PUT("/paymentMethod/:paymentMethodId", controller.UpdatePaymentMethod)
		// Payment
		v1.GET("/group/:groupId/payment", controller.FindPayment)
		v1.POST("/group/:groupId/payment", controller.AddPayment)
		v1.PUT("/payment/:paymentId", controller.UpdatePayment)
	}
	return v1
}
