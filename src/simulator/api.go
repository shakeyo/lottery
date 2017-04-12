package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/model"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Print("processing HelloServer")
	io.WriteString(w, "hello, world!\n")
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	log.Print("processing Authenticate")
	log.Print(req.PostForm.Encode())
	log.Print(req.Form.Encode())

	user := &model.User{
		UID: 1000,
		Property: model.UserProperty{
			Money: 1000,
			Gold:  1000,
			Score: 10000000,
		},
		UserBaseInfo: model.UserBaseInfo{
			NickName: "ceshi",
			VIPLevel: 1,
			Avatar:   "",
			Rights:   0,
		},
	}

	data, _ := json.Marshal(user)
	io.WriteString(w, string(data))

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

}

func main() {
	log.Print("startup")
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/authenticate", Authenticate)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
