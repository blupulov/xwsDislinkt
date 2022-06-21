package customHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blupulov/xwsDislinkt/api-gateway/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type CustomHandler struct {
	userClientAddress      string
	postClientAddress      string
	followingClientAddress string
	companyClientAddress   string
}

func NewCustomHandler(userClientAddr, postClientAddr, followingClientAddr, companyClientAddr string) *CustomHandler {
	return &CustomHandler{
		userClientAddress:      userClientAddr,
		postClientAddress:      postClientAddr,
		followingClientAddress: followingClientAddr,
		companyClientAddress:   companyClientAddr,
	}
}

func (ch *CustomHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/custom/user/{userId}/followers", ch.GetUserFollowers)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/custom/user/{userId}/following", ch.GetFollowingUsers)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/custom/user/{userId}/company/{companyId}", ch.EnableCompanyAndPromoteUser)
}

func (ch *CustomHandler) EnableCompanyAndPromoteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userId := pathParams["userId"]
	companyId := pathParams["companyId"]
	if userId == "" || companyId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := ch.enableCompany(companyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ch.promoteUserToCompanyOwner(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *CustomHandler) GetFollowingUsers(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userId := pathParams["userId"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followingUsersIds, err := ch.getFollowingUsers(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users []model.User
	for _, userID := range *followingUsersIds {
		user, err := ch.getUserById(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users = append(users, *user)
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", jsonUsers)
}

func (ch *CustomHandler) GetUserFollowers(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userId := pathParams["userId"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userFollowersIds, err := ch.getUserFollowers(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users []model.User
	for _, userID := range *userFollowersIds {
		user, err := ch.getUserById(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users = append(users, *user)
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", jsonUsers)
}
