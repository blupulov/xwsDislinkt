package model

type EducationItem struct {
	SchoolName       string `bson:"schoolName" json:"schoolName"`
	SchoolType       string `bson:"schoolType" json:"schoolType"` //greska u modelu
	DurationInMonths int    `bson:"durationInMonths" json:"durationInMonths"`
}
