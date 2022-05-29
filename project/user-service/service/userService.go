package service

import (
	"github.com/blupulov/xwsDislinkt/user-service/dto"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	ui model.UserInterface
}

func NewUserService(ui model.UserInterface) *UserService {
	return &UserService{
		ui: ui,
	}
}

func (us *UserService) GetById(id primitive.ObjectID) (*model.User, error) {
	return us.ui.GetById(id)
}

func (us *UserService) Register(user *model.User) error {
	return us.ui.Register(user)
}

func (us *UserService) GetAll() ([]*model.User, error) {
	return us.ui.GetAll()
}

func (us *UserService) DeleteById(id primitive.ObjectID) error {
	return us.ui.DeleteById(id)
}

func (us *UserService) Login(password, username string) (*dto.TokenDto, error) {
	return us.ui.Login(password, username)
}

func (us *UserService) AddExpirience(expirience *model.WorkExperienceItem, expirienceID primitive.ObjectID) error {
	return us.ui.AddExpirience(expirience, expirienceID)
}

func (us *UserService) AddSkill(skill *model.SkillItem, userID primitive.ObjectID) error {
	return us.ui.AddSkill(skill, userID)
}

func (us *UserService) AddEducation(education *model.EducationItem, userID primitive.ObjectID) error {
	return us.ui.AddEducation(education, userID)
}

func (us *UserService) AddInterest(interest *model.InterestItem, userID primitive.ObjectID) error {
	return us.ui.AddInterest(interest, userID)
}

func (us *UserService) ChangeUser(userId primitive.ObjectID, dto *dto.ChangeUserDto) error {
	return us.ui.ChangeUser(userId, dto)
}

func (us *UserService) GetManyUsersById(usersIds []string) (*[]model.User, error) {
	return us.ui.GetManyUsersById(usersIds)
}

func (us *UserService) GetUserByUsername(username string) (*model.User, error) {
	return us.ui.GetUserByUsername(username)
}
