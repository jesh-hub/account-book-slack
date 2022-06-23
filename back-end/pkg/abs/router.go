package abs

import (
	"github.com/gin-gonic/gin"
)

func NewAbsRouterV1(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.POST("/api/user", Signup)
	}
	return v1
}
