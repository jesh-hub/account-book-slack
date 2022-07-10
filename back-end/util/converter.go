package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjectId(id string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(id)
	ErrorHandlerInternal(err)
	return objectId
}
