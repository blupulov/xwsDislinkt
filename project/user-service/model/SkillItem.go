package model

type SkillItem struct {
	SkillName        string `bson:"skillName" json:"skillName"`
	SkillDescription string `bson:"skillDescription" json:"skillDescription"`
}
