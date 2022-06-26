package abs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"
)

func init() {
	ConnectDB()
}

func TestInsert(t *testing.T) {
	user := newUser("choshsh@test.com")

	collection := GetCollection(DB, "user")
	result, err := insertOne(collection, user)
	errorHandlerInternal(err)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)
}

func TestFindOne(t *testing.T) {
	email := "choshsh@test.com"

	var result User

	findOptions := FindOptions{
		Filter: bson.D{{"_id", email}},
	}
	collection := GetCollection(DB, "user")
	findOne(collection, findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result.Email)
}

func TestFindOneByOptions(t *testing.T) {
	email := "choshsh@test.com"
	var result User

	collection := GetCollection(DB, "user")
	findOptions := FindOptions{
		Filter: bson.D{{"_id", email}},
		Opts:   bson.D{{"regDate", 0}},
	}
	findOne(collection, findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)
}

func TestFindMany(t *testing.T) {
	var result []User

	collection := GetCollection(DB, "user")
	findOptions := FindOptions{
		Filter: bson.M{
			"regDate": bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now().Add(-9 * time.Hour))},
		},
		Opts: nil,
	}
	err := findMany(collection, findOptions, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prettyPrint(result))
	assert.Greater(t, len(result), 0)
}
