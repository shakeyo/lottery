package model

type Session struct {
	SID int64
	User
	Vaild      bool
	LastActive uint64
}

func (self *Session) GetUID() int64 {
	return self.UID
}
