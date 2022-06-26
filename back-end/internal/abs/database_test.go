package abs

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

type SecretsStruct struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title" json:"title"`
	Text  string             `bson:"text" json:"text"`
}

func init() {
	ConnectDB()
}

func TestConnect(t *testing.T) {
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
	assert.NotEmpty(t, result.Id)
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
	findMany(collection, findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.Greater(t, len(result), 0)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
