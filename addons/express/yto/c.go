/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-03-30 16:12:41
 * @LastEditTime: 2022-09-22 13:49:03
 * @FilePath: \PithyGo\addons\express\yto\c.go
 */
package yto

import (
	"PithyGo/addons/express/typestruct"
	"PithyGo/common"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/ebar-go/curl"
)

const (
	apiUrl    = "https://openapi.yto.net.cn:11443/open/"
	client_id = ""
	secretKey = ""
)

func data_sign(data string) string {
	return base64.StdEncoding.EncodeToString(common.Md5Byte(data + secretKey))
}

/**
 * @summary: 圆通快递轨迹订阅
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
	return postReq("subscribe_adapter", `{"client_id":"`+client_id+`","logistics_interface":"{\"clientId\":\"`+client_id+`\",\"waybillNo\":\"`+waybillNo+`\"}","msg_type":"online"}`)
}

/**
 * @summary: 快递查询
 * @description:
 * @param {string} ordersn
 * @Author: Wynters
 */
func TRACE_QUERY(ordersn string) (typestruct.TraceQueryStruct, error) {
	d, err := postReq("track_query_adapter", `{"NUMBER":"`+ordersn+`"}`)
	var info typestruct.TraceQueryStruct

	if err == nil {
		a, _ := d.MarshalJSON()
		println(string(a), "--------------AAAAAA-------------")
	}

	/*








		if err != nil {
			return info, err
		}
		j := d.Get("traceLogs").GetIndex(0)
		info.Brand = "HTKY"
		info.Order = "DESC"
		info.No = j.Get("mailNo").MustString()
		dArry, _ := j.GetPath("traces", "trace").Array()

		for _, v := range dArry {
			var dataInfo typestruct.TraceQueryStructData
			data := v.(map[string]interface{})

			dataInfo.Context = data["remark"].(string)
			dataInfo.Status = data["scanType"].(string)
			dataInfo.Time = data["acceptTime"].(string)

			info.Data = append(info.Data, dataInfo)
		}
	*/
	return info, nil
}
func postReq(method, param string) (*simplejson.Json, error) {

	var (
		timestamp    = strconv.FormatInt(time.Now().UnixMilli(), 10)
		paramJson, _ = json.Marshal(param)
		postData     = `{"timestamp":"` + timestamp + `","param":` + string(paramJson) + `,"sign":"` + data_sign(param+method+"v1") + `","format":"JSON"}`
	)

	response, err := curl.PostJson(apiUrl+method+"/v1/kCxMfO/"+client_id+"_STD", strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	//println(response.String(), postData, apiUrl+"/"+method+"/v1/kCxMfO/"+client_id)
	//req, _ := curl.Post(host, strings.NewReader(params.Encode()))
	j, _ := simplejson.NewJson(response.Byte())
	return j, nil
}
