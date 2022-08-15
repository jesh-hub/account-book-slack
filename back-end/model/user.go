package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Picture          string `json:"picture"`
}
