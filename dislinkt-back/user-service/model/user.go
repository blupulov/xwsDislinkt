package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	FirstName    string             `bson:"firstName" json:"firstName"`
	LastName     string             `bson:"lastName" json:"lastName"`
	BirthDate    time.Time          `bson:"birthDate" json:"birthDate"`
	Email        string             `bson:"Email" json:"Email"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	PhoneNumber  string             `bson:"phoneNumber" json:"phoneNumber"`
	Biography    string             `bson:"biography" json:"biography"`
	Role         UserRole           `bson:"role" json:"role"`
	ProfileImage string             `bson:"profileImage" json:"profileImage"` //base64

	Education      []EducationItem      `bson:"educationCollection" json:"educationCollection"`
	WorkExperience []WorkExperienceItem `bson:"workExperienceCollection" json:"workExperienceCollection"`
	Skills         []SkillItem          `bson:"skills" json:"skills"`
	Interests      []InterestItem       `bson:"interestsCollection" json:"interestCollection"`
	BlockedUsers   []string             `bson:"blockedUsers" json:"blockedUsers"`
}

type UserRole int8

const (
	ROLE_USER UserRole = iota
	ROLE_ADMIN
	ROLE_COMPANY_OWNER
)

func (role UserRole) String() string {
	switch role {
	case ROLE_USER:
		return "USER"
	case ROLE_ADMIN:
		return "ADMIN"
	case ROLE_COMPANY_OWNER:
		return "COMPANY OWNER"
	}
	return "USER"
}
