package model

import (
	"strconv"
)

type AuthToken struct {
	UID   int64
	Token string
}

func (u *AuthToken) String() string {
	return "[AuthToken: " + strconv.FormatInt(u.UID, 10) +
		", AuthToken: " + u.Token +
		"]"
}
