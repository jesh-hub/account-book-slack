package api

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Login(c *gin.Context) {
	loginParam := &service.LoginParam{}
	if err := c.ShouldBindJSON(loginParam); err != nil {
		c.JSON(http.StatusBadRequest, util.NewAppError(err))
		return
	}

	user, err := service.Login(loginParam)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
		return
	}
	c.JSON(http.StatusOK, user)
}
