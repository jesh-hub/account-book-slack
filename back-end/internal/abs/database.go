package abs

import (
	"abs/pkg/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	host := util.GodotEnv("DB_HOST")
	auth := util.GodotEnv("DB_AUTH")
	dbName := util.GodotEnv("DB_NAME")

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI("mongodb+srv://"+auth+"@"+host+"/?retryWrites=true&w=majority"))
	errorHandlerInternal(err)
	//
	//uri := "mongodb+srv://" + auth + "@" + host + "/?retryWrites=true&w=majority"
	//serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	//
	//// Timeout 설정을 위한 Context생성
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//client, err := mongo.Connect(ctx, clientOptions)
	//errorHandlerInternal(err)
	//
	//// Ping the primary
	//if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully connected and pinged.")
	//return client
}
