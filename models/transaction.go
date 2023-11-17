package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Description string             `bson:"description" json:"description"`
	Amount      float32            `bson:"amount" json:"amount"`
}
