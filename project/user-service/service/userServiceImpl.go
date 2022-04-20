package service

import (
	"context"

	"github.com/blupulov/xwsDislinkt/user-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user-service"
	COLLECTION = "users"
)

type UserServiceImpl struct {
	usersCollection *mongo.Collection
}

func NewUserServiceImpl(client *mongo.Client) model.UserInterface {
	uc := client.Database(DATABASE).Collection(COLLECTION)
	return &UserServiceImpl{
		usersCollection: uc,
	}
}

func (ps *UserServiceImpl) GetById(id primitive.ObjectID) (*model.User, error) {
	panic("Not implemented")
}

func (ps *UserServiceImpl) DeleteById(id primitive.ObjectID) error {
	panic("Not implemented")
}

func (ps *UserServiceImpl) GetAll() ([]*model.User, error) {
	filter := bson.D{{}}
	return ps.filterAll(filter)
}

func (ps *UserServiceImpl) Register(user *model.User) error {

	user.Id = primitive.NewObjectID()

	_, err := ps.usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

//private method for applying filters on db
func (ps *UserServiceImpl) filterAll(filter interface{}) ([]*model.User, error) {
	cur, err := ps.usersCollection.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cur)
}

//private method for decoding all data in cursor after filtering db
func decode(cur *mongo.Cursor) (users []*model.User, err error) {
	for cur.Next(context.TODO()) {
		var user model.User
		err = cur.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cur.Err()
	return
}
