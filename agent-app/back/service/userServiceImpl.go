package service

import (
	"context"
	"errors"
	"time"

	"github.com/blupulov/xwsDislinkt/agent-app/back/dto"
	"github.com/blupulov/xwsDislinkt/agent-app/back/model"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER_COLLECTION = "users"
	SecretKey       = "secret"
)

type UserServiceImpl struct {
	usersCollection *mongo.Collection
}

func NewUserServiceImpl(client *mongo.Client) model.UserInterface {
	uc := client.Database(DATABASE).Collection(USER_COLLECTION)
	return &UserServiceImpl{
		usersCollection: uc,
	}
}

func (us *UserServiceImpl) PromoteUserToCompanyOwner(userId primitive.ObjectID) error {
	findFilter := bson.M{"_id": userId}
	updateFilter := bson.M{"$set": bson.M{"role": model.ROLE_COMPANY_OWNER}}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)

	return sr.Err()
}

func (ps *UserServiceImpl) GetById(id primitive.ObjectID) (*model.User, error) {
	filter := bson.M{"_id": id}
	return ps.filterUser(filter)
}

func (us *UserServiceImpl) DeleteById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	return us.usersCollection.FindOneAndDelete(context.TODO(), filter).Err()
}

func (us *UserServiceImpl) GetAll() ([]*model.User, error) {
	filter := bson.D{{}}
	return us.filterAllUsers(filter)
}

func (us *UserServiceImpl) Register(user *model.User) error {
	user.Id = primitive.NewObjectID()

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(password)

	userExists, err := us.findUserByUsername(user.Username)
	if userExists != nil {
		return errors.New("User already exists")
	}

	user.Role = model.ROLE_USER

	_, err = us.usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserServiceImpl) GetUserByUsername(username string) (user *model.User, err error) {
	user, err = us.findUserByUsername(username)
	if err != nil {
		return
	}

	return
}

func (us *UserServiceImpl) Login(password, username string) (*dto.TokenDto, error) {
	user, err := us.findUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	dto := dto.TokenDto{
		Role:  user.Role.String(),
		Token: token,
		Id:    user.Id.Hex(),
	}

	return &dto, nil
}

//****************************************************************************************************
//PRIVATE METHODS
//****************************************************************************************************

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

func (us *UserServiceImpl) filterAllUsers(filter interface{}) ([]*model.User, error) {
	cur, err := us.usersCollection.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeUser(cur)
}

func decodeUser(cur *mongo.Cursor) (users []*model.User, err error) {
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

func (us *UserServiceImpl) filterUser(filter interface{}) (user *model.User, err error) {
	cur := us.usersCollection.FindOne(context.TODO(), filter)

	err = cur.Decode(&user)

	return
}

func (us *UserServiceImpl) isNewUsernameTaken(userId primitive.ObjectID, username string) bool {
	user, err := us.findUserByUsername(username)
	if err != nil {
		return false
	}

	if user.Id != userId {
		return true
	}

	return false
}
