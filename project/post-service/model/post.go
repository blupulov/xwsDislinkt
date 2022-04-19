package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	PostComment   string             `bson:"postComment" json:"postComment"`
	PostImage     string             `bson:"postImage" json:"postImage"`
	PostOwnerId   string             `bson:"postOwnerId" json:"postOwnerId"`
	PostDate      time.Time          `bson:"postDate" json:"postDate"`
	LikeNumber    int                `bson:"postLikeNumber" json:"postLikeNumber"`
	DislikeNumber int                `bson:"postDislikeNumber" json:"postDislikeNumber"`
	FansIds       []string           `bson:"fansIds" json:"fansIds"`
	HatersIds     []string           `bson:"hatersIds" json:"hatersIds"`
	Comments      []Comment          `bson:"comments" json:"comments"`
}

//Verovatno ce i ovde morati da se doda Id
type Comment struct {
	CommentOwnerId string    `bson:"commentOwnerId" json:"commentOwnerId"`
	CommentContent string    `bson:"commentContent" json:"commentContent"`
	CreationDate   time.Time `bson:"creationDate" json:"creationDate"`
}
