package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CompanyInterface interface {
	GetAll() ([]*Company, error)
	GetAllUnAcceptedCompanies() ([]*Company, error)
	GetAllByOwnerId(ownerId string) ([]*Company, error)
	GetById(companyId primitive.ObjectID) (*Company, error)
	CreateCompany(company *Company) error
	DeleteById(companyId primitive.ObjectID) error
	CreateComment(companyId primitive.ObjectID, comment *Comment) error
	DeleteCommentById(companyId primitive.ObjectID, commentId primitive.ObjectID) error
	CreateJob(companyId primitive.ObjectID, job *Job) error
	DeleteJobById(companyId primitive.ObjectID, jobId primitive.ObjectID) error
	EnableCompany(companyId primitive.ObjectID) error
}
