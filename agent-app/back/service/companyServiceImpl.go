package service

import (
	"context"

	"github.com/blupulov/xwsDislinkt/agent-app/back/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE           = "agent-app"
	COMPANY_COLLECTION = "companies"
)

type CompanyServiceImpl struct {
	companiesCollection *mongo.Collection
}

func NewCompanyServiceImpl(client *mongo.Client) model.CompanyInterface {
	cc := client.Database(DATABASE).Collection(COMPANY_COLLECTION)
	return &CompanyServiceImpl{
		companiesCollection: cc,
	}
}

func (cs *CompanyServiceImpl) GetAll() ([]*model.Company, error) {
	filter := bson.D{{}}
	return cs.filterAll(filter)
}

func (cs *CompanyServiceImpl) GetAllUnAcceptedCompanies() ([]*model.Company, error) {
	filter := bson.M{"isAccepted": false}
	return cs.filterAll(filter)
}

func (cs *CompanyServiceImpl) GetAllByOwnerId(ownerId string) ([]*model.Company, error) {
	filter := bson.M{"ownerId": ownerId, "isAccepted": true}
	return cs.filterAll(filter)
}

func (cs *CompanyServiceImpl) GetById(companyId primitive.ObjectID) (*model.Company, error) {
	filter := bson.M{"_id": companyId}
	return cs.filter(filter)
}

func (cs *CompanyServiceImpl) CreateCompany(company *model.Company) error {
	company.Id = primitive.NewObjectID()

	_, err := cs.companiesCollection.InsertOne(context.TODO(), company)

	return err
}

func (cs *CompanyServiceImpl) DeleteById(companyId primitive.ObjectID) error {
	filter := bson.M{"_id": companyId}

	_, err := cs.companiesCollection.DeleteOne(context.TODO(), filter)

	return err
}

func (cs *CompanyServiceImpl) CreateComment(companyId primitive.ObjectID, comment *model.Comment) error {
	comment.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": companyId}
	updateFilter := bson.M{"$push": bson.M{"comments": comment}}

	sr := cs.companiesCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func (cs *CompanyServiceImpl) DeleteCommentById(companyId primitive.ObjectID, commentId primitive.ObjectID) error {
	findFilter := bson.M{"_id": companyId}
	updateFilter := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": commentId}}}

	sr := cs.companiesCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func (cs *CompanyServiceImpl) CreateJob(companyId primitive.ObjectID, job *model.Job) error {
	job.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": companyId}
	updateFilter := bson.M{"$push": bson.M{"jobs": job}}

	sr := cs.companiesCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func (cs *CompanyServiceImpl) DeleteJobById(companyId primitive.ObjectID, jobId primitive.ObjectID) error {
	findFilter := bson.M{"_id": companyId}
	updateFilter := bson.M{
		"$pull": bson.M{
			"jobs": bson.M{
				"_id": jobId}}}

	sr := cs.companiesCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func (cs *CompanyServiceImpl) EnableCompany(companyId primitive.ObjectID) error {
	findFilter := bson.M{"_id": companyId}
	updateFilter := bson.M{"$set": bson.M{"isAccepted": true}}

	sr := cs.companiesCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func decode(cur *mongo.Cursor) (companies []*model.Company, err error) {
	for cur.Next(context.TODO()) {
		var company model.Company
		err = cur.Decode(&company)
		if err != nil {
			return
		}
		companies = append(companies, &company)
	}
	err = cur.Err()
	return
}

func (cs *CompanyServiceImpl) filter(filter interface{}) (company *model.Company, err error) {
	cur := cs.companiesCollection.FindOne(context.TODO(), filter)

	err = cur.Decode(&company)

	return
}

func (cs *CompanyServiceImpl) filterAll(filter interface{}) ([]*model.Company, error) {
	cur, err := cs.companiesCollection.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return decode(cur)
}
