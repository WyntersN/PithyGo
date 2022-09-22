/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-17 21:00:29
 * @LastEditTime: 2022-06-04 11:32:42
 * @FilePath: \PithyGo\addons\kuaibao\api.go
 */
package kuaibao

import (
	"errors"

	"github.com/bitly/go-simplejson"
)

type AddressInfo struct {
	Name     string //姓名
	Mobile   string //手机号码
	Province string //省
	City     string //市
	District string //区
	Detail   string //地址
}

func AutoAddressGit(address string) *simplejson.Json {

	json := requestPostGetAddress(address)
	j, _ := simplejson.NewJson(json)
	//jsonData, _ := j.MarshalJSON()
	return j

}

/**
 * @summary:获取电子面单
 * @description:
 * @param {string} shipper
 * @param {string} source
 * @param {string} orderns
 * @param {string} goods
 * @param {AddressInfo} recipient
 * @param {AddressInfo} sender
 * @Author: Wynters
 */
func GetWaybillCode(shipper string, source string, orderns string, goods string, recipient AddressInfo, sender AddressInfo) ([]byte, error) {

	json, err := requestPost("get.waybill.number", `
	{
		"cp_code": "`+shipper+`",
		"account_no": "xiaoyoutong",
		"account_password": "`+source+`",
		"order_data": [{
			"order_id": "`+orderns+`",
			"goods_name": "`+goods+`",
			"recipient": {
				"address": {
					"province": "`+recipient.Province+`",
					"district": "`+recipient.District+`",
					"city": "`+recipient.City+`",
					"detail": "`+recipient.Detail+`"
				},
				"mobile": "`+recipient.Mobile+`",
				"name": "`+recipient.Name+`"
			},
			"sender": {
				"address": {
					"province": "`+sender.Province+`",
					"district": "`+sender.District+`",
					"city": "`+sender.City+`",
					"detail": "`+sender.Detail+`"
				},
				"mobile": "`+sender.Mobile+`",
				"name": "`+sender.Name+`"
			}
		}]
	}
	`)

	if err != nil {
		return nil, err
	}

	j, _ := simplejson.NewJson(json)

	code, _ := j.Get("code").Int()
	if code == 0 {
		jsonData, _ := j.MarshalJSON()
		return jsonData, nil
	}
	msg, _ := j.Get("msg").String()
	return nil, errors.New(msg)

}

/**
 * @summary: 发送打印请求
 * @description:
 * @param {string} shipper
 * @param {string} source
 * @param {string} orderns
 * @param {string} waybill_code
 * @param {string} goods
 * @param {string} note
 * @param {string} routing_info
 * @param {AddressInfo} recipient
 * @param {AddressInfo} sender
 * @Author: Wynters
 */
func PrintWaybill(shipper string, source string, orderns string, waybill_code string, goods string, note string, routing_info string, recipient AddressInfo, sender AddressInfo) ([]byte, error) {

	json, err := requestPost("cloud.print.waybill", `
	{
		"sequence": "1/1",
		"print_type":3,
		"account_no": "xiaoyoutong",
		"agent_id": "`+source+`",
		"print_data": [{
			"cp_code": "`+shipper+`",
			"tid": "`+orderns+`",
			"waybill_code":"`+waybill_code+`",
			"goods_name": "`+goods+`",
			"note":"`+note+`",
			"routing_info":`+routing_info+`,
			"recipient": {
				"address": {
					"province": "`+recipient.Province+`",
					"district": "`+recipient.District+`",
					"city": "`+recipient.City+`",
					"detail": "`+recipient.Detail+`"
				},
				"mobile": "`+recipient.Mobile+`",
				"name": "`+recipient.Name+`"
			},
			"sender": {
				"address": {
					"province": "`+sender.Province+`",
					"district": "`+sender.District+`",
					"city": "`+sender.City+`",
					"detail": "`+sender.Detail+`"
				},
				"mobile": "`+sender.Mobile+`",
				"name": "`+sender.Name+`"
			}
		}]
	}
	`)

	if err != nil {
		return nil, err
	}

	j, _ := simplejson.NewJson(json)

	code, _ := j.Get("code").Int()
	if code == 0 {
		jsonData, _ := j.MarshalJSON()
		return jsonData, nil
	}
	msg, _ := j.Get("msg").String()
	return nil, errors.New(msg)

}

