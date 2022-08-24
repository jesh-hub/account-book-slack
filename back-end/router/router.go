package router

import (
	"abs/router/api"
	"github.com/gin-gonic/gin"
)

func NewAbsRouterV1(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.POST("/login", api.Login)
		// group
		v1.GET("/group/:groupId", api.FindGroupById)
		v1.GET("/group", api.FindGroup)
		v1.POST("/group", api.AddGroup)
		v1.PUT("/group/:groupId", api.UpdateGroup)
		v1.DELETE("/group/:groupId", api.DeleteGroup)
		// payment method
		v1.GET("/group/:groupId/paymentMethod", api.FindPaymentMethodByGroupId)
		v1.POST("/group/:groupId/paymentMethod", api.AddPaymentMethod)
		v1.PUT("/paymentMethod/:paymentMethodId", api.UpdatePaymentMethod)
		// Payment
		v1.GET("/group/:groupId/payment", api.FindPayment)
		v1.POST("/group/:groupId/payment", api.AddPayment)
		v1.PUT("/payment/:paymentId", api.UpdatePayment)
	}
	return v1
}
