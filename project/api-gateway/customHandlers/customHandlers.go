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
}

func NewCustomHandler(userClientAddr, postClientAddr, followingClientAddr string) *CustomHandler {
	return &CustomHandler{
		userClientAddress:      userClientAddr,
		postClientAddress:      postClientAddr,
		followingClientAddress: followingClientAddr,
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
