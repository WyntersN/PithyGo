/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-12-08 06:43:39
 * @LastEditTime: 2022-09-22 13:49:18
 * @FilePath: \PithyGo\addons\baidu\api.go
 */
package baidu

import (
	"PithyGo/service"
	"net/url"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/ebar-go/curl"
)

func getToken() string {
	tokenTmp, _ := service.GetCache("baidu_access_token")
	//server.LogInfo.Println(tokenTmp)
	if tokenTmp != nil {
		//	println("------------------------------11111111111111----------------------------")
		return tokenTmp.(string)
	}
	//println("------------------------------2222222222222----------------------------")
	type jsonData struct {
		AccessToken   string `json:"access_token"`
		ExpiresIn     int64  `json:"expires_in"`
		RefreshToken  string `json:"refresh_token"`
		Scope         string `json:"scope"`
		SessionKey    string `json:"session_key"`
		SessionSecret string `json:"session_secret"`
	}
	var (
		reqData       jsonData
		token         string
		client_id     = ""
		client_secret = ""
	)

	req, _ := curl.Get("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + client_id + "&client_secret=" + client_secret)

	req.BindJson(&reqData)

	//println(string(req.Byte()))

	if reqData.AccessToken != "" {
		token = reqData.AccessToken
		service.SetCacheDefault("baidu_access_token", token, time.Second*time.Duration(reqData.ExpiresIn))
		//return token
	}
	return token

}

func AutoAddress(address string) *simplejson.Json {
	var ( //24.ce17cce5b1be48c1d0a4c694d3bc420f.2592000.1642119202.282335-23170253
		req, _ = curl.PostJson("https://aip.baidubce.com/rpc/2.0/nlp/v1/address?access_token="+getToken(), strings.NewReader(`{"text":"`+address+`"}`))
		j, _   = simplejson.NewJson(req.Byte())
	)
	return j
}

func AutoOcr(image string) string {
	params := make(url.Values)
	params.Set("image", image)
	var (
		request, _ = curl.Post("https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic?access_token="+getToken(), strings.NewReader(params.Encode()))
		j, _       = simplejson.NewJson(request.Byte())
		str        = j.Get("words_result").GetIndex(0).Get("words").MustString()
	)

	return str

}
