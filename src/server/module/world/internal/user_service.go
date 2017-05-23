package internal

import (
	"server/model"
	"sync"
)

var UserServiceInstance = makeUserService()

type UserService struct {
	users map[int64]*model.User
	sync.RWMutex
}

func makeUserService() *UserService {
	return &UserService{
		users: make(map[int64]*model.User),
	}
}

func (c *UserService) AddUser(userID int64, user *model.User) {

	_, exists := c.users[userID]
	if exists {
		return
	}

	c.users[userID] = user
}

func (c *UserService) RemoveUser(userID int64) bool {
	return false
}

func (c *UserService) GetUserByUID(userID int64) *model.User {
	return nil
}

func (c *UserService) SearchUserByToken(token string) *model.User {
	return nil
}
