package api

import (
	"abs/model"
	"abs/service"
	"abs/util"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

// AddGroup
// @Summary Add group
// @Description 그룹 등록
// @Tags group
// @Accept json
// @Produce json
// @Param groupAdd body model.GroupAdd true "GroupAdd"
// @Success 200 {object} model.Group
// @Failure 400 {object} util.AppError
// @Failure 500 {object} util.AppError
// @Router /v1/group [post]
func AddGroup(c *gin.Context) {
	groupAdd := &model.GroupAdd{}
	if err := c.ShouldBindJSON(groupAdd); err != nil {
		c.JSON(http.StatusBadRequest, util.NewAppError(err))
		return
	}

	group, err := service.AddGroup(groupAdd)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
		return
	}
	c.JSON(http.StatusOK, group)
}

// FindGroup
// @Summary Find group by parameter
// @Tags group
// @Accept json
// @Produce json
// @Param email query string true "그룹 조회 by email"
// @Param withPaymentMethod query bool false "결제수단 함께 조회하는지 여부"
// @Success 200 {array} model.Group
// @Failure 400 {object} util.AppError
// @Failure 500 {object} util.AppError
// @Router /v1/group [get]
func FindGroup(c *gin.Context) {
	param := service.NewFindGroupParam()
	param.Email = c.Query("email")

	if v, ok := c.GetQuery("withPaymentMethod"); ok {
		var err error
		param.WithPaymentMethod, err = strconv.ParseBool(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.NewAppError(err))
			return
		}
	}

	groups, err := service.FindGroupByEmail(param)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
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

	// 결제수단 함께 조회 여부 파라미터 처리
	// type 변환 : string -> bool
	if v, ok := c.GetQuery("withPaymentMethod"); ok {
		var err error
		param.WithPaymentMethod, err = strconv.ParseBool(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.NewAppError(err))
			return
		}
	}

	group, err := service.FindGroupById(param)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
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
// @Failure 400 {object} util.AppError
// @Failure 500 {object} util.AppError
// @Router /v1/group/{groupId} [put]
func UpdateGroup(c *gin.Context) {
	id := c.Param("groupId")
	groupUpdate := &model.GroupUpdate{}

	if err := c.ShouldBindJSON(groupUpdate); err != nil {
		c.JSON(http.StatusBadRequest, util.NewAppError(err))
		return
	}

	group, err := service.UpdateGroup(id, groupUpdate)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
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
// @Failure 500 {object} util.AppError
// @Router /v1/group/{groupId} [delete]
func DeleteGroup(c *gin.Context) {
	id := c.Param("groupId")

	err := service.DeleteGroup(id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, util.NewAppError(err))
		return
	}
	c.Status(http.StatusOK)
}
