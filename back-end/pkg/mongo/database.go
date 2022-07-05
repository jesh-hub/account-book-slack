package mongo

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBUser struct {
	Host string
	Auth string
	DB   string
}

//func (du *DBUser) ConnectDB() (*mongo.Client, context.Context, context.CancelFunc) {
//	uri := "mongodb+srv://" + du.Auth + "@" + du.Host + "/?retryWrites=true&w=majority"
//	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
//	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
//
//	// Timeout 설정을 위한 Context생성
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	client, err := mongo.Connect(ctx, clientOptions)
//	errorHandler(err)
//
//	// Ping the primary
//	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
//		panic(err)
//	}
//
//	fmt.Println("Successfully connected and p	inged.")
//	return client, ctx, cancel
//}
func (du *DBUser) ConnectDB() (*mongo.Client, context.Context, context.CancelFunc) {
	err := mgm.SetDefaultConfig(nil, du.DB, options.Client().ApplyURI("mongodb://"+du.Auth+"@"+du.Host))
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
