/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-17 23:31:50
 * @LastEditTime: 2022-08-24 14:24:29
 * @FilePath: \PithyGo\addons\kuaibao\core.go
 */
package kuaibao

import (
	"PithyGo/common"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ebar-go/curl"
)

const (
	apiUrl = "https://kop.kuaidihelp.com/api"
	appId  = "111719"
	appKey = "ef556fa73aded5259d2603258d4d975068ca4358"
)

func requestPost(method string, data string) ([]byte, error) {

	//	println(data)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sign := common.Md5(appId + method + timestamp + appKey)

	params := make(url.Values)
	params.Add("app_id", appId)
	params.Add("method", method)
	params.Add("sign", sign)
	params.Add("ts", timestamp)
	params.Add("data", data)

	request, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := curl.Send(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.String()), nil
}

func requestPostGetAddress(address string) []byte {

	var (
		data        = strings.NewReader(`{"address":"` + address + `"}`)
		response, _ = curl.PostJson("https://wangzc.wang/smAddress", data)
	)
	return response.Byte()
}
