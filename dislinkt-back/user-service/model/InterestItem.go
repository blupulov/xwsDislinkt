package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InterestItem struct {
	Id                  primitive.ObjectID `bson:"_id" json:"_id"`
	InterestName        string             `bson:"interestName" json:"interestName"`
	InterestDescription string             `bson:"interestDescription" json:"interestDescription"`
}
