package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EducationItem struct {
	Id               primitive.ObjectID `bson:"_id" json:"_id"`
	SchoolName       string             `bson:"schoolName" json:"schoolName"`
	SchoolType       string             `bson:"schoolType" json:"schoolType"` //greska u modelu
	DurationInMonths int                `bson:"durationInMonths" json:"durationInMonths"`
}
