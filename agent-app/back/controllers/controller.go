package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/agent-app/back/dto"
	"github.com/blupulov/xwsDislinkt/agent-app/back/model"
	"github.com/blupulov/xwsDislinkt/agent-app/back/service"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	cs *service.CompanyService
	us *service.UserService
}

func NewController(cs *service.CompanyService, us *service.UserService) *Controller {
	return &Controller{
		cs: cs,
		us: us,
	}
}

func (c *Controller) GetAllCompanies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	companies, err := c.cs.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonCompanies, err := json.Marshal(companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonCompanies)
}

func (c *Controller) GetAllUnAcceptedCompanies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	companies, err := c.cs.GetAllUnAcceptedCompanies()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonCompanies, err := json.Marshal(companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonCompanies)
}

func (c *Controller) GetAllByOwnerId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	companies, err := c.cs.GetAllByOwnerId(p.ByName("ownerId"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonCompanies, err := json.Marshal(companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonCompanies)
}

func (c *Controller) GetCompanyById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	companyId, err := primitive.ObjectIDFromHex(p.ByName("companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	company, err := c.cs.GetById(companyId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonCompany, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonCompany)
}

func (c *Controller) CreateCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	company := model.Company{}
	json.NewDecoder(r.Body).Decode(&company)

	err := c.cs.CreateCompany(&company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) EnableCompany(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	companyId, err := primitive.ObjectIDFromHex(p.ByName("companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ownerId, err := primitive.ObjectIDFromHex(p.ByName("ownerId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.cs.EnableCompany(companyId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.us.PromoteUserToCompanyOwner(ownerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//****************************************************************************************************
//USERS
//****************************************************************************************************

func (c *Controller) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)
	err := c.us.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := primitive.ObjectIDFromHex(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *model.User

	user, err = c.us.GetById(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", user)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginDto dto.LoginDto
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := c.us.Login(loginDto.Password, loginDto.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonToken)
}
