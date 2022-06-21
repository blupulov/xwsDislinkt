package service

import (
	"github.com/blupulov/xwsDislinkt/company-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	ci model.CompanyInterface
}

func NewCompanyService(ui model.CompanyInterface) *CompanyService {
	return &CompanyService{
		ci: ui,
	}
}

func (cs *CompanyService) GetAll() ([]*model.Company, error) {
	return cs.ci.GetAll()
}

func (cs *CompanyService) GetAllUnAcceptedCompanies() ([]*model.Company, error) {
	return cs.ci.GetAllUnAcceptedCompanies()
}

func (cs *CompanyService) GetAllByOwnerId(ownerId string) ([]*model.Company, error) {
	return cs.ci.GetAllByOwnerId(ownerId)
}

func (cs *CompanyService) GetById(companyId primitive.ObjectID) (*model.Company, error) {
	return cs.ci.GetById(companyId)
}

func (cs *CompanyService) CreateCompany(company *model.Company) error {
	return cs.ci.CreateCompany(company)
}

func (cs *CompanyService) DeleteById(companyId primitive.ObjectID) error {
	return cs.ci.DeleteById(companyId)
}

func (cs *CompanyService) CreateComment(companyId primitive.ObjectID, comment *model.Comment) error {
	return cs.ci.CreateComment(companyId, comment)
}

func (cs *CompanyService) DeleteCommentById(companyId primitive.ObjectID, commentId primitive.ObjectID) error {
	return cs.ci.DeleteCommentById(companyId, commentId)
}

func (cs *CompanyService) CreateJob(companyId primitive.ObjectID, job *model.Job) error {
	return cs.ci.CreateJob(companyId, job)
}

func (cs *CompanyService) DeleteJobById(companyId primitive.ObjectID, jobId primitive.ObjectID) error {
	return cs.ci.DeleteJobById(companyId, jobId)
}

func (cs *CompanyService) EnableCompany(companyId primitive.ObjectID) error {
	return cs.ci.EnableCompany(companyId)
}
