package api

import (
	//	"encoding/xml"
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

func (self *service) AuthUser(userID int64, userToken string) (string, int) {

	v := url.Values{}
	v.Set("UserID", strconv.FormatInt(userID, 10))
	v.Set("Token", userToken)
	v.Set("Version", self.cfg.Version)

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

func (self *service) ModifyUserProperty(userID int64, token string) bool {
	return false
}
