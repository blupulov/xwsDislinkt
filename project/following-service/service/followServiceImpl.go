package service

import (
	"github.com/blupulov/xwsDislinkt/following-service/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowServiceImpl struct {
	driver neo4j.Driver
}

func NewFollowServiceImpl(d neo4j.Driver) model.ServiceInterface {
	return &FollowServiceImpl{
		driver: d,
	}
}

//****************************************************************************************************
//PUBLIC METHODS
//****************************************************************************************************

func (fs *FollowServiceImpl) Insert(userId string) error {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return insertUser(tx, userId)
	})

	return err
}

func (fs *FollowServiceImpl) Follow(sourceUserId, targetUserId string) error {
	if err := fs.prepareUsers(sourceUserId, targetUserId); err != nil {
		return err
	}

	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return insertFollow(tx, sourceUserId, targetUserId)
	})

	return err
}

func (fs *FollowServiceImpl) SendRequest(sourceUserId, targetUserId string) error {
	if err := fs.prepareUsers(sourceUserId, targetUserId); err != nil {
		return err
	}

	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return insertRequest(tx, sourceUserId, targetUserId)
	})

	return err
}

func (fs *FollowServiceImpl) UnFollow(sourceUserId, targetUserId string) error {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return deleteFollow(tx, sourceUserId, targetUserId)
	})

	return err
}

func (fs *FollowServiceImpl) RequestAnswer(requestSenderId, requestReceiverId string, accept bool) error {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	if accept {
		_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			if _, err := deleteRequest(tx, requestSenderId, requestReceiverId); err != nil {
				tx.Rollback()
				return nil, err
			}
			if _, err := insertFollow(tx, requestSenderId, requestReceiverId); err != nil {
				tx.Rollback()
				return nil, err
			}
			return nil, nil
		})
		return err
	} else {
		_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return deleteRequest(tx, requestSenderId, requestReceiverId)
		})
		return err
	}
}

func (fs *FollowServiceImpl) GetAllUserFollowers(userId string) ([]*model.User, error) {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return getAllUserFollowers(tx, userId)
	})
	if err != nil {
		return nil, err
	}

	users := result.([]*model.User)
	return users, nil
}

func (fs *FollowServiceImpl) GetAllFollowingUsers(userId string) ([]*model.User, error) {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return getAllFollowingUsers(tx, userId)
	})
	if err != nil {
		return nil, err
	}

	users := result.([]*model.User)
	return users, nil
}

func (fs *FollowServiceImpl) GetAllReceivedRequests(userId string) ([]*model.User, error) {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return getAllReceivedRequests(tx, userId)
	})
	if err != nil {
		return nil, err
	}

	users := result.([]*model.User)
	return users, nil
}

func (fs *FollowServiceImpl) GetAllSentRequests(userId string) ([]*model.User, error) {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return getAllSentRequests(tx, userId)
	})
	if err != nil {
		return nil, err
	}

	users := result.([]*model.User)
	return users, nil
}

func (fs *FollowServiceImpl) DeleteRequest(requestSenderId, requestReceiverId string) error {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return deleteRequest(tx, requestSenderId, requestReceiverId)
	})
	return err
}

//****************************************************************************************************
//PRIVATE QUERY FUNCTIONS
//****************************************************************************************************

func insertFollow(tx neo4j.Transaction, sUserId, tUserId string) (interface{}, error) {
	query := "MATCH(n:User {userId: $sUserId}), (m:User {userId: $tUserId})" +
		"CREATE (n) -[:FOLLOW]-> (m)"
	params := map[string]interface{}{
		"sUserId": sUserId,
		"tUserId": tUserId,
	}

	return tx.Run(query, params)
}

func insertRequest(tx neo4j.Transaction, sUserId, tUserId string) (interface{}, error) {
	query := "MATCH(n:User {userId: $sUserId}), (m:User {userId: $tUserId})" +
		"CREATE (n) -[:REQUEST]-> (m)"
	params := map[string]interface{}{
		"sUserId": sUserId,
		"tUserId": tUserId,
	}

	return tx.Run(query, params)
}

