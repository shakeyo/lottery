package msg

func init() {
	Processor.Register(&F2C_RoomNotify{})
	Processor.Register(&C2F_JoinRoom{})
	Processor.Register(&F2C_JoinRoom_Ack{})
	Processor.Register(&C2F_QuitRoom{})
	Processor.Register(&F2C_QuitRoom_Ack{})
	Processor.Register(&C2F_SendMsg{})
	Processor.Register(&F2C_SendMsg_Ack{})
	Processor.Register(&F2C_MsgList{})
}

/*聊天服务器*/

const (
	ChatConversation_Private = iota
	ChatConversation_Group
	ChatConversation_ChatRoom
	ChatConversation_System
	ChatConversation_CustomerService
)

type ChatConversation struct {
	Type   int
	Target string
}

type ChatMessage struct {
	Conversation ChatConversation
	Sender       int64
	SentTime     int64
	MsgContent   string
}

type ChatMember struct {
	UserID int64
	Name   string
	Status uint64
}

type C2F_JoinRoom struct {
	RoomID string
}

type F2C_JoinRoom_Ack struct {
	Err     string
	RoomID  string
	MsgList []*ChatMessage
	//Members  []*ChatMember
	Members []int64
}

type F2C_RoomNotify struct {
	Event  string
	RoomID string
	Member ChatMember
}

type C2F_QuitRoom struct {
	RoomID string
}

type F2C_QuitRoom_Ack struct {
}

type C2F_SendMsg struct {
	Conversation ChatConversation
	Message      string
}

type F2C_SendMsg_Ack struct {
	Err string
}

type C2F_FetchHistoryMessage struct {
	Conversation ChatConversation
}

type F2C_MsgList struct {
	MsgList []*ChatMessage
}
