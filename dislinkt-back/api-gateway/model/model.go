package model

type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	ProfileImage string `json:"image"`
}
