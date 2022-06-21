package dto

import "time"

type ChangeUserDto struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Biography   string    `json:"biography"`
	BirthDate   time.Time `json:"birthDate"`
}
