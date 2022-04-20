package model

type InterestItem struct {
	InterestName        string `bson:"interestName" json:"interestName"`
	InterestDescription string `bson:"interestDescription" json:"interestDescription"`
}
