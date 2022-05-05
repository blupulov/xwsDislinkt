package model

import (
	"github.com/blupulov/xwsDislinkt/user-service/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInterface interface {
	GetById(id primitive.ObjectID) (*User, error) // treba uraditi
	GetAll() ([]*User, error)
	Register(user *User) error
	DeleteById(id primitive.ObjectID) error // treba uraditi
	Login(password, username string) (string, error)
	AddExpirience(expirience *WorkExperienceItem, expirienceID primitive.ObjectID) error
	AddSkill(skill *SkillItem, userID primitive.ObjectID) error
	AddEducation(education *EducationItem, userID primitive.ObjectID) error
	AddInterest(interest *InterestItem, userID primitive.ObjectID) error
	ChangeUser(userID primitive.ObjectID, dto *dto.ChangeUserDto) error // treba istraziti i uraditi
}