func deleteFollow(tx neo4j.Transaction, sUserId, tUserId string) (interface{}, error) {
	query := "MATCH(n:User {userId: $sUserId}) -[f:FOLLOW]-> (m:User {userId: $tUserId}) DELETE f"
	params := map[string]interface{}{
		"sUserId": sUserId,
		"tUserId": tUserId,
	}

	return tx.Run(query, params)
}

func deleteRequest(tx neo4j.Transaction, sUserId, tUserId string) (interface{}, error) {
	query := "MATCH(n:User {userId: $sUserId}) -[r:REQUEST]-> (m:User {userId: $tUserId}) DELETE r"
	params := map[string]interface{}{
		"sUserId": sUserId,
		"tUserId": tUserId,
	}

	return tx.Run(query, params)
}

func getAllUserFollowers(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "MATCH (n:User) -[:FOLLOW]-> (m:User {userId: $userId}) return n.userId as userId"
	params := map[string]interface{}{
		"userId": userId,
	}

	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for result.Next() {
		uId := result.Record().GetByIndex(0).(string)
		user := &model.User{UserId: uId}
		users = append(users, user)
	}
	return users, nil
}

func getAllFollowingUsers(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "MATCH (n:User {userId: $userId}) -[:FOLLOW]-> (m:User) return m.userId as userId"
	params := map[string]interface{}{
		"userId": userId,
	}

	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for result.Next() {
		uId := result.Record().GetByIndex(0).(string)
		user := &model.User{UserId: uId}
		users = append(users, user)
	}
	return users, nil
}

func getAllReceivedRequests(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "MATCH (n:User) -[:REQUEST]-> (m:User {userId: $userId}) return n.userId as userId"
	params := map[string]interface{}{
		"userId": userId,
	}

	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for result.Next() {
		uId := result.Record().GetByIndex(0).(string)
		user := &model.User{UserId: uId}
		users = append(users, user)
	}
	return users, nil
}

func getAllSentRequests(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "MATCH (n:User {userId: $userId}) -[:REQUEST]-> (m:User) return m.userId as userId"
	params := map[string]interface{}{
		"userId": userId,
	}

	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for result.Next() {
		uId := result.Record().GetByIndex(0).(string)
		user := &model.User{UserId: uId}
		users = append(users, user)
	}
	return users, nil
}

func insertUser(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "CREATE (:User {userId: $userId})"
	params := map[string]interface{}{
		"userId": userId,
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func findUser(tx neo4j.Transaction, userId string) (interface{}, error) {
	query := "MATCH (n:User {userId: $userId}) return n.userId as userId"
	params := map[string]interface{}{
		"userId": userId,
	}
	result, err := tx.Run(query, params)
	user := &model.User{}
	for result.Next() {
		userId, _ := result.Record().Get("userId")
		user.UserId = userId.(string)
	}
	return user, err
}

//****************************************************************************************************
//PRIVATE METHODS
//****************************************************************************************************

func (fs *FollowServiceImpl) isUserExist(userId string) (bool, error) {
	session := fs.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return findUser(tx, userId)
	})
	if err != nil {
		return false, err
	}

	user := result.(*model.User)
	if user.UserId == "" {
		return false, nil
	}

	return true, nil
}

func (fs *FollowServiceImpl) prepareUsers(sourceUserId, targetUserId string) error {
	existSource, err := fs.isUserExist(sourceUserId)
	if err != nil {
		return err
	}

	existTarget, err := fs.isUserExist(targetUserId)
	if err != nil {
		return err
	}

	if !existSource {
		if err := fs.Insert(sourceUserId); err != nil {
			return err
		}
	}

	if !existTarget {
		if err := fs.Insert(targetUserId); err != nil {
			return err
		}
	}

	return nil
}
