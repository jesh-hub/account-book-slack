package abs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var groupCollection = GetCollection(DB, "group")

type Group struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Users     []string           `json:"users" bson:"users" binding:"required"`
	RegUserId string             `json:"regUserId" bson:"regUserId"`
	RegDate   primitive.DateTime `json:"regDate" bson:"regDate"`
	ModUserId string             `json:"modUserId" bson:"modUserId"`
	ModDate   primitive.DateTime `json:"modDate" bson:"modDate"`
}

func newGroup(name string, users []string, regUserId string) Group {
	return Group{
		Name:      name,
		Users:     users,
		RegUserId: regUserId,
		RegDate:   primitive.NewDateTimeFromTime(time.Now()),
		ModUserId: "",
		ModDate:   primitive.NewDateTimeFromTime(time.Now()),
	}
}

func FindGroupById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errorHandler(c, 400, errors.New("invalid parameter"))
		return
	}
	groupId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	var groups []Group
	findOptions := FindOptions{
		Filter: bson.M{"_id": groupId},
	}
	if err = findMany(groupCollection, findOptions, &groups); err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, groups)
}

func FindGroupByEmail(c *gin.Context) {
	email := c.Query("email")
	if len(email) == 0 {
		errorHandler(c, 400, errors.New("invalid parameter"))
		return
	}

	var groups []Group
	findOptions := FindOptions{
		Filter: bson.M{"users": email},
	}
	if err := findMany(groupCollection, findOptions, &groups); err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, groups)
}

func AddGroup(c *gin.Context) {
	var group Group
	if err := c.ShouldBindJSON(&group); err != nil {
		errorHandler(c, 400, err)
		return
	}

	result, err := insertOne(groupCollection, group)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	groupId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	var group Group
	if err := c.ShouldBindJSON(&group); err != nil {
		errorHandler(c, 400, err)
		return
	}

	group.ModDate = primitive.NewDateTimeFromTime(time.Now())
	result, err := replaceOne(groupCollection, groupId, group)
	if err != nil {
		errorHandler(c, 400, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
