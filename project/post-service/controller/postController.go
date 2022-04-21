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
