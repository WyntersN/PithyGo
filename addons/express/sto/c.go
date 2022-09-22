/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-12-17 15:17:42
 * @LastEditTime: 2022-09-22 13:48:57
 * @FilePath: \PithyGo\addons\express\sto\c.go
 */
package sto

import (
	"PithyGo/common"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/ebar-go/curl"
)

const (
	apiUrl    = "https://cloudinter-linkgatewayonline.sto.cn/gateway/link.do"
	appKey    = ""
	secretKey = ""
	from_code = ""
)

func data_digest(content string) string {
	return base64.StdEncoding.EncodeToString(common.Md5Byte(content + secretKey))
}

/**
 * @summary: 申通快递轨迹订阅
 * @description:
 * @param {string} ordersn
 * @Author: Wynters
 */

// type subscribeInfoList struct {
// 	SubscribeInfoList []SubscribeInfo `json:"subscribeInfoList"`
// }

func TRACE_PLATFORM_SUBSCRIBE(waybillNo string) (*simplejson.Json, error) {

	// var subscribeInfoList subscribeInfoList
	// subscribeInfoList.SubscribeInfoList = SubscribeInfo
	return postReq("STO_TRACE_PLATFORM_SUBSCRIBE", `{"subscribeInfoList":[{"waybillNo":"`+waybillNo+`"}]}`, "sto_trace_platform", "sto_trace_platform")
}

/**
 * @summary: 申通快递查询
 * @description:
 * @param {string} ordersn
 * @Author: Wynters
 */
func TRACE_QUERY_COMMON(ordersn string) {

}
func postReq(api_name, content, to_appkey, to_code string) (*simplejson.Json, error) {
	params := make(url.Values)
	params.Set("content", content)
	params.Set("data_digest", data_digest(content))
	params.Set("api_name", api_name)
	params.Set("from_appkey", appKey)
	params.Set("from_code", from_code)
	params.Set("to_appkey", to_appkey)
	params.Set("to_code", to_code)

	request, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	//println(params.Encode())

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := curl.Send(request)
	if err != nil {
		return nil, err
	}
	//req, _ := curl.Post(host, strings.NewReader(params.Encode()))
	j, _ := simplejson.NewJson(response.Byte())
	return j, nil
}
