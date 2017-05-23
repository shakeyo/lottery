package msg

const (
	IM_Tips_MemberEnter = iota
	IM_Tips_MemberLeave
	IM_Tips_SetAdmin
	IM_Tips_UnSetAdmin
	IM_Tips_KickMember
	IM_Tips_Mute
	IM_Tips_UnMute
	IM_Tips_
)
const (
	IM_Conversation_Private = iota
	IM_Conversation_Group
	IM_Conversation_ChatRoom
	IM_Conversation_System
	IM_Conversation_Service
)

const (
	//已投递
	IM_MsgStatus_Delivered = 0x0000
	//没有目标
	IM_MsgStatus_Rejected_NoTarget = 0x0001
	//频率过快
	IM_MsgStatus_Rejected_LimitExceeded = 0x0002
	//文字被过滤(非法字符）
	IM_MsgStatus_Rejected_Filtered = 0x0004
)

const (
	IM_Feature_PhizMsg  = 0x0000
	IM_Feature_PropMsg  = 0x0001
	IM_Feature_VoiceMsg = 0x0002
	IM_Feature_VOIPCall = 0x0004
)

const (
	//满人
	IM_GroupErrCode_FullMember = 800 + iota
	//没有权限操作
	IM_GroupErrCode_NoPermission
	//目标不存在或者不在群里
	IM_GroupErrCode_NoTarget
	//超过最大创建数量
	IM_GroupErrCode_MaxCreated
)

const (

	//圈子是否启用验证
	IM_GroupSettings_Authentication = 0x01
	//GroupSettings_Rookie = 0x01
)

const (
	//同意加入圈子
	IM_GroupApplyOpinion_Agree = iota
	//拒绝加入
	IM_GroupApplyOpinion_Reject
	//忽略申请
	IM_GroupApplyOpinion_Ignore
)

const (
	//公共红包
	IM_MoneyMsg_Scope_Public = iota
	//定向红包
	IM_MoneyMsg_Scope_Specific
)
