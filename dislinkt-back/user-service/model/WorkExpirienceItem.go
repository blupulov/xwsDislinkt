package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type WorkExperienceItem struct {
	Id                   primitive.ObjectID `bson:"_id" json:"_id"`
	CompanyName          string             `bson:"companyName" json:"companyName"`
	WorkDurationInMonths int                `bson:"workDurationInMonths" json:"workDurationInMonths"`
	EmploymentField      string             `bson:"employmentField" json:"employmentField"`
}
