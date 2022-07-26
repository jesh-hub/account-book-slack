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
		v1.GET("/group/:id", controller.FindGroupById)
		v1.GET("/group", controller.FindGroupByEmail)
		v1.POST("/group", controller.AddGroup)
		v1.PUT("/group/:id", controller.UpdateGroup)
		// payment method
		v1.GET("/paymentMethod", controller.FindPaymentMethodByGroupId)
		v1.POST("/paymentMethod", controller.AddPaymentMethod)
		v1.PUT("/paymentMethod/:id", controller.UpdatePaymentMethod)
		// Payment
		v1.GET("/payment", controller.FindPayment)
		v1.POST("/payment", controller.AddPayment)
		v1.PUT("/payment/:id", controller.UpdatePayment)
	}
	return v1
}
