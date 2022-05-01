package handler

import (
	ub "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUbFromUser(user *model.User) *ub.User {
	ubUser := &ub.User{
		Id:             user.Id.Hex(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		BirthDate:      timestamppb.New(user.BirthDate),
		Email:          user.Email,
		Username:       user.Username,
		Password:       user.Password,
		Education:      []*ub.EducationItem{}, //za potrebe testiranja,sve dole
		WorkExpirience: []*ub.WorkExpirienceItem{},
		Skills:         []*ub.SkillItem{},
		Interests:      []*ub.Interests{},
		BlockedUsers:   []string{},
	}

	return ubUser
}

func mapUserFromUb(ubUser *ub.User) *model.User {
	return &model.User{
		FirstName:      ubUser.FirstName,
		LastName:       ubUser.LastName,
		BirthDate:      ubUser.GetBirthDate().AsTime(),
		Email:          ubUser.Email,
		Username:       ubUser.Username,
		Password:       ubUser.Password,
		Education:      []model.EducationItem{}, //za potrebe testiranja
		WorkExperience: []model.WorkExperienceItem{},
		Skills:         []model.SkillItem{},
		Interests:      []model.InterestItem{},
		BlockedUsers:   []string{},
	}
}
