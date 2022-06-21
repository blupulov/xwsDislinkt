package model

import (
	"github.com/blupulov/xwsDislinkt/agent-app/back/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInterface interface {
	GetById(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Register(user *User) error
	DeleteById(id primitive.ObjectID) error
	Login(password, username string) (*dto.TokenDto, error)
	//ChangeUser(userID primitive.ObjectID, dto *dto.ChangeUserDto) error
	GetUserByUsername(username string) (*User, error)
	PromoteUserToCompanyOwner(userId primitive.ObjectID) error
}
