/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-12-17 13:26:05
 * @LastEditTime: 2022-07-02 11:41:26
 * @FilePath: \PithyGo\addons\express\api.go
 */
package express

import (
	"PithyGo/addons/express/best"
	"PithyGo/addons/express/sffw"
	"PithyGo/addons/express/sto"
	"PithyGo/addons/express/typestruct"
	"PithyGo/addons/express/yto"
	"strings"

	"github.com/bitly/go-simplejson"
)

/**
 * @summary: 物流订阅
 * @description:
 * @param {*}
 * @Author: Wynters
 */
func TRACE_PLATFORM_SUBSCRIBE(cp_code string, waybillNo string) (*simplejson.Json, error) {
	var (
		req *simplejson.Json
		err error
	)
	switch strings.ToUpper(cp_code) {
	case "STO":
		req, err = sto.TRACE_PLATFORM_SUBSCRIBE(waybillNo)
	case "HTKY":
		req, err = best.TRACE_PLATFORM_SUBSCRIBE(waybillNo)
	case "SFFW":
		req, err = sffw.TRACE_PLATFORM_SUBSCRIBE(waybillNo)
	case "YTO":
		req, err = yto.TRACE_PLATFORM_SUBSCRIBE(waybillNo)
	}
	return req, err
}

func TRACE_QUERY(cp_code, ordersn string) (typestruct.TraceQueryStruct, error) {

	var (
		req typestruct.TraceQueryStruct
		err error
	)
	switch cp_code {
	case "HTKY":
		req, err = best.TRACE_QUERY(ordersn)
	case "SFFW":
		req, err = sffw.TRACE_QUERY(ordersn)
	case "YTO":
		req, err = yto.TRACE_QUERY(ordersn)
	}

	return req, err

}
