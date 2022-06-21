package customHandlers

import (
	"context"

	"github.com/blupulov/xwsDislinkt/api-gateway/grpcClients"
	fs "github.com/blupulov/xwsDislinkt/common/proto/services/following-service"
)

func (ch *CustomHandler) getUserFollowers(userId string) (*[]string, error) {
	followingClient := grpcClients.NewFollowingClient(ch.followingClientAddress)

	res, err := followingClient.GetAllUserFollowers(context.TODO(), &fs.GetAllUserFollowersRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	var followers []string
	for _, userID := range res.UsersIds {
		followers = append(followers, userID)
	}

	return &followers, nil
}

func (ch *CustomHandler) getFollowingUsers(userId string) (*[]string, error) {
	followingClient := grpcClients.NewFollowingClient(ch.followingClientAddress)

	res, err := followingClient.GetAllFollowingUsers(context.TODO(), &fs.GetAllFollowingUsersRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	var following []string
	for _, userID := range res.UsersIds {
		following = append(following, userID)
	}

	return &following, nil
}
