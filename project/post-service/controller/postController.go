package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/post-service/model"
	"github.com/blupulov/xwsDislinkt/post-service/service"
	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	ps *service.PostService
}

func NewPostController(ps *service.PostService) *PostController {
	return &PostController{
		ps: ps,
	}
}

func (pc PostController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (pc PostController) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := model.Post{}

	json.NewDecoder(r.Body).Decode(&post)
	err := pc.ps.Insert(&post)
	if err != nil {
		//log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pc PostController) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
