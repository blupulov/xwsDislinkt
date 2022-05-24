package handler

import (
	pb "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUbFromUser(user *model.User) *pb.User {
	ubUser := &pb.User{
		Id:             user.Id.Hex(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		BirthDate:      timestamppb.New(user.BirthDate),
		Email:          user.Email,
		Username:       user.Username,
		Password:       user.Password,
		Education:      []*pb.EducationItem{}, //za potrebe testiranja,sve dole
		WorkExpirience: []*pb.WorkExpirienceItem{},
		Skills:         []*pb.SkillItem{},
		Interests:      []*pb.Interests{},
		BlockedUsers:   []string{},
	}

	return ubUser
}

func mapUserForRegistration(newUser *pb.UserRegistrationModel) *model.User {
	return &model.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		//BirthDate:      newUser.GetBirthdate().AsTime(),
		Email:          newUser.Email,
		Username:       newUser.Username,
		Password:       newUser.Password,
		Education:      make([]model.EducationItem, 0), //za potrebe testiranja
		WorkExperience: make([]model.WorkExperienceItem, 0),
		Skills:         make([]model.SkillItem, 0),
		Interests:      make([]model.InterestItem, 0),
		BlockedUsers:   make([]string, 0),
	}
}
