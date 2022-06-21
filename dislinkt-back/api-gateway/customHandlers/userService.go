package customHandlers

import (
	"context"

	"github.com/blupulov/xwsDislinkt/api-gateway/grpcClients"
	"github.com/blupulov/xwsDislinkt/api-gateway/model"
	us "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
)

func (ch *CustomHandler) promoteUserToCompanyOwner(userId string) error {
	userClient := grpcClients.NewUserClient(ch.userClientAddress)

	_, err := userClient.PromoteUserToCompanyOwner(context.TODO(), &us.PromoteUserRequest{UserId: userId})

	return err
}

func (ch *CustomHandler) getUserById(userId string) (*model.User, error) {
	userClient := grpcClients.NewUserClient(ch.userClientAddress)

	res, err := userClient.GetById(context.TODO(), &us.GetByIdRequest{Id: userId})
	if err != nil {
		return nil, err
	}

	return mapUser(res), nil
}

func mapUser(res *us.GetByIdResponse) *model.User {
	return &model.User{
		Id:           res.User.Id,
		Username:     res.User.Username,
		FirstName:    res.User.FirstName,
		LastName:     res.User.LastName,
		Email:        res.User.Email,
		ProfileImage: res.User.Image,
	}
}
