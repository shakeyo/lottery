package internal

import (
	"reflect"
	"server/model"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var (
	//临时用
	msgList []*msg.ChatMessage
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2F_JoinRoom{}, handleJoinRoom)
	handleMsg(&msg.C2F_QuitRoom{}, handleQuitRoom)
	handleMsg(&msg.C2F_SendMsg{}, handleSendMsg)

	skeleton.RegisterChanRPC("UserLogOff", rpcUserLogOff)
	skeleton.RegisterChanRPC("UserLogged", rpcUserLogged)
}

func rpcUserLogOff(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	user := a.UserData().(*model.User)

	/*	channel.Broadcast(&msg.F2C_MemberChangedNotify{
		Name:     user.NickName,
		RoomName: m.RoomName,
		Event:    "Leave",
	})*/
}

func rpcUserLogged(args []interface{}) {

}

func handleJoinRoom(args []interface{}) {
	m := args[0].(*msg.C2F_JoinRoom)
	a := args[1].(gate.Agent)

	sendMsg := &msg.F2C_JoinRoom_Ack{}
	sendAckFunc := func(err string) {
		sendMsg.Err = err
		sendMsg.RoomID = m.RoomID
		a.WriteMsg(sendMsg)
	}

	if a.UserData() == nil {
		sendAckFunc("forbid")
		return
	}

	user := a.UserData().(*model.User)

	log.Debug("User:%v Join room:%v", user.UID, m.RoomID)

	channel, ret := ChannelServiceInstance.Channel(m.RoomID)
	if ret == false {
		channel = ChannelServiceInstance.NewChannel(m.RoomID)
	}
	channel.Add(&a)

	sendMsg.Members = channel.Members()
	a.WriteMsg(&sendMsg)

	channel.Broadcast(&msg.F2C_ChannelNotify{
		RoomID: m.RoomID,
		Event:  "Enter",
		Member: msg.ChatMember{
			UserID: user.UID,
			Name:   user.NickName,
			Status: 0,
		},
	})

	a.WriteMsg(&msg.F2C_MsgList{
		MsgList: msgList,
	})

}

func handleQuitRoom(args []interface{}) {
	m := args[0].(*msg.C2F_QuitRoom)
	a := args[1].(gate.Agent)

	sendMsg := &msg.F2C_QuitRoom_Ack{}
	sendAckFunc := func(err string) {
		a.WriteMsg(sendMsg)
	}

	if a.UserData() == nil {
		sendAckFunc("forbid")
		return
	}

	user := a.UserData().(*model.User)

	log.Debug("User:%v quit room:%v", user.UID, m.RoomID)

	channel, ret := ChannelServiceInstance.Channel(m.RoomID)
	if ret == false {
		sendAckFunc("unexpected")
		return
	}
	channel.Leave(user.UID)
	channel.Broadcast(&msg.F2C_ChannelNotify{
		RoomID: m.RoomID,
		Event:  "Leave",
		Member: msg.ChatMember{
			UserID: user.UID,
			Name:   user.NickName,
			Status: 0,
		},
	})
	sendAckFunc("ok")
}

func handleSendMsg(args []interface{}) {
	m := args[0].(*msg.C2F_SendMsg)
	a := args[1].(gate.Agent)

	sendMsg := &msg.F2C_SendMsg_Ack{}
	sendAckFunc := func(err string) {
		sendMsg.Err = err
		a.WriteMsg(sendMsg)
	}

	if a.UserData() == nil {
		sendAckFunc("forbid")
		return
	}

	user := a.UserData().(*model.User)

	log.Debug("User:%v say:%v type:%v target:%v", user.UID, m.Message, m.Conversation.Type, m.Conversation.Target)
	channel, ret := ChannelServiceInstance.Channel(m.Conversation.Target)
	if ret == false {
		sendAckFunc("no exists room")
		return
	}

	msgList = append(msgList, &msg.ChatMessage{
		Conversation: msg.ChatConversation{
			Target: m.Conversation.Target,
			Type:   m.Conversation.Type,
		},
		Sender:     user.UID,
		SentTime:   time.Now().Unix(),
		MsgContent: m.Message,
	})

	channel.Broadcast(&msg.F2C_MsgList{
		MsgList: msgList[len(msgList)-1:],
	})

	sendAckFunc("ok")
}
