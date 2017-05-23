/*
聊天室部分
*/

package msg

func init() {
	Processor.Register(&C2F_JoinChatRoom{})
	Processor.Register(&C2F_LeaveChatRoom{})
	Processor.Register(&C2F_CreateChatRoom{})
}

//用户加入频道
type C2F_JoinChatRoom struct {
	ID string
}

//用户申请加入频道返回
type F2C_JoinChatRoom_Ack struct {
	Code   int
	ErrStr string
}

//用户退出频道(所有人退出则自动解散）
type C2F_LeaveChatRoom struct {
	ID string
}

//退出频道返回
type F2C_LeaveChatRoom_Ack struct {
	Code   int
	ErrStr string
}

//创建频道(特定管理用户）
type C2F_CreateChatRoom struct {
	Name        string
	Description string
}

//创建返回
type F2C_CreateChatRoom_Ack struct {
	Code   int
	ErrStr string
}
