package service

import "github.com/blupulov/xwsDislinkt/following-service/model"

type FollowService struct {
	fi model.ServiceInterface
}

func NewFollowService(si model.ServiceInterface) *FollowService {
	return &FollowService{
		fi: si,
	}
}

func (fs *FollowService) Insert(userId string) error {
	return fs.fi.Insert(userId)
}

func (fs *FollowService) Follow(sourceUserId, targetUserId string) error {
	return fs.fi.Follow(sourceUserId, targetUserId)
}

func (fs *FollowService) UnFollow(sourceUserId, targetUserId string) error {
	return fs.fi.UnFollow(sourceUserId, targetUserId)
}

func (fs *FollowService) SendRequest(sourceUserId, targetUserId string) error {
	return fs.fi.SendRequest(sourceUserId, targetUserId)
}

func (fs *FollowService) GetAllUserFollowers(userId string) ([]*model.User, error) {
	return fs.fi.GetAllUserFollowers(userId)
}

func (fs *FollowService) GetAllFollowingUsers(userId string) ([]*model.User, error) {
	return fs.fi.GetAllFollowingUsers(userId)
}

func (fs *FollowService) RequestAnswer(requestSenderId, requestReceverId string, answer bool) error {
	return fs.fi.RequestAnswer(requestSenderId, requestReceverId, answer)
}

func (fs *FollowService) GetAllReceivedRequests(userId string) ([]*model.User, error) {
	return fs.fi.GetAllReceivedRequests(userId)
}

func (fs *FollowService) GetAllSentRequests(userId string) ([]*model.User, error) {
	return fs.fi.GetAllSentRequests(userId)
}

func (fs *FollowService) DeleteRequest(requestSenderId, requestReceiverId string) error {
	return fs.fi.DeleteRequest(requestSenderId, requestReceiverId)
}
