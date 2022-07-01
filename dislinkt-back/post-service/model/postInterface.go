package model

import (
	"github.com/blupulov/xwsDislinkt/post-service/dto"
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
	DeleteCommentById(commentId, postId string) error
	CreatePostFromJob(dto *dto.ShareJobDto) error
}
