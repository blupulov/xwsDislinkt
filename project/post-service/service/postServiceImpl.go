package service

import (
	"context"
	"time"

	"github.com/blupulov/xwsDislinkt/post-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//TODO: ADD DTO-OBJECTS FOR CREATING AND UPDATING
const (
	DATABASE   = "post-service"
	COLLECTION = "posts"
)

//Implementation of methods in PostInterface
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
	panic("Not implemented")
}

func (ps *PostServiceImpl) GetAll() ([]*model.Post, error) {
	filter := bson.D{{}}
	return ps.filterAll(filter)
}

func (ps *PostServiceImpl) GetAllByOwnerId(ownerId string) ([]*model.Post, error) {
	panic("Not implemented")
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
	panic("Not implemented")
}

func (ps *PostServiceImpl) Like(userId string, postId primitive.ObjectID) error {
	if err := ps.updateLikeIfHater(userId, postId); err == nil {
		return nil
	}
	err := ps.updateLikeIfNotHater(userId, postId)
	return err
}

func (ps *PostServiceImpl) Dislike(userId string, postId primitive.ObjectID) error {
	panic("Not implemented")
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
