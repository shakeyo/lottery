package api

import (
	//"encoding/json"
	"io"
	"log"
	"net/http"
	//"server/model"
)

/*
	发送系统消息
	用于广播消息等（比如跑马灯，比如封盘，开奖等）
*/
func SendSystemMessage(w http.ResponseWriter, req *http.Request) {
	log.Print("processing SendSystemMessage")
	io.WriteString(w, "SendSystemMessage!\n")
}

/*
	设置服务器状态

*/
func SetServiceStatus(w http.ResponseWriter, req *http.Request) {
	log.Print("processing SetServiceStatus")
	io.WriteString(w, "SetServiceStatus!\n")
}
