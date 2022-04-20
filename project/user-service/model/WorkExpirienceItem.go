package model

type WorkExpirienceItem struct {
	CompanyName          string `bson:"companyName" json:"companyName"`
	WorkDurationInMonths int    `bson:"workDurationInMonths" json:"workDurationInMonths"`
	EmploymentField      string `bson:"employmentField" json:"employmentField"`
}
