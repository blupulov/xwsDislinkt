package handler

import (
	"context"

	pb "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	"github.com/blupulov/xwsDislinkt/post-service/service"
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
		post := mapPost(p)
		response.Post = append(response.Post, post)
	}

	return response, nil
}
