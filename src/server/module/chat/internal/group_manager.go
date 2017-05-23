package internal

type GroupObserver interface {
	OnGroupMemberAdd(g *Group, uid int64)
	OnGroupMemberRemove(g *Group, uid int64)
}

type GroupManager struct {
	mutex    sync.Mutex
	groups   map[int64]*Group
	observer GroupObserver
	ping     string
}

func NewGroupManager() *GroupManager {
	now := time.Now().Unix()
	r := fmt.Sprintf("ping_%d", now)
	for i := 0; i < 4; i++ {
		n := rand.Int31n(26)
		r = r + string('a'+n)
	}

	m := new(GroupManager)
	m.groups = make(map[int64]*Group)
	m.ping = r
	return m
}

func (self *GroupManager) GetGroups() []*Group {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	groups := make([]*Group, 0, len(self.groups))
	for _, group := range self.groups {
		groups = append(groups, group)
	}
	return groups
}

func (self *GroupManager) FindGroup(gid int64) *Group {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if group, ok := self.groups[gid]; ok {
		return group
	}
	return nil
}

func (self *GroupManager) FindUserGroups(appid int64, uid int64) []*Group {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	groups := make([]*Group, 0, 4)
	for _, group := range self.groups {
		if group.appid == appid && group.IsMember(uid) {
			groups = append(groups, group)
		}
	}
	return groups
}

func (s *GroupManager) CreateGroup(appid int64, uid int64) []*Group {

}

func (s *GroupManager) Disband(appid int64, uid int64) []*Group {

}

func (s *GroupManager) AddMember(appid int64, uid int64) []*Group {

}

func (s *GroupManager) RemoveMember(appid int64, uid int64) []*Group {

}
