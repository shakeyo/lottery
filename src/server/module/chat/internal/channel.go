package internal

import (
	//	"reflect"
	"sync"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type SessionFilter func(*gate.Agent) bool

type Channel struct {
	sync.RWMutex
	name    string                // channel name
	uidMap  map[int64]*gate.Agent // uid map to session pointer
	members []int64               // all user ids
	service *ChannelService       // channel service which contain current channel
}

func newChannel(n string, cs *ChannelService) *Channel {
	return &Channel{
		name:    n,
		service: cs,
		uidMap:  make(map[int64]*gate.Agent)}
}

func (c *Channel) Member(uid int64) *gate.Agent {
	c.RLock()
	defer c.RUnlock()

	return c.uidMap[uid]
}

func (c *Channel) Members() []int64 {
	c.RLock()
	defer c.RUnlock()

	return c.members
}

// Push message to partial client, which filter return true
func (c *Channel) Multicast(v interface{}, filter SessionFilter) error {

	log.Debug("Type=Multicast Data=%+v", v)

	c.RLock()
	defer c.RUnlock()

	for _, s := range c.uidMap {
		if !filter(s) {
			continue
		}
		(*s).WriteMsg(v)
	}

	return nil
}

// Push message to all client
func (c *Channel) Broadcast(v interface{}) error {

	log.Debug("Type=Broadcast Data=%+v", v)

	c.RLock()
	defer c.RUnlock()

	for _, s := range c.uidMap {
		(*s).WriteMsg(v)
	}

	return nil

}

func (c *Channel) IsContainUser(uid int64) bool {
	c.RLock()
	defer c.RUnlock()

	if _, ok := c.uidMap[uid]; ok {
		return true
	}

	return false
}

func (c *Channel) IsContainAgent(session *gate.Agent) bool {
	c.RLock()
	defer c.RUnlock()

	for _, s := range c.uidMap {
		if s == session {
			return true
		}
	}

	return false
}

func (c *Channel) Add(session *gate.Agent, uid int64) {
	c.Lock()
	defer c.Unlock()

	c.uidMap[uid] = session
	c.members = append(c.members, uid)
}

func (c *Channel) Leave(uid int64) {
	if !c.IsContainUser(uid) {
		return
	}

	c.Lock()
	defer c.Unlock()

	var temp []int64
	for i, u := range c.members {
		if u == uid {
			temp = append(temp, c.members[:i]...)
			c.members = append(temp, c.members[(i+1):]...)
			break
		}
	}
	delete(c.uidMap, uid)
}

func (c *Channel) LeaveAll() {
	c.Lock()
	defer c.Unlock()

	c.uidMap = make(map[int64]*gate.Agent)
	c.members = make([]int64, 0)
}

func (c *Channel) Count() int {
	c.RLock()
	defer c.RUnlock()

	return len(c.uidMap)
}

func (c *Channel) Destroy() {
	c.service.Lock()
	defer c.service.Unlock()

	delete(c.service.channels, c.name)
}
