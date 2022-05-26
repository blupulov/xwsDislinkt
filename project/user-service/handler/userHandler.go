package handler

import (
	"context"

	pb "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	us *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		us: userService,
	}
}

func (uh *UserHandler) GetAll(ctx context.Context, r *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := uh.us.GetAll()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}

	for _, u := range users {
		user := mapPbUserFromModel(u)
		response.Users = append(response.Users, user)
	}

	return response, nil
}

func (uh *UserHandler) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	newUser := mapUserForRegistration(r.NewUser)
	err := uh.us.Register(newUser)

	if err != nil {
		response.Status = "Not created"
	} else {
		response.Status = "Created"
	}

	return &response, err
}

func (uh *UserHandler) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := uh.us.Login(r.Credentials.Password, r.Credentials.Username)
	if err != nil {
		return nil, err
	}

	var response pb.LoginResponse
	response.Role = token.Role
	response.UserId = token.Id
	response.Token = token.Token

	return &response, nil
}

func (uh *UserHandler) GetById(ctx context.Context, r *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	var response pb.GetByIdResponse

	userId, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, err
	}

	user, err := uh.us.GetById(userId)
	if err != nil {
		return nil, err
	}

	response.User = mapPbUserFromModel(user)

	return &response, nil
}

func (uh *UserHandler) AddSkill(ctx context.Context, r *pb.AddSkillRequest) (*pb.AddSkillResponse, error) {
	var response pb.AddSkillResponse

	userId, err := primitive.ObjectIDFromHex(r.UserId)
	if err != nil {
		return nil, err
	}

	newSkill := mapUserForAddingSkill(r.NewSkill)
	err = uh.us.AddSkill(newSkill, userId)

	if err != nil {
		response.Status = "Not created"
	} else {
		response.Status = "Created"
	}

	return &response, err
}

func (uh *UserHandler) GetManyUsersById(ctx context.Context, r *pb.GetManyUsersByIdRequest) (*pb.GetManyUsersByIdResponse, error) {
	var response pb.GetManyUsersByIdResponse

	users, err := uh.us.GetManyUsersById(r.UsersIds.UsersIds)
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		response.Users = append(response.Users, mapPbUserFromModel(&user))
	}

	return &response, nil
}
