package service

import (
	"context"
	"errors"
	"time"

	"github.com/blupulov/xwsDislinkt/user-service/dto"
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
	filter := bson.M{"_id": id}
	return ps.filter(filter)
}

func (ps *UserServiceImpl) DeleteById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	return ps.usersCollection.FindOneAndDelete(context.TODO(), filter).Err()
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

	user.Role = model.ROLE_USER

	_, err = ps.usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
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

func (us *UserServiceImpl) AddExpirience(expirience *model.WorkExperienceItem, userID primitive.ObjectID) error {
	expirience.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": userID}
	updateFilter := bson.M{
		"$push": bson.M{"workExperienceCollection": expirience},
	}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

func (us *UserServiceImpl) AddSkill(skill *model.SkillItem, userID primitive.ObjectID) error {
	skill.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": userID}
	updateFilter := bson.M{
		"$push": bson.M{"skills": skill},
	}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

func (us *UserServiceImpl) AddEducation(education *model.EducationItem, userID primitive.ObjectID) error {
	education.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": userID}
	updateFilter := bson.M{
		"$push": bson.M{"educationCollection": education},
	}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

func (us *UserServiceImpl) AddInterest(interest *model.InterestItem, userID primitive.ObjectID) error {
	interest.Id = primitive.NewObjectID()

	findFilter := bson.M{"_id": userID}
	updateFilter := bson.M{
		"$push": bson.M{"interestsCollection": interest},
	}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

func (us *UserServiceImpl) ChangeUser(userID primitive.ObjectID, dto *dto.ChangeUserDto) error {
	if us.isNewUsernameTaken(userID, dto.Username) {
		return errors.New("username is already taken")
	}

	findFilter := bson.M{"_id": userID}
	updateFilter := bson.M{
		"$set": bson.M{
			"firstName":   dto.FirstName,
			"lastName":    dto.LastName,
			"username":    dto.Username,
			"biography":   dto.Biography,
			"Email":       dto.Email,
			"phoneNumber": dto.PhoneNumber,
			"birthDate":   dto.BirthDate,
		},
	}

	sr := us.usersCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
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

func (us *UserServiceImpl) GetManyUsersById(usersIds []string) (*[]model.User, error) {
	var users []model.User

	for _, id := range usersIds {
		userId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		user, err := us.GetById(userId)
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
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

func (ps *UserServiceImpl) filter(filter interface{}) (user *model.User, err error) {
	cur := ps.usersCollection.FindOne(context.TODO(), filter)

	err = cur.Decode(&user)

	return
}
