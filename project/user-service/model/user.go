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
	Password     string             `bson:"password" json:"password"` // hes?
	PhoneNumber  string             `bson:"phoneNumber" json:"phoneNumber"`
	Biography    string             `bson:"biography" json:"biography"`
	IsAdmin      bool               `bson:"isAdmin" json:"isAdmin"`
	ProfileImage string             `bson:"profileImage" json:"profileImage"` //base64

	Education      []EducationItem      `bson:"educationCollection" json:"educationCollection"` // Nije bas dobra praksa imenovanja, konsultovati se
	WorkExperience []WorkExperienceItem `bson:"workExperienceCollection" json:"workExperienceCollection"`
	Skills         []SkillItem          `bson:"skills" json:"skills"`
	Interests      []InterestItem       `bson:"interestsCollection" json:"interestCollection"`
	BlockedUsers   []string             `bson:"blockedUsers" json:"blockedUsers"`
}
