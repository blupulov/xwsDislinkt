package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	CompanyName string             `bson:"companyName" json:"companyName"`
	OwnerId     string             `bson:"ownerId" json:"ownerId"`
	IsAccepted  bool               `bson:"isAccepted" json:"isAccepted"`
	Description string             `bson:"description" json:"description"`
	Info        Info               `bson:"info" json:"info"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	OpenJobs    []Job              `bson:"jobs" json:"openJobs"`
}

type Job struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	JobName        string             `bson:"jobName" json:"jobName"`
	JobDescription string             `bson:"jobDescription" json:"jobDescription"`
}

type Comment struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	CommentOwnerId string             `bson:"commentOwnerId" json:"commentOwnerId"`
	Grade          int                `bson:"grade" json:"grade"`
	Comment        string             `bson:"comment" json:"comment"`
}

type Info struct {
	Street      string `bson:"street" json:"street"`
	City        string `bson:"city" json:"city"`
	Country     string `bson:"country" json:"country"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Email       string `bson:"email" json:"email"`
}
