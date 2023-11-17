package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description"`
	Amount      float32            `bson:"amount"`
}
