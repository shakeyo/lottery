package msg

import (
	"github.com/name5566/leaf/network/json"
)

var (
	Processor = json.NewProcessor()
)

func init() {
	Processor.Register(&C2W_Auth{})
	Processor.Register(&W2C_Auth_Ack{})
	Processor.Register(&C2W_Heartbeat{})
	Processor.Register(&W2C_Heartbeat{})
}

type C2W_Heartbeat struct {
}

type W2C_Heartbeat struct {
}

type C2W_Auth struct {
	UID       int64
	Timestamp int64
	Nonce     string
	Signature string
	Token     string
}

type W2C_Auth_Ack struct {
	Code int
}
