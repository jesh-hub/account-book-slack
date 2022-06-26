package abs

import (
	"github.com/gin-gonic/gin"
)

const URL_PREFIX = "/abs"

func NewAbsRouterV1(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group(URL_PREFIX + "/v1")
	{
		v1.POST("/login", Login)
	}

	return v1
}
