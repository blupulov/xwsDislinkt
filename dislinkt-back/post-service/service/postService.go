package service

import (
	"github.com/blupulov/xwsDislinkt/post-service/dto"
	"github.com/blupulov/xwsDislinkt/post-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Only for call methods of PostInterface
type PostService struct {
	pi model.PostInterface
}

func NewPostService(pi model.PostInterface) *PostService {
	return &PostService{
		pi: pi,
	}
}

func (ps *PostService) GetById(id primitive.ObjectID) (*model.Post, error) {
	return ps.pi.GetById(id)
}

func (ps *PostService) GetAll() ([]*model.Post, error) {
	return ps.pi.GetAll()
}

func (ps *PostService) GetAllByOwnerId(ownerId string) ([]*model.Post, error) {
	return ps.pi.GetAllByOwnerId(ownerId)
}

func (ps *PostService) Insert(post *model.Post) error {
	return ps.pi.Insert(post)
}

func (ps *PostService) DeleteById(id primitive.ObjectID) error {
	return ps.pi.DeleteById(id)
}

func (ps *PostService) Like(userId string, postId primitive.ObjectID) error {
	return ps.pi.Like(userId, postId)
}

func (ps *PostService) Dislike(userId string, postId primitive.ObjectID) error {
	return ps.pi.Dislike(userId, postId)
}

func (ps *PostService) AddComment(comment *model.Comment, postId primitive.ObjectID) error {
	return ps.pi.AddComment(comment, postId)
}

func (ps *PostService) DeleteCommentById(commentId, postId string) error {
	return ps.pi.DeleteCommentById(commentId, postId)
}

func (ps *PostService) CreatePostFromJob(dto *dto.ShareJobDto) error {
	return ps.pi.CreatePostFromJob(dto)
}
