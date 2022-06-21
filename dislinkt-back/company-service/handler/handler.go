package handler

import (
	"context"

	pb "github.com/blupulov/xwsDislinkt/common/proto/services/company-service"
	"github.com/blupulov/xwsDislinkt/company-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyHandler struct {
	pb.UnimplementedCompanyServiceServer
	cs *service.CompanyService
}

func NewCompanyHandler(companyService *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		cs: companyService,
	}
}

func (ch *CompanyHandler) GetAllUnAccepted(ctx context.Context, r *pb.GetAllUnAcceptedRequest) (*pb.GetAllUnAcceptedResponse, error) {
	var response pb.GetAllUnAcceptedResponse

	companies, err := ch.cs.GetAllUnAcceptedCompanies()
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		response.Companies = append(response.Companies, mapPbCompanyFromCompany(company))
	}

	return &response, nil
}

func (ch *CompanyHandler) GetAll(ctx context.Context, r *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	var response pb.GetAllResponse

	companies, err := ch.cs.GetAll()
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		response.Companies = append(response.Companies, mapPbCompanyFromCompany(company))
	}

	return &response, nil
}

func (ch *CompanyHandler) GetAllByOwnerId(ctx context.Context, r *pb.GetAllByOwnerIdRequest) (*pb.GetAllByOwnerIdResponse, error) {
	var response pb.GetAllByOwnerIdResponse

	companies, err := ch.cs.GetAllByOwnerId(r.OwnerId)
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		response.Companies = append(response.Companies, mapPbCompanyFromCompany(company))
	}

	return &response, nil
}

func (ch *CompanyHandler) GetById(ctx context.Context, r *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	var response pb.GetByIdResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	company, err := ch.cs.GetById(companyId)
	if err != nil {
		return nil, err
	}

	response.Company = mapPbCompanyFromCompany(company)

	return &response, nil
}

func (ch *CompanyHandler) CreateCompany(ctx context.Context, r *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	var response pb.CreateCompanyResponse

	company := mapCompanyFromPbCompany(r.Company)

	err := ch.cs.CreateCompany(company)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) DeleteCompany(ctx context.Context, r *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	var response pb.DeleteCompanyResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	err = ch.cs.DeleteById(companyId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) EnableCompany(ctx context.Context, r *pb.EnableCompanyRequest) (*pb.EnableCompanyResponse, error) {
	var response pb.EnableCompanyResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	err = ch.cs.EnableCompany(companyId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) CreateJob(ctx context.Context, r *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	var response pb.CreateJobResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	job := mapJobFromPbJob(r.Job)

	err = ch.cs.CreateJob(companyId, job)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) DeleteJob(ctx context.Context, r *pb.DeleteJobRequest) (*pb.DeleteJobResponse, error) {
	var response pb.DeleteJobResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	jobId, err := primitive.ObjectIDFromHex(r.JobId)
	if err != nil {
		return nil, err
	}

	err = ch.cs.DeleteJobById(companyId, jobId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	var response pb.CreateCommentResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	comment := mapCommentFromPbComment(r.Comment)

	err = ch.cs.CreateComment(companyId, comment)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (ch *CompanyHandler) DeleteComment(ctx context.Context, r *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	var response pb.DeleteCommentResponse

	companyId, err := primitive.ObjectIDFromHex(r.CompanyId)
	if err != nil {
		return nil, err
	}

	commentId, err := primitive.ObjectIDFromHex(r.CommentId)
	if err != nil {
		return nil, err
	}

	err = ch.cs.DeleteCommentById(companyId, commentId)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
