package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lastName" bson:"lastName"`
	Picture          string `json:"picture" bson:"picture"`
}
