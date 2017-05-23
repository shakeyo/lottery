package msg

func init() {
	Processor.Register(&C2F_SendMsg{})
	Processor.Register(&F2C_SendMsg_Ack{})
}

/*聊天公共*/

//文本(表情）消息
type TextMessage struct {
	Msg string
	Ext string
}

//命令消息
type CommandMessage struct {
	Action string
	Args   []string
}

//图片消息
type ImageMessage struct {
	Thumb    string
	FileUri  string
	FileSize int
	Width    int
	Height   int
}

//红包消息内容
type MoneyMessage struct {
	//定向红包，还是公共红包
	Scope int
	//红包金额（定向是固定金额，公共红包是总额）
	Money int
	//公共红包有效，红包数量
	Amount int
	//固定红包有效，接收者
	Receivers []int
}

//互动道具消息内容
type PropMessage struct {
	PropID  int
	Payload string
}

//礼物赠送消息内容
type GiftMessage struct {
	GiftID  int
	Payload string
}

//消息
type Message struct {
	MsgID      int64
	Body       string
	From       string
	To         string
	CreateTime int64
	Status     int
}

//发送消息
type C2F_SendMsg struct {
	ID   string
	Type string
	To   string
	Body string
}

//发送返回
type F2C_SendMsg_Ack struct {
	Code int
}

//新消息
type F2C_NewMsgList struct {
	MsgList []Message
}
