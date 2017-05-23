/*
圈子部分
*/

package msg

func init() {
	//Processor.Register(&C2F_JoinChatRoom{})
}

//圈子
type ChatGroupBrief struct {
	ID          string
	Name        string
	Head        string
	Description string
}

//玩家圈子
type UserChatGroup struct {
	ID        string
	Profile   int
	EnterTime int64
	Rights    int64
}

//用户创建圈子
type C2F_CreateGroup struct {
	//Group ChatGroup
}

//创建圈子返回
type F2C_CreateGroup_Ack struct {
	Code int
}

//用户加入圈子
type C2F_JoinGroup struct {
	ID        string
	Introduce string
	Password  string
}

//用户申请加入频道返回
type F2C_JoinGroup_Ack struct {
	ID int
}

//用户退出圈子(群主则为解散）
type C2F_QuitGroup struct {
	ID string
}

//退出圈子返回
type F2C_QuitGroup_Ack struct {
	Code int
}

//加入圈子申请通知
type F2C_GroupJoin_Notify struct {
	ID        string
	Applier   string
	ApplierID int64
	Introduce string
}

//处理申请
type C2F_DealGroupApply struct {
	//group10:user1001
	ApplyID string
	Opinion int
	Notes   string
}

//处理申请返回
type F2C_DealGroupApply_Ack struct {
	Code int
}

//从圈子里踢掉
type C2F_DelGroupMembers struct {
	ID      string
	Members []int
}

//圈子踢人返回
type F2C_DelGroupMembers_Ack struct {
	Code int
}

type F2C_RemoveFromGroup struct {
	ID string
}

//搜索返回
type F2C_RemoveFromGroup_Ack struct {
	Code int
}

//圈子设置变更（名称，头像，简介，权限，管理员等）
type F2C_GroupChanged_Notify struct {
}

//根据关键字搜索圈子
type C2F_SearchGroups struct {
	Keywords string
}

//搜索返回
type F2C_SearchGroups_Ack struct {
	Code    int
	Results []ChatGroupBrief
}

//设置成员角色
type C2F_SetGroupMemberRole struct {
	ID     string
	Member int64
	Have   bool
}

//设置角色返回
type F2C_SetGroupMemberRole_Ack struct {
	Code int
}

//修改圈子信息
type C2F_UpdateGroupInfo struct {
	ID     string
	Member int64
	Have   bool
}

//修改圈子信息返回
type F2C_UpdateGroupInfo_Ack struct {
	Code int
}
