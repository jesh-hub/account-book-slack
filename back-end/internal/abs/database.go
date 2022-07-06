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
}
