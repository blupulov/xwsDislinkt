package model

type WorkExperienceItem struct {
	CompanyName          string `bson:"companyName" json:"companyName"`
	WorkDurationInMonths int    `bson:"workDurationInMonths" json:"workDurationInMonths"`
	EmploymentField      string `bson:"employmentField" json:"employmentField"`
}
