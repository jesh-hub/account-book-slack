package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	param := service.LoginParam{}
	if err := c.ShouldBindJSON(param); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	user, err := service.Login(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
