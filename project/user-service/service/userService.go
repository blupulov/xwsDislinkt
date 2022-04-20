package service

import (
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
