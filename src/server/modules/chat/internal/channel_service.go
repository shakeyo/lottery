package internal

import (
	"sync"
)

var ChannelServiceInstance = makeChannelService()

type ChannelService struct {
	channels map[string]*Channel // all server channels
	sync.RWMutex
}

func makeChannelService() *ChannelService {
	return &ChannelService{
		channels: make(map[string]*Channel),
	}
}

func (c *ChannelService) NewChannel(name string) *Channel {
	c.Lock()
	defer c.Unlock()

	channel := newChannel(name, c)
	c.channels[name] = channel
	return channel
}

// Get channel by channel name
func (c *ChannelService) Channel(name string) (*Channel, bool) {
	c.RLock()
	defer c.RUnlock()

	channel, ok := c.channels[name]
	return channel, ok
}

// Get all members in channel by channel name
func (c *ChannelService) Members(name string) []int64 {
	c.RLock()
	defer c.RUnlock()

	if channel, ok := c.channels[name]; ok {
		return channel.Members()
	}
	return make([]int64, 0)
}

// Destroy channel by channel name
func (c *ChannelService) DestroyChannel(name string) {
	c.RLock()
	c.RUnlock()

	if channel, ok := c.channels[name]; ok {
		channel.Destroy()
	}
}
