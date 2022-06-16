package mongo

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
)

var dbuser DBUser

type SecretsStruct struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title" json:"title"`
	Text  string             `bson:"text" json:"text"`
}

func init() {
	godotenv.Load()
	dbuser = DBUser{
		Auth: os.Getenv("DB_AUTH"),
		Host: os.Getenv("DB_HOST"),
		DB:   "abs",
	}
}

func TestConnectDB(t *testing.T) {
	dbuser.ConnectDB()
}

func TestInsert(t *testing.T) {
	client, ctx, cancel := dbuser.ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbuser.DB).Collection("user")

	obj := SecretsStruct{
		Title: "choshsh",
		Text:  "12312321321",
	}
	result, err := coll.InsertOne(ctx, obj)
	errorHandler(err)

	fmt.Printf("%+v\n", result)
	assert.NotEmpty(t, result)
}

func TestRead(t *testing.T) {
	client, ctx, cancel := dbuser.ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbuser.DB).Collection("user")

	var obj2 SecretsStruct
	err := coll.FindOne(ctx, bson.D{{"title", "choshsh"}}).Decode(&obj2)
	errorHandler(err)

	fmt.Printf("%+v\n", obj2)
	assert.NotEmpty(t, obj2)
}
