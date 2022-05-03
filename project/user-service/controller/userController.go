package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blupulov/xwsDislinkt/user-service/dto"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//Dodati i vracanje ID-ja
//logout samo se napravi f-ja koja postavlja exipres u proslost i salje prazan cookie
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginDto dto.LoginDto
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := uc.us.Login(loginDto.Password, loginDto.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cookie := http.Cookie{Name: "jwt", Value: token, Expires: time.Now().Add(time.Hour * 24)}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusAccepted)
}

func (uc *UserController) AddExpirience(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//var expirience model.WorkExperienceItem
	expirience := model.WorkExperienceItem{}
	json.NewDecoder(r.Body).Decode(&expirience)
	err = uc.us.AddExpirience(&expirience, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
