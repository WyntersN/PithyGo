/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-12-19 17:19:30
 * @LastEditTime: 2022-09-22 13:49:09
 * @FilePath: \PithyGo\addons\express\best\c.go
 */
package best

import (
	"PithyGo/addons/express/typestruct"
	"PithyGo/common"
	"net/http"
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/ebar-go/curl"
)

const (
	apiUrl     = "http://edi-q9.ns.800best.com/kd/api/process"
	partnerID  = ""
	partnerKey = ""
)

func TRACE_QUERY(waybill_no string) (typestruct.TraceQueryStruct, error) {
	var info typestruct.TraceQueryStruct
	d, err := postReq("KD_TRACE_QUERY", `{"mailNos":{"mailNo":["`+waybill_no+`"]}}`)
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

	return info, nil
}

func TRACE_PLATFORM_SUBSCRIBE(waybill_no string) (*simplejson.Json, error) {
	return postReq("KD_SCAN_REGISTER_NOTIFY", `{"acceptCityName":"杭州市","acceptCountyName":"西湖区","acceptManAddress":"教工路1号","acceptManName":"xx","acceptManPhone":"138xxxx","acceptProvinceName":"浙江省","mailNo":"`+waybill_no+`","sendCityName":"杭州市","sendCountyName":"西湖区","sendManName":"xx","sendProvinceName":"浙江省"}`)
}

func TRACE_ADDRESS_ANOMALY(provinceName, cityName, countyName, address string) (*simplejson.Json, error) {
	return postReq("KD_WAYBILL_ADDRESS_ANOMALY", `{"provinceName":"`+provinceName+`","cityName":"`+cityName+`","countyName":"`+countyName+`","address":"`+address+`"}`)
}

func postReq(serviceType, bizData string) (*simplejson.Json, error) {
	params := make(url.Values)
	params.Set("sign", sign(bizData))
	params.Set("partnerID", partnerID)
	params.Set("serviceType", serviceType)
	params.Set("bizData", bizData)

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
	//	println(string(response.Byte()))
	//req, _ := curl.Post(host, strings.NewReader(params.Encode()))
	j, _ := simplejson.NewJson(response.Byte())
	return j, nil
}

func sign(bizData string) string {
	return common.Md5(bizData + partnerKey)
}
