package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 		primitive.ObjectID `bson:"_id,omitempty" json:"id,omimtempty"`
	Email 	string             `bson:"email" json:"email"`
	password 	string             `bson:"password" json:"-"`
	Todos 	[]primitive.ObjectID `bson:"todos" json:"todos"`
}