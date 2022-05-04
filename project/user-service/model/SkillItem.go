package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SkillItem struct {
	Id               primitive.ObjectID `bson: "_id" json: "_id"`
	SkillName        string             `bson:"skillName" json:"skillName"`
	SkillDescription string             `bson:"skillDescription" json:"skillDescription"`
}
