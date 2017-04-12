package gate

import (
	"server/modules/chat"
	//"server/game"
	"server/modules/world"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2W_Auth{}, world.ChanRPC)
	msg.Processor.SetRouter(&msg.C2F_JoinRoom{}, chat.ChanRPC)
	msg.Processor.SetRouter(&msg.C2F_SendMsg{}, chat.ChanRPC)
	msg.Processor.SetRouter(&msg.C2F_QuitRoom{}, chat.ChanRPC)
	//msg.Processor.SetRouter(&msg.FireActionReq{}, game.ChanRPC)
}
