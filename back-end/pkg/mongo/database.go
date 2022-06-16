package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type DBUser struct {
	Host string
	Auth string
	DB   string
}

func (du *DBUser) ConnectDB() (*mongo.Client, context.Context, context.CancelFunc) {
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

	fmt.Println("Successfully connected and p	inged.")
	return client, ctx, cancel
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
