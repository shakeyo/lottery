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

const (
	ErrorCode_Unknown         = 500
	ErrorCode_OK              = 200
	ErrorCode_InvaildToken    = 302
	ErrorCode_VersionLimit    = 201
	ErrorCode_InvaildData     = 414
	ErrorCode_AccountForbid   = 422
	ErrorCode_InvaildProtocol = 509
)

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
