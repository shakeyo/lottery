package api

import (
	"server/model"
)

type UserAPI interface {
	AuthUser(userID int64, userToken string) (*model.User, int)
	ModifyUserProperty(userID int64, token string) bool
}
