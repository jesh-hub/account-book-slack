package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGroup(c *gin.Context) {
	group := &service.Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	group, err := service.AddGroup(group)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func FindGroupById(c *gin.Context) {
	param := service.NewFindGroupParam()
	param.Id = c.Param("id")
	param.WithPaymentMethod = c.Query("withPaymentMethod")

	group, err := service.FindGroupById(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func FindGroupByEmail(c *gin.Context) {
	param := service.NewFindGroupParam()
	param.Email = c.Query("email")
	param.WithPaymentMethod = c.Query("withPaymentMethod")

	groups, err := service.FindGroupByEmail(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	group := &service.Group{}

	if err := c.ShouldBindJSON(&group); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	param := service.UpdateGroupParam{
		Id:    id,
		Group: group,
	}

	group, err := service.UpdateGroup(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, group)
}
