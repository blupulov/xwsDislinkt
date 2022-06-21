package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/agent-app/back/service"
	"github.com/julienschmidt/httprouter"
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
