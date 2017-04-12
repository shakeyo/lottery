package model

import (
	"strconv"
)

type UserBaseInfo struct {
	NickName string // 昵称

	VIPLevel int    //vip等级
	Avatar   string //头像

	Rights uint64 //权限
}

type UserProperty struct {
	Money uint64 //余额
	Gold  uint64 //元宝
	Score uint64 //积分
}

type User struct {
	UID int64 // 用户id
	UserBaseInfo
	Property UserProperty
}

const (
	IS_GM = iota
	IS_Forbid
	IS_MUTE
)

func (u *User) String() string {
	return "[UID: " + strconv.FormatInt(u.UID, 10) +
		", NickName: " + u.NickName +
		", Money: " + strconv.FormatUint(u.Property.Money, 10) +
		", Gold: " + strconv.FormatUint(u.Property.Gold, 10) +
		", Score: " + strconv.FormatUint(u.Property.Score, 10) +
		"]"
}