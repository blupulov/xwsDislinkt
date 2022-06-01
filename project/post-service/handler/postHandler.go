package handler

import (
	"context"

	pb "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	"github.com/blupulov/xwsDislinkt/post-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	ps *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		ps: postService,
	}
}

func (ph *PostHandler) GetAll(ctx context.Context, r *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := ph.ps.GetAll()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		//Post is array
		Post: []*pb.Post{},
	}

	for _, p := range posts {
		post := mapPbFromPost(p)
		response.Post = append(response.Post, post)
	}

	return response, nil
}

func (ph *PostHandler) Insert(ctx context.Context, r *pb.InsertRequest) (*pb.InsertResponse, error) {
	var response pb.InsertResponse

	newPost := mapPostFromPb(r.Post)
	err := ph.ps.Insert(newPost)
	if err != nil {
		response.Status = "Not created"
	} else {
		response.Status = "Created"
	}

	return &response, err
}

func (ph *PostHandler) Like(ctx context.Context, r *pb.LikeRequest) (res *pb.LikeResponse, err error) {
	res = &pb.LikeResponse{}
	postId, err := primitive.ObjectIDFromHex(r.PostId)
	if err != nil {
		res.Status = "Bad post ID"
		return
	}

	err = ph.ps.Like(r.UserId, postId)
	if err != nil {
		res.Status = "Post did not liked"
	} else {
		res.Status = "Post liked"
	}

	return
}

func (ph *PostHandler) GetAllUserPosts(ctx context.Context, r *pb.GetAllUserPostsRequest) (*pb.GetAllUserPostsResponse, error) {
	var response pb.GetAllUserPostsResponse

	posts, err := ph.ps.GetAllByOwnerId(r.UserId)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		response.Posts = append(response.Posts, mapPbFromPost(post))
	}

	return &response, nil
}

func (ph *PostHandler) CommentPost(ctx context.Context, r *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	var response pb.CommentPostResponse

	postId, err := primitive.ObjectIDFromHex(r.PostId)
	if err != nil {
		return nil, err
	}

	comment := mapPostCommentFromPb(r.Comment)
	err = ph.ps.AddComment(comment, postId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ph *PostHandler) DeletePost(ctx context.Context, r *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	var response pb.DeletePostResponse

	postId, err := primitive.ObjectIDFromHex(r.PostId)
	if err != nil {
		return nil, err
	}

	err = ph.ps.DeleteById(postId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// func (ph *PostHandler) Dislike(ctx context.Context, r *pb.DislikeRequest) (res *pb.DislikeResponse, err error) {
// 	res = &pb.DislikeResponse{}
// 	postId, err := primitive.ObjectIDFromHex(r.PostId)
// 	if err != nil {
// 		res.Status = "Bad post ID"
// 		return
// 	}

// 	err = ph.ps.Dislike(r.userId, r.postId)
// 	if err != nil {
// 		res.Status = "Didnt dislike post"
// 	} else {
// 		res.Status = "Post disliked"
// 	}

// 	return
// }
