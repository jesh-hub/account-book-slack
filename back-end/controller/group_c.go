package controller

import (
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddGroup
// @Summary Add group
// @Description 그룹 등록
// @Tags group
// @Accept json
// @Produce json
// @Param payment body service.Group true "Group"
// @Success 200 {object} service.Group
// @Router /v1/group [post]
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

// FindGroupById
// @Summary Find group by group id
// @Description groupId로 그룹 조회
// @Tags group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} service.Group
// @Router /v1/group/{id} [get]
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

// FindGroupByEmail
// @Summary Find group by email
// @Description email로 그룹 조회
// @Tags group
// @Accept json
// @Produce json
// @Param email query string true "email"
// @Success 200 {array} service.Group
// @Router /v1/group [get]
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

// UpdateGroup
// @Summary Update group
// @Description 그룹 수정
// @Tags group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param group body service.Group true "Group"
// @Success 200 {object} service.Group
// @Router /v1/group/{id} [put]
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
