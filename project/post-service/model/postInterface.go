package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostInterface interface {
	GetById(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	GetAllByOwnerId(ownerId string) ([]*Post, error)
	Insert(post *Post) error
	DeleteById(id primitive.ObjectID) error
	Like(userId string, postId primitive.ObjectID) error
	Dislike(userId string, postId primitive.ObjectID) error
	AddComment(comment *Comment, postId primitive.ObjectID) error
	//delete comment by Id, or (ownerId and creation time)??
}
