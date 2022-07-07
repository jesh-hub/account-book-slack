package abs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name" binding:"required"`
	Users            []string `json:"users" bson:"users" binding:"required"`
	RegUserId        string   `json:"regUserId" bson:"regUserId"`
	ModUserId        string   `json:"modUserId" bson:"modUserId"`
}

func newGroup(name string, users []string, regUserId string) Group {
	return Group{
		Name:      name,
		Users:     users,
		RegUserId: regUserId,
		ModUserId: "",
	}
}

func FindGroupById(c *gin.Context) {
	groupId := c.Param("id")
	if len(groupId) == 0 {
		errorHandler(c, 400, errors.New("invalid parameter"))
		return
	}

	groupColl := mgm.Coll(&Group{})
	group := &Group{}
	_ = groupColl.FindByID(groupId, group)

	c.JSON(http.StatusOK, group)
}

func FindGroupByEmail(c *gin.Context) {
	email := c.Query("email")
	if len(email) == 0 {
		errorHandler(c, 400, errors.New("invalid parameter"))
		return
	}

	groupColl := mgm.Coll(&Group{})
	groups := &[]Group{}
	_ = groupColl.SimpleFind(groups, bson.M{"users": email})

	c.JSON(http.StatusOK, groups)
}

func AddGroup(c *gin.Context) {
	group := &Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		errorHandler(c, 400, err)
		return
	}

	groupColl := mgm.Coll(&Group{})
	err := groupColl.Create(group)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, *group)
}

func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errorHandler(c, 400, errors.New("parameter is invalid"))
		return
	}

	groupColl := mgm.Coll(&Group{})
	group := &Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		errorHandler(c, 400, err)
		return
	}

	err := groupColl.Update(group)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, group)
}
