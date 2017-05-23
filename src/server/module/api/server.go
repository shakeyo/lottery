package api

import (
	"log"
	"net/http"
)

func init() {
	log.Print("api server startup")
	http.HandleFunc("/user", KickOffUser)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
