package internal

import "sync"

type Group struct {
	gid     int64
	mutex   sync.Mutex
	members []int
}

func NewGroup(gid int64, members []int64) *Group {
	group := new(Group)
	group.gid = gid
	group.members = members
	return group
}

func (group *Group) Members() []int {
	return group.members
}

//修改成员，在副本修改，避免读取时的lock
func (group *Group) AddMember(uid int64) {
	group.mutex.Lock()
	defer group.mutex.Unlock()
	members := group.members.Clone()
	members.Add(uid)
	group.members = members
}

func (group *Group) RemoveMember(uid int64) {
	group.mutex.Lock()
	defer group.mutex.Unlock()
	members := group.members.Clone()
	members.Remove(uid)
	group.members = members
}

func (group *Group) IsMember(uid int64) bool {
	_, ok := group.members[uid]
	return ok
}

func (group *Group) IsEmpty() bool {
	return len(group.members) == 0
}
