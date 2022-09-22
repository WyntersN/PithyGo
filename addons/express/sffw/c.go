/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-03-21 19:42:48
 * @LastEditTime: 2022-09-22 13:48:49
 * @FilePath: \PithyGo\addons\express\sffw\c.go
 */
package sffw

import (
	"PithyGo/addons/express/typestruct"
	"PithyGo/common"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/ebar-go/curl"
)

const (
	//apiUrl = "https://fns-edi.sit.sf-express.com:8020/" //测试环境
	apiUrl = "https://x5edi.fwx-network.com/"
	appKey = ""
)

func TRACE_QUERY(waybill_no string) (typestruct.TraceQueryStruct, error) {

	var info typestruct.TraceQueryStruct

	d, err := postReq("v2/trace/query", `{"mailNos":{"mailNo":["`+waybill_no+`"]}}`)

	if err == nil {
		a, _ := d.MarshalJSON()
		println(string(a), "---------------------------")
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

func TRACE_PLATFORM_SUBSCRIBE(waybill_no string) (*simplejson.Json, error) {
	return postReq("v1/trace/register", `{"mailNo":"`+waybill_no+`"}`)
}

func postReq(route, body string) (*simplejson.Json, error) {
	var sign, timestamp = sign(body)
	request, err := http.NewRequest(http.MethodPost, apiUrl+route, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("appKey", appKey)
	request.Header.Add("timestamp", timestamp)
	request.Header.Add("sign", sign)
	request.Header.Add("format", "json")
	request.Header.Add("v", "1.0")
	response, err := curl.Send(request)
	if err != nil {
		return nil, err
	}
	//println(response.String(), sign, appKey, timestamp, sign, body)
	//req, _ := curl.Post(host, strings.NewReader(params.Encode()))
	j, _ := simplejson.NewJson(response.Byte())
	return j, nil
}

// sign=base64(md5(请求body+签名appKey+timestamp))
func sign(body string) (string, string) {
	var (
		timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
		sign      = common.Md5Byte(body + appKey + timestamp)
	)
	return base64.StdEncoding.EncodeToString(sign), timestamp

}
