package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/post-service/model"
	"github.com/blupulov/xwsDislinkt/post-service/service"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct {
	ps *service.PostService
}

func NewPostController(ps *service.PostService) *PostController {
	return &PostController{
		ps: ps,
	}
}

func (pc *PostController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := pc.ps.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonPosts)
}

func (pc *PostController) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := model.Post{}

	json.NewDecoder(r.Body).Decode(&post)
	err := pc.ps.Insert(&post)
	if err != nil {
		//log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (pc *PostController) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	postId, err := primitive.ObjectIDFromHex(p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	post, err := pc.ps.GetById(postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonPost)
}

func (pc *PostController) DeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	postId, err := primitive.ObjectIDFromHex(p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = pc.ps.DeleteById(postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (pc *PostController) Like(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	postId, err := primitive.ObjectIDFromHex(p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := p.ByName("userId")

	err = pc.ps.Like(userId, postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (pc *PostController) Dislike(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	postId, err := primitive.ObjectIDFromHex(p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := p.ByName("userId")

	err = pc.ps.Dislike(userId, postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (pc *PostController) AddComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	postId, err := primitive.ObjectIDFromHex(p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	comment := model.Comment{}
	json.NewDecoder(r.Body).Decode(&comment)
	err = pc.ps.AddComment(&comment, postId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (pc *PostController) DeleteCommentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := pc.ps.DeleteCommentById(p.ByName("commentId"), p.ByName("postId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pc *PostController) GetAllByOwnerId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	posts, err := pc.ps.GetAllByOwnerId(p.ByName("ownerId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonPosts)
}
