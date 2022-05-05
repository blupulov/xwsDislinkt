package model

type ServiceInterface interface {
	Insert(userId string) error

	Follow(soureUserId, targetUserId string) error
	UnFollow(sourceUserId, targetUserId string) error

	SendRequest(sourceUserId, targetUserId string) error
	DeleteRequest(requestSenderId, requestReceiverId string) error

	GetAllUserFollowers(userId string) ([]*User, error)
	GetAllFollowingUsers(userId string) ([]*User, error)
	GetAllReceivedRequests(userId string) ([]*User, error)
	GetAllSentRequests(userId string) ([]*User, error)

	RequestAnswer(requestSenderId, requestReceiverId string, accept bool) error
}
