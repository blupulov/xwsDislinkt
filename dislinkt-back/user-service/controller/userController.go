package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blupulov/xwsDislinkt/user-service/dto"
	"github.com/blupulov/xwsDislinkt/user-service/model"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	"github.com/dgrijalva/jwt-go"
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
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *model.User

	user, err = uc.us.GetById(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v\n", user)
}

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

	jsonToken, err := json.Marshal(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "%s\n", jsonToken)
}

func (uc *UserController) AddExpirience(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	expirience := model.WorkExperienceItem{}
	json.NewDecoder(r.Body).Decode(&expirience)
	err = uc.us.AddExpirience(&expirience, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) AddSkill(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	skill := model.SkillItem{}
	json.NewDecoder(r.Body).Decode(&skill)
	err = uc.us.AddSkill(&skill, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) AddEducation(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	education := model.EducationItem{}
	json.NewDecoder(r.Body).Decode(&education)
	err = uc.us.AddEducation(&education, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) AddInterest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	interest := model.InterestItem{}
	json.NewDecoder(r.Body).Decode(&interest)
	err = uc.us.AddInterest(&interest, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) DeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.us.DeleteById(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) ChangeUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto := dto.ChangeUserDto{}
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.us.ChangeUser(userId, &dto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) GetUserByUsername(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")

	var user *model.User

	user, err := uc.us.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v\n", user)

}

/*
tokenString := "<YOUR TOKEN STRING>"
claims := jwt.MapClaims{}
token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte("<YOUR VERIFICATION KEY>"), nil
})
// ... error handling

// do something with decoded claims
for key, val := range claims {
    fmt.Printf("Key: %v, value: %v\n", key, val)
}
*/

func (uc *UserController) ReadToken(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(p.ByName("token"), claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(service.SecretKey), nil
	})
	if err != nil {
		fmt.Println("1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	exp := claims["exp"].(float64)
	fmt.Println(time.Unix(int64(exp), 0).UTC())
	w.WriteHeader(http.StatusOK)
}
