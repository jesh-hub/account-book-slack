package abs

import (
	"github.com/gin-gonic/gin"
)

const URL_PREFIX = "/abs"

func NewAbsRouterV1(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group(URL_PREFIX + "/v1")
	{
		v1.POST("/login", Login)
		// group
		v1.GET("/group/:id", FindGroupById)
		v1.GET("/group", FindGroupByEmail)
		v1.POST("/group", AddGroup)
		v1.PUT("/group/:id", UpdateGroup)
		// payment method
		v1.GET("/paymentMethod", FindPaymentMethodByGroupId)
		v1.POST("/paymentMethod", AddPaymentMethod)
	}

	return v1
}
