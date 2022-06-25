package abs

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
)

type SecretsStruct struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title" json:"title"`
	Text  string             `bson:"text" json:"text"`
}

func init() {
	godotenv.Load()
	dbUser = DbUser{
		Host: os.Getenv("DB_HOST"),
		Auth: os.Getenv("DB_AUTH"),
		Name: os.Getenv("DB_NAME"),
	}
}

func TestConnect(t *testing.T) {
	dbUser.Connect()
}

func TestInsert(t *testing.T) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	obj := SecretsStruct{
		Title: "choshsh123123123",
		Text:  "1231232132112312312",
	}

	id := insertOne("user", obj)

	fmt.Println(prettyPrint(id))
	assert.NotEmpty(t, id)
}

func TestFindOne(t *testing.T) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	var result SecretsStruct

	findOptions := FindOptions{
		Filter: bson.D{{"title", "choshsh123123123"}},
	}
	findOne("user", findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result.Text)
}

func TestFindOneByOptions(t *testing.T) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	var result SecretsStruct

	findOptions := FindOptions{
		Filter: bson.D{{"title", "choshsh123123123"}},
		Opts:   bson.D{{"text", 0}},
	}
	findOne("user", findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.NotEmpty(t, result)
}

func TestFindMany(t *testing.T) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	var result []SecretsStruct

	findOptions := FindOptions{
		Filter: bson.D{{"title", "choshsh123123123"}},
		Opts:   nil,
	}
	findMany("user", findOptions, &result)

	fmt.Println(prettyPrint(result))
	assert.Greater(t, len(result), 0)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
