package controller

import (
	"net/http"
	"strconv"

	"github.com/blupulov/xwsDislinkt/following-service/service"
	"github.com/julienschmidt/httprouter"
)

type FollowController struct {
	fs *service.FollowService
}

func NewFollowController(fs *service.FollowService) *FollowController {
	return &FollowController{
		fs: fs,
	}
}

func (fc *FollowController) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := fc.fs.Insert(p.ByName("userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (fc *FollowController) Follow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := fc.fs.Follow(p.ByName("userId"), p.ByName("targetUserId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (fc *FollowController) UnFollow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := fc.fs.UnFollow(p.ByName("userId"), p.ByName("targetUserId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (fc *FollowController) SendRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := fc.fs.SendRequest(p.ByName("userId"), p.ByName("targetUserId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (fc *FollowController) RequestAnswer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	answer, err := strconv.ParseBool(p.ByName("answer"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = fc.fs.RequestAnswer(p.ByName("requestSenderId"), p.ByName("requestReceiverId"), answer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
