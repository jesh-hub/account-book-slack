package abs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type DbUser struct {
	Host string
	Auth string
	Name string
}

type FindOptions struct {
	Filter bson.D
	Opts   bson.D
}

var dbUser DbUser

// connect
func (du *DbUser) Connect() (*mongo.Client, context.Context, context.CancelFunc) {
	uri := "mongodb+srv://" + du.Auth + "@" + du.Host + "/?retryWrites=true&w=majority"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	// Timeout 설정을 위한 Context생성
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	errorHandler(err)

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	return client, ctx, cancel
}

// findOne
func findOne(collection string, findOptions FindOptions, result interface{}) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbUser.Name).Collection(collection)

	var opts *options.FindOneOptions
	if findOptions.Opts != nil {
		opts = options.FindOne().SetProjection(findOptions.Opts)
	} else {
		opts = nil
	}
	err := coll.FindOne(ctx, findOptions.Filter, opts).Decode(result)
	errorHandler(err)
}

// findMany
func findMany(collection string, findOptions FindOptions, result interface{}) {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbUser.Name).Collection(collection)

	var opts *options.FindOptions
	if findOptions.Opts != nil {
		opts = options.Find().SetProjection(findOptions.Opts)
	} else {
		opts = nil
	}
	cursor, err := coll.Find(ctx, findOptions.Filter, opts)
	errorHandler(err)

	err = cursor.All(context.TODO(), result)
	errorHandler(err)
}

// insertOne
func insertOne(collection string, data interface{}) interface{} {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbUser.Name).Collection(collection)
	result, err := coll.InsertOne(ctx, data)
	errorHandler(err)

	return result.InsertedID
}

// insertMany
func insertMany(collection string, data ...interface{}) []interface{} {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbUser.Name).Collection(collection)
	result, err := coll.InsertMany(ctx, data)
	errorHandler(err)

	return result.InsertedIDs
}

// updateOne
func updateOne(collection string, id string, data interface{}) interface{} {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbUser.Name).Collection(collection)
	result, err := coll.UpdateByID(ctx, id, data)
	errorHandler(err)

	return result.UpsertedID
}

// deleteOne
func deleteOne(collection string, id string) int64 {
	client, ctx, cancel := dbUser.Connect()
	defer client.Disconnect(ctx)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	coll := client.Database(dbUser.Name).Collection(collection)
	result, err := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	errorHandler(err)

	return result.DeletedCount
}
