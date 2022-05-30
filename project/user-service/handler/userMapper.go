package handler

import (
	pb "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"github.com/blupulov/xwsDislinkt/user-service/dto"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPbUserFromModel(user *model.User) *pb.User {
	pbUser := &pb.User{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: timestamppb.New(user.BirthDate),
		Email:     user.Email,
		Username:  user.Username,
		Biography: user.Biography,
		Role:      user.Role.String(),
		// DODATI DOBAVLJENJE NIZOVA IZPOD
		Education:      []*pb.EducationItem{}, //za potrebe testiranja,sve dole
		WorkExpirience: []*pb.WorkExpirienceItem{},
		Skills:         []*pb.SkillItem{},
		Interests:      []*pb.Interests{},
		BlockedUsers:   []string{},
		PhoneNumber:    user.PhoneNumber,
	}

	return pbUser
}

func mapUserForRegistration(newUser *pb.UserRegistrationModel) *model.User {
	return &model.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		//BirthDate:      newUser.GetBirthdate().AsTime(),
		Email:          newUser.Email,
		Username:       newUser.Username,
		Password:       newUser.Password,
		Biography:      newUser.Biography,
		PhoneNumber:    newUser.PhoneNumber,
		Education:      make([]model.EducationItem, 0), //za potrebe testiranja
		WorkExperience: make([]model.WorkExperienceItem, 0),
		Skills:         make([]model.SkillItem, 0),
		Interests:      make([]model.InterestItem, 0),
		BlockedUsers:   make([]string, 0),
	}
}

func mapUserForAddingSkill(newSkill *pb.AddSkillModel) *model.SkillItem {
	return &model.SkillItem{
		SkillName:        newSkill.SkillName,
		SkillDescription: newSkill.SkillDescription,
	}
}

func mapUserForAddingExpirience(newExpirience *pb.AddExpirienceModel) *model.WorkExperienceItem {
	return &model.WorkExperienceItem{
		CompanyName:          newExpirience.CompanyName,
		WorkDurationInMonths: int(newExpirience.WorkDurationInMonths),
		EmploymentField:      newExpirience.EmploymentField,
	}
}

func mapUserForAddingEducation(newEducation *pb.AddEducationModel) *model.EducationItem {
	return &model.EducationItem{
		SchoolName:       newEducation.SchoolName,
		SchoolType:       newEducation.SchoolType,
		DurationInMonths: int(newEducation.DurationInMonths),
	}
}

func mapUserForAddingInterest(newInterest *pb.AddInterestModel) *model.InterestItem {
	return &model.InterestItem{
		InterestName:        newInterest.InterestName,
		InterestDescription: newInterest.InterestDescription,
	}
}

func mapChangeUserDtoFromPb(userInfo *pb.ChangeUserInfo) *dto.ChangeUserDto {
	return &dto.ChangeUserDto{
		Username:    userInfo.Username,
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Biography:   userInfo.Biography,
		BirthDate:   userInfo.GetBirthDate().AsTime(),
	}
}
