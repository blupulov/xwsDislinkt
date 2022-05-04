package model

type ServiceInterface interface {
	Insert(userId string) error
	Follow(soureUserId, targetUserId string) error
	GetAllUserFollowers(userId string) ([]User, error)
	GetAllFollowingUsers(userId string) ([]User, error)
	GetAllRequestForFollowingOfUser(userId string) ([]User, error)
	SendRequest(sourceUserId, targetUserId string) error
	UnFollow(sourceUserId, targetUserId string) error
	RequestAnswer(requestSenderId, requestReceiverId string, accept bool) error
}
