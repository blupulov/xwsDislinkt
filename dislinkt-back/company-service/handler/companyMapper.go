package handler

import (
	pb "github.com/blupulov/xwsDislinkt/common/proto/services/company-service"
	"github.com/blupulov/xwsDislinkt/company-service/model"
)

func mapPbCompanyFromCompany(company *model.Company) *pb.Company {
	pbCompany := &pb.Company{
		Id:          company.Id.Hex(),
		CompanyName: company.CompanyName,
		OwnerId:     company.OwnerId,
		Description: company.Description,
		Info: &pb.Info{
			Street:      company.Info.Street,
			City:        company.Info.City,
			Country:     company.Info.Country,
			Email:       company.Info.Email,
			PhoneNumber: company.Info.PhoneNumber,
		},
	}

	for _, job := range company.OpenJobs {
		pbCompany.OpenJobs = append(pbCompany.OpenJobs, mapPbJobFromJob(&job))
	}

	for _, comment := range company.Comments {
		pbCompany.Comments = append(pbCompany.Comments, mapPbCommentFromComment(&comment))
	}

	return pbCompany
}

func mapPbCommentFromComment(comment *model.Comment) *pb.Comment {
	return &pb.Comment{
		Id:             comment.Id.Hex(),
		CommentOwnerId: comment.CommentOwnerId,
		Comment:        comment.Comment,
		Grade:          int32(comment.Grade),
	}
}

func mapPbJobFromJob(job *model.Job) *pb.Job {
	return &pb.Job{
		Id:             job.Id.Hex(),
		JobDescription: job.JobDescription,
		JobName:        job.JobName,
	}
}

func mapCompanyFromPbCompany(pbCompany *pb.Company) *model.Company {
	return &model.Company{
		CompanyName: pbCompany.CompanyName,
		OwnerId:     pbCompany.OwnerId,
		IsAccepted:  false,
		Description: pbCompany.Description,
		Info: model.Info{
			Street:      pbCompany.Info.Street,
			City:        pbCompany.Info.City,
			Country:     pbCompany.Info.Country,
			Email:       pbCompany.Info.Email,
			PhoneNumber: pbCompany.Info.PhoneNumber,
		},
		OpenJobs: make([]model.Job, 0),
		Comments: make([]model.Comment, 0),
	}
}

func mapJobFromPbJob(pbJob *pb.Job) *model.Job {
	return &model.Job{
		JobName:        pbJob.JobName,
		JobDescription: pbJob.JobDescription,
	}
}

func mapCommentFromPbComment(pbComment *pb.Comment) *model.Comment {
	return &model.Comment{
		CommentOwnerId: pbComment.CommentOwnerId,
		Comment:        pbComment.Comment,
		Grade:          int(pbComment.Grade),
	}
}
