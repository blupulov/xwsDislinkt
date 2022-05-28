package handler

import (
	"context"

	pb "github.com/blupulov/xwsDislinkt/common/proto/services/following-service"
	"github.com/blupulov/xwsDislinkt/following-service/service"
)

type FollowingHandler struct {
	fs *service.FollowService
	pb.UnimplementedFollowingServiceServer
}

func NewFollowingHandler(followService *service.FollowService) *FollowingHandler {
	return &FollowingHandler{
		fs: followService,
	}
}

func (fh *FollowingHandler) GetAllUserFollowers(ctx context.Context, r *pb.GetAllUserFollowersRequest) (*pb.GetAllUserFollowersResponse, error) {
	var response pb.GetAllUserFollowersResponse

	users, err := fh.fs.GetAllUserFollowers(r.UserId)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		response.UsersIds = append(response.UsersIds, user.UserId)
	}

	return &response, nil
}

func (fh *FollowingHandler) GetAllFollowingUsers(ctx context.Context, r *pb.GetAllFollowingUsersRequest) (*pb.GetAllFollowingUsersResponse, error) {
	var response pb.GetAllFollowingUsersResponse

	users, err := fh.fs.GetAllFollowingUsers(r.UserId)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		response.UsersIds = append(response.UsersIds, user.UserId)
	}

	return &response, nil
}
