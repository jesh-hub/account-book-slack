package controller

import (
	"abs/model"
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddGroup
// @Summary Add group
// @Description 그룹 등록
// @Tags group
// @Accept json
// @Produce json
// @Param payment body model.Group true "Group"
// @Success 200 {object} model.Group
// @Router /v1/group [post]
func AddGroup(c *gin.Context) {
	group := &model.Group{}
	if err := c.ShouldBindJSON(group); err != nil {
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

// FindGroup
// @Summary Find group by paramter
// @Description email로 그룹 조회
// @Tags group
// @Accept json
// @Produce json
// @Param email query string true "email"
// @Param withPaymentMethod query bool false "결제수단 함께 조회하는지 여부"
// @Success 200 {array} model.Group
// @Router /v1/group [get]
func FindGroup(c *gin.Context) {
	param := service.NewFindGroupParam()
	param.Email = c.Query("email")

	if v, ok := c.GetQuery("withPaymentMethod"); ok {
		param.WithPaymentMethod, _ = strconv.ParseBool(v)
	}

	groups, err := service.FindGroupByEmail(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

// FindGroupById
// @Summary Find group by group id
// @Description groupId로 그룹 조회
// @Tags group
// @Accept json
// @Produce json
// @Param groupId path string true "Group ID"
// @Param withPaymentMethod query bool false "결제수단 함께 조회하는지 여부"
// @Success 200 {object} model.Group
// @Router /v1/group/{groupId} [get]
func FindGroupById(c *gin.Context) {
	param := service.NewFindGroupParam()
	param.Id = c.Param("groupId")

	if v, ok := c.GetQuery("withPaymentMethod"); ok {
		param.WithPaymentMethod, _ = strconv.ParseBool(v)
	}

	group, err := service.FindGroupById(param)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

// UpdateGroup
// @Summary Update group
// @Description 그룹 수정
// @Tags group
// @Accept json
// @Produce json
// @Param groupId path string true "Group ID"
// @Param group body model.GroupUpdate true "GroupUpdate"
// @Success 200 {object} model.Group
// @Router /v1/group/{groupId} [put]
func UpdateGroup(c *gin.Context) {
	id := c.Param("groupId")
	groupUpdate := &model.GroupUpdate{}

	if err := c.ShouldBindJSON(groupUpdate); err != nil {
		util.ErrorHandler(c, 400, err)
		return
	}

	group, err := service.UpdateGroup(id, groupUpdate)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

// DeleteGroup
// @Summary Delete group
// @Description 그룹 삭제
// @Tags group
// @Accept json
// @Param groupId path string true "Group ID"
// @Success 200
// @Router /v1/group/{groupId} [delete]
func DeleteGroup(c *gin.Context) {
	id := c.Param("groupId")

	err := service.DeleteGroup(id)
	if err != nil {
		util.ErrorHandler(c, 500, err)
		return
	}
	c.Status(http.StatusOK)
}
