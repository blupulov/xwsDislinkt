package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/user-service/model"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := uc.us.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonUsers)
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)
	err := uc.us.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
