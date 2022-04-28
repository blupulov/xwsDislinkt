package service

import (
	"context"
	"errors"
	"time"

	"github.com/blupulov/xwsDislinkt/user-service/model"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	DATABASE   = "user-service"
	COLLECTION = "users"
	SecretKey  = "secret"
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

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(password)

	userExists, err := ps.findUserByUsername(user.Username)
	if userExists != nil {
		return errors.New("User already exists")
	}

	_, err = ps.usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserServiceImpl) Login(password, username string) (string, error) {
	user, err := us.findUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		//internal server error
		return "", err
	}

	return token, nil
}

func (us *UserServiceImpl) findUserByUsername(username string) (*model.User, error) {
	filter := bson.M{"username": username}
	sr := us.usersCollection.FindOne(context.TODO(), filter)
	if sr.Err() != nil {
		return nil, sr.Err()
	}

	var user model.User
	sr.Decode(&user)

	return &user, nil
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
