package api

import (
	//	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	//	"bytes"
	"encoding/json"

	"github.com/name5566/leaf/log"
)

type service struct {
	cfg RemoteAPICfg
}

type serviceRespData struct {
	Code int
	Msg  string
	Data string
}

func NewUserAPI() *UserAPI {
	return nil
}

func NewSystemAPI() *SystemAPI {
	return nil
}

func (self *service) makePostForm(userID int64, userToken string) url.Values {
	v := url.Values{}
	v.Set("UserID", strconv.FormatInt(userID, 10))
	v.Set("Token", userToken)
	v.Set("Version", self.cfg.Version)
	return v
}

func (self *service) doHttpPost(path string, form url.Values) (string, int) {

	url := fmt.Sprintf("%s/%s", self.cfg.HostUrl, path)
	resp, err := http.PostForm(url, form)
	if err != nil {
		log.Error("http post error:%v", err)
		return "", -1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("http io error:%v", err)
		return "", -1
	}
	log.Debug(string(body))

	var ret serviceRespData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Error("json unmarshal error:%v", err)
		return "", -1
	}

	if ret.Code != 0 {
		log.Error("request error,code:%v msg:%v", ret.Code, ret.Msg)
		return "", ret.Code
	}

	return ret.Data, 0
}

func (self *service) Authenticate(userID int64, userToken string) (string, int) {

	v := self.makePostForm(userID, userToken)

	resp, err := http.PostForm(self.cfg.HostUrl+"/Authenticate", v)
	if err != nil {
		log.Error("http post error:%v", err)
		return "", -1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("http io error:%v", err)
		return "", -1
	}
	log.Debug(string(body))

	var ret serviceRespData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Error("json unmarshal error:%v", err)
		return "", -1
	}

	if ret.Code != 0 {
		log.Error("request error,code:%v msg:%v", ret.Code, ret.Msg)
		return "", ret.Code
	}

	return ret.Data, 0
}

func (self *service) ModifyProperty(userID int64, token string) bool {
	return false
}
