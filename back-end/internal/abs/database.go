package abs

import (
	"abs/pkg/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type FindOptions struct {
	Filter interface{}
	Opts   interface{}
}

var DB = ConnectDB()
var dbName string

func ConnectDB() *mongo.Client {
	host := util.GodotEnv("DB_HOST")
	auth := util.GodotEnv("DB_AUTH")
	dbName = util.GodotEnv("DB_NAME")

	uri := "mongodb+srv://" + auth + "@" + host + "/?retryWrites=true&w=majority"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	// Timeout 설정을 위한 Context생성
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	errorHandlerInternal(err)

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

// findOne
func findOne(collection *mongo.Collection, findOptions FindOptions, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	var opts *options.FindOneOptions
	if findOptions.Opts != nil {
		opts = options.FindOne().SetProjection(findOptions.Opts)
	} else {
		opts = nil
	}
	err := collection.FindOne(ctx, findOptions.Filter, opts).Decode(result)
	return err
}

// findMany
func findMany(collection *mongo.Collection, findOptions FindOptions, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	var opts *options.FindOptions
	if findOptions.Opts != nil {
		opts = options.Find().SetProjection(findOptions.Opts)
	} else {
		opts = nil
	}
	cursor, err := collection.Find(ctx, findOptions.Filter, opts)
	if err != nil {
		return err
	}

	err = cursor.All(context.TODO(), result)
	if err != nil {
		return err
	}
	return nil
}

// insertOne
func insertOne(collection *mongo.Collection, data interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	result, err := collection.InsertOne(ctx, data)
	return result.InsertedID, err
}

// insertMany
func insertMany(collection *mongo.Collection, data ...interface{}) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	result, err := collection.InsertMany(ctx, data)
	return result.InsertedIDs, err
}

// updateOne
func updateOne(collection *mongo.Collection, id string, data interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	update := bson.M{"$set": data}
	result, err := collection.UpdateByID(ctx, id, update)

	return result, err
}

// deleteOne
func deleteOne(collection *mongo.Collection, id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	return result.DeletedCount, err
}
