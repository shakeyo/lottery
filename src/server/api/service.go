package api

import (
	"server/model"
	//	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	//	"bytes"
	"encoding/json"
)

type APIConfig struct {
	HostUrl     string
	Version     string
	OAuthEnable bool
}

type serviceAPI struct {
	cfg   APIConfig
	token model.AuthToken
}

func NewUserAPI() *UserAPI {
	return nil
}

func NewSystemAPI() *SystemAPI {
	return nil
}

func (self *serviceAPI) AuthUser(userID int64, userToken string) (*model.User, int) {

	v := url.Values{}
	v.Set("UserID", strconv.FormatInt(userID, 10))
	v.Set("Token", userToken)
	v.Set("Version", self.cfg.Version)

	resp, err := http.PostForm(self.cfg.HostUrl+"/Authenticate", v)

	if err != nil {
		fmt.Println("http post error:", err)
		return nil, 1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http io error:", err)
		return nil, 1
	}
	fmt.Println(string(body))

	type RespData struct {
		Code int
		Msg  string
		Data string
	}

	var ret RespData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return nil, 1
	}

	return nil, 0
}

func (self *serviceAPI) ModifyUserProperty(userID int64, token string) bool {
	return false
}
