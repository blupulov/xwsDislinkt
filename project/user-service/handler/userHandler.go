package handler

import (
	"context"

	ub "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	ub.UnimplementedUserServiceServer
	us *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		us: userService,
	}
}

func (uh *UserHandler) GetAll(ctx context.Context, r *ub.GetAllRequest) (*ub.GetAllResponse, error) {
	users, err := uh.us.GetAll()
	if err != nil {
		return nil, err
	}

	response := &ub.GetAllResponse{
		Users: []*ub.User{},
	}

	for _, u := range users {
		user := mapUbFromUser(u)
		response.Users = append(response.Users, user)
	}

	return response, nil
}

func (uh *UserHandler) Register(ctx context.Context, r *ub.RegisterRequest) (*ub.RegisterResponse, error) {
	var response ub.RegisterResponse

	newUser := mapUserFromUb(r.User)
	err := uh.us.Register(newUser)

	if err != nil {
		response.Status = "Not created"
	} else {
		response.Status = "Created"
	}

	return &response, err
}
