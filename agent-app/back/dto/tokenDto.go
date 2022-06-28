package dto

type TokenDto struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	Id    string `json:"userId"`
}
