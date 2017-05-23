package api

import (
	//"encoding/json"
	"io"
	"log"
	"net/http"
	//"server/model"
)

/*
	踢掉某个玩家
*/
func KickOffUser(w http.ResponseWriter, req *http.Request) {
	log.Print("processing KickOffUser")
	io.WriteString(w, "KickOffUser!\n")
}

/*
	踢掉所有玩家
*/
func KickOffAllUser(w http.ResponseWriter, req *http.Request) {
	log.Print("processing KickOffAllUser")
	io.WriteString(w, "KickOffAllUser!\n")
}

/*
	通知服务器玩家数据变更
*/
func SyncUserData(w http.ResponseWriter, req *http.Request) {
	log.Print("processing SyncUserData")
	io.WriteString(w, "SyncUserData!\n")
}

/*
	推送通知给玩家
	用于需要及时到达客户端的通知，比如充值到账，新邮件或自定义消息
	通知有区别与消息
*/
func PushNoticeToUser(w http.ResponseWriter, req *http.Request) {
	log.Print("processing PushNoticeToUser")
	io.WriteString(w, "PushNoticeToUser!\n")
}