/**
 * @summary: 获取面单余量
 * @description:
 * @param {string} shipper
 * @param {string} source
 * @Author: Wynters
 */
func QueryBalance(shipper, source string) (int, error) {

	json, err := requestPost("account.waybill.balance", `
	{
        "customer_name": "xiaoyoutong",
        "shipper_type": "`+shipper+`",
        "customer_password": "`+source+`"
    }
	`)

	if err != nil {
		return -1, err
	}

	j, _ := simplejson.NewJson(json)

	code, _ := j.Get("code").Int()
	if code == 0 {
		count, _ := j.Get("data").Get("result").Get("count").Int()
		return count, nil
	}
	msg, _ := j.Get("msg").String()
	return -1, errors.New(msg)
}

func SelectExpressInfo(ecode, eun string) ([]byte, error) {

	json, err := requestPost("express.info.get", `
	{
        "exp_company_code": "`+ecode+`",
        "waybill_no": "`+eun+`"
    }
	`)

	if err != nil {
		return nil, err
	}

	j, _ := simplejson.NewJson(json)

	code, _ := j.Get("code").Int()
	if code == 0 {
		jsonData, _ := j.Get("data").GetIndex(0).MarshalJSON()
		return jsonData, nil
	}
	msg, _ := j.Get("msg").String()
	return nil, errors.New(msg)
}

func AutoAddressResolve(address string) ([]byte, error) {

	json, err := requestPost("cloud.address.resolve", `
	{
        "text": "`+address+`",
        "multimode":false,
        "customer_password":false
    }
	`)

	if err != nil {
		return nil, err
	}

	j, _ := simplejson.NewJson(json)

	code, _ := j.Get("code").Int()
	if code == 0 {
		jsonData, _ := j.MarshalJSON()
		return jsonData, nil
	}
	msg, _ := j.Get("msg").String()
	return nil, errors.New(msg)

}

// func AutoAddress(address string) ([]byte, error) {

// 	json, err := requestPost("cloud.address.cleanse", `
// 	{
//         "text": "`+address+`",
//         "multimode":false,
//         "customer_password":false
//     }
// 	`)

// 	if err != nil {
// 		return nil, err
// 	}

// 	j, _ := simplejson.NewJson(json)

// 	code, _ := j.Get("code").Int()
// 	if code == 0 {
// 		jsonData, _ := j.MarshalJSON()
// 		return jsonData, nil
// 	}
// 	msg, _ := j.Get("msg").String()
// 	return nil, errors.New(msg)

// }

// func AutoAddressEx(address string) ([]byte, error) {

// 	params := make(url.Values)

// 	params.Add("type", "address")
// 	params.Add("request_data", `{"data":{"text":"浙江省义乌市惠民路700号","multimode":false,"is_parse":0,"cleanTown":"0"}}`)
// 	request, _ := http.NewRequest(http.MethodPost, "https://open.kuaidihelp.com/Demo/DemoHandle/run", strings.NewReader(params.Encode()))
// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	response, err := curl.Send(request)

// 	if err != nil {
// 		return nil, err
// 	}
// 	resJson := strings.ReplaceAll(response.String(), "::", ":")
// 	j, err := simplejson.NewJson([]byte(resJson))
// 	if err != nil {
// 		println(err.Error(), "----------------", resJson)
// 	}

// 	code := j.Get("code").MustInt()
// 	if code == 0 {
// 		return []byte(resJson), nil
// 	}
// 	msg, _ := j.Get("msg").String()
// 	return nil, errors.New(msg)

// }
