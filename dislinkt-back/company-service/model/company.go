package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Id          primitive.ObjectID `bson:"_id"`
	CompanyName string             `bson:"companyName"`
	OwnerId     string             `bson:"ownerId"`
	IsAccepted  bool               `bson:"isAccepted"`
	Description string             `bson:"description"`
	Info        Info               `bson:"info"`
	Comments    []Comment          `bson:"comments"`
	OpenJobs    []Job              `bson:"jobs"`
}

type Job struct {
	Id             primitive.ObjectID `bson:"_id"`
	JobName        string             `bson:"jobName"`
	JobDescription string             `bson:"jobDescription"`
}

type Comment struct {
	Id             primitive.ObjectID `bson:"_id"`
	CommentOwnerId string             `bson:"commentOwnerId"`
	Grade          int                `bson:"grade"`
	Comment        string             `bson:"comment"`
}

type Info struct {
	Street      string `bson:"street"`
	City        string `bson:"city"`
	Country     string `bson:"country"`
	PhoneNumber string `bson:"phoneNumber"`
	Email       string `bson:"email"`
}
