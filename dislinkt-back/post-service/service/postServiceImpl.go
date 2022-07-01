package service

import (
	"context"
	"errors"
	"time"

	"github.com/blupulov/xwsDislinkt/post-service/dto"
	"github.com/blupulov/xwsDislinkt/post-service/model"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post-service"
	COLLECTION = "posts"
)

type PostServiceImpl struct {
	postsCollection *mongo.Collection
}

func NewPostServiceImpl(client *mongo.Client) model.PostInterface {
	pc := client.Database(DATABASE).Collection(COLLECTION)
	return &PostServiceImpl{
		postsCollection: pc,
	}
}

func (ps *PostServiceImpl) GetById(id primitive.ObjectID) (*model.Post, error) {
	filter := bson.M{"_id": id}
	return ps.filter(filter)
}

func (ps *PostServiceImpl) GetAll() ([]*model.Post, error) {
	filter := bson.D{{}}
	return ps.filterAll(filter)
}

func (ps *PostServiceImpl) GetAllByOwnerId(ownerId string) ([]*model.Post, error) {
	filter := bson.M{"postOwnerId": ownerId}
	return ps.filterAll(filter)
}

func (ps *PostServiceImpl) Insert(post *model.Post) error {
	post.PostDate = time.Now()
	post.Id = primitive.NewObjectID()

	_, err := ps.postsCollection.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostServiceImpl) DeleteById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	sr := ps.postsCollection.FindOneAndDelete(context.TODO(), filter)
	return sr.Err()
}

func (ps *PostServiceImpl) Like(userId string, postId primitive.ObjectID) error {
	if err := ps.updateLikeIfHater(userId, postId); err == nil {
		return nil
	}
	err := ps.updateLikeIfNotHater(userId, postId)
	return err
}

func (ps *PostServiceImpl) Dislike(userId string, postId primitive.ObjectID) error {
	if err := ps.updateDislikeIfFan(userId, postId); err == nil {
		return nil
	}
	err := ps.updateDislikeIfNotFan(userId, postId)
	return err
}

func (ps *PostServiceImpl) AddComment(comment *model.Comment, postId primitive.ObjectID) error {
	comment.Id = primitive.NewObjectID()
	comment.CreationDate = time.Now()
	findFilter := bson.M{"_id": postId}
	updateFilter := bson.M{
		"$push": bson.M{"comments": comment},
	}
	sr := ps.postsCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

//Only comment owner// postOwner??
func (ps *PostServiceImpl) DeleteCommentById(commentId, postId string) error {
	pId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	cId, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return err
	}
	var filterArray []primitive.ObjectID
	filterArray = append(filterArray, cId)
	findFilter := bson.M{"_id": pId}
	updateFilter := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": cId,
			},
		},
	}
	sr := ps.postsCollection.FindOneAndUpdate(context.TODO(), findFilter, updateFilter)
	return sr.Err()
}

//private method for applay filters on db
func (ps *PostServiceImpl) filterAll(filter interface{}) ([]*model.Post, error) {
	cur, err := ps.postsCollection.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cur)
}

//private method for decode all data in cursor after filtering db
func decode(cur *mongo.Cursor) (posts []*model.Post, err error) {
	for cur.Next(context.TODO()) {
		var post model.Post
		err = cur.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cur.Err()
	return
}

//private method for like post
func (ps *PostServiceImpl) updateLikeIfHater(userId string, postId primitive.ObjectID) error {
	//$in - work with array only
	var userIds []string
	userIds = append(userIds, userId)

	checkFilter := bson.M{
		"_id":       postId,
		"hatersIds": bson.M{"$in": userIds},
	}

	updateFilter := bson.D{
		{"$inc", bson.D{{"postLikeNumber", 1}}},
		{"$push", bson.D{{"fansIds", userId}}},
		{"$pull", bson.D{{"hatersIds", userId}}},
		{"$inc", bson.D{{"postDislikeNumber", -1}}},
	}

	//throw error if cant find post with specific checkFilter
	ur := ps.postsCollection.FindOneAndUpdate(context.TODO(), checkFilter, updateFilter)
	return ur.Err()
}

//private method for like post
func (ps *PostServiceImpl) updateLikeIfNotHater(userId string, postId primitive.ObjectID) error {
	findFilter := bson.M{"_id": postId}
	updateFilter := bson.D{
		{"$inc", bson.D{{"postLikeNumber", 1}}},
		{"$push", bson.D{{"fansIds", userId}}},
	}

	_, err := ps.postsCollection.UpdateOne(context.TODO(), findFilter, updateFilter)
	return err
}

func (ps *PostServiceImpl) filter(filter interface{}) (post *model.Post, err error) {
	cur := ps.postsCollection.FindOne(context.TODO(), filter)

	err = cur.Decode(&post)

	return
}

func (ps *PostServiceImpl) updateDislikeIfFan(userId string, postId primitive.ObjectID) error {
	var userIds []string
	userIds = append(userIds, userId)

	checkFilter := bson.M{
		"_id":     postId,
		"fansIds": bson.M{"$in": userIds},
	}

	updateFilter := bson.D{
		{"$inc", bson.D{{"postDislikeNumber", 1}}},
		{"$push", bson.D{{"hatersIds", userId}}},
		{"$pull", bson.D{{"fansIds", userId}}},
		{"$inc", bson.D{{"postLikeNumber", -1}}},
	}

	ur := ps.postsCollection.FindOneAndUpdate(context.TODO(), checkFilter, updateFilter)
	return ur.Err()
}

func (ps *PostServiceImpl) updateDislikeIfNotFan(userId string, postId primitive.ObjectID) error {
	findFilter := bson.M{"_id": postId}
	updateFilter := bson.D{
		{"$inc", bson.D{{"postDislikeNumber", 1}}},
		{"$push", bson.D{{"hatersIds", userId}}},
	}

	_, err := ps.postsCollection.UpdateOne(context.TODO(), findFilter, updateFilter)
	return err
}

func (ps *PostServiceImpl) CreatePostFromJob(dto *dto.ShareJobDto) error {
	postOwnerId, err := validateApiToken(dto.Token)
	if err != nil {
		return err
	}

	var post model.Post
	post.PostOwnerId = postOwnerId
	post.PostComment = "Job name: " + dto.JobName + "\n" +
		"Company name: " + dto.CompanyName + "\n" +
		"DisJobUsername: " + dto.DisJobUsername
	post.Comments = make([]model.Comment, 0)
	post.HatersIds = make([]string, 0)
	post.FansIds = make([]string, 0)

	err = ps.Insert(&post)
	if err != nil {
		return err
	}

	return nil
}

func validateApiToken(token string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	iss := claims["iss"].(string)
	exp := claims["exp"].(float64)

	now := time.Now().UTC()
	expires := time.Unix(int64(exp), 0).UTC()

	if now.After(expires) {
		return "", errors.New("Api token expires")
	}

	return iss, nil
}
