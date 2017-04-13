package internal

import (
	"reflect"
	"server/modules/world/internal/remote_api"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var userAPI = remote_api.NewUserAPI()

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2W_Auth{}, handleAuth)
	handleMsg(&msg.C2W_Heartbeat{}, handleHeartBeat)
}

func handleHeartBeat(args []interface{}) {
	//m := args[0].(*msg.C2W_Heartbeat)
	a := args[1].(gate.Agent)
	a.WriteMsg(&msg.W2C_Heartbeat{})
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.C2W_Auth)
	a := args[1].(gate.Agent)

	log.Debug("user auth:%v", m.UID)

	sendMsg := &msg.W2C_Auth_Ack{}
	sendAckFunc := func(err int) {
		sendMsg.Code = err
		a.WriteMsg(sendMsg)
	}

	if a.UserData() != nil {
		log.Error("repeated auth attempted, User:%v", m.UID)
		sendAckFunc(-1)
		return
	}

	user, ret := (*userAPI).AuthUser(m.UID, m.Token)
	if ret != 0 {
		sendAckFunc(ret)
	}

	sendAckFunc(0)

	a.SetUserData(user)

	log.Error("user logged, User:%v", a.UserData())
}
