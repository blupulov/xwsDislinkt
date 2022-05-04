package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInterface interface {
	GetById(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Register(user *User) error
	DeleteById(id primitive.ObjectID) error
	Login(password, username string) (string, error)
	AddExpirience(expirience *WorkExperienceItem, expirienceID primitive.ObjectID) error
	AddSkill(skill *SkillItem, userID primitive.ObjectID) error
}
