/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-12-17 10:56:09
 * @LastEditTime: 2022-06-15 17:16:05
 * @FilePath: \PithyGo\addons\express\callback.go
 */
package express

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Callback struct {
	Ctx     iris.Context
	Session *sessions.Session
}
type fwfwFromData struct {
	MailNo              string                            `json:"mailNo"`
	BillTraceScanRecord []fwfwFromDataBillTraceScanRecord `json:"billTraceScanRecord"`
}

type fwfwFromDataBillTraceScanRecord struct {
	Remark        string `json:"remark"`
	AcceptAddress string `json:"acceptAddress"`
	AcceptTime    string `json:"acceptTime"`
	ScanType      string `json:"scanType"`
}

/*
func (c *Callback) PostBy(express_code string) {
	//轨迹推送
	//	service.LOG.Sugar().Info("-------------express-POST--------------", c.Ctx.PostValue("content"), express_code)
	var data structure.WmsExpress
	data.Code = strings.ToLower(express_code)
	switch express_code {
	case "sto":

		var (
			content = c.Ctx.PostValue("content")
			j, _    = simplejson.NewJson([]byte(content))
			trace   = j.Get("trace")
			json, _ = j.MarshalJSON()
			weight  string
		)

		//println(content)
		data.Name = "申通快递"
		data.WaybillNo = j.Get("waybillNo").MustString()
		data.OptTime, _ = time.Parse("2006-01-02 15:04:05", trace.Get("opTime").MustString())
		data.OptProvince = trace.Get("opOrgProvinceName").MustString()
		data.OptCity = trace.Get("opOrgCityName").MustString()
		data.Content = trace.Get("memo").MustString()
		//	data.Status, _ = strconv.Atoi(trace.Get("opOrgType").MustString("-1"))
		data.StatusStr = trace.Get("scanType").MustString("未知状态")
		data.OriginalData = string(json)
		data.Status = statusToStr(data.StatusStr)
		if data.Status == 1 {
			weight = trace.Get("weight").MustString("-1")
			data.Weight, _ = strconv.ParseFloat(weight, 64)
		}
	//	saveExpress(data)
		c.Ctx.Write([]byte(`{"success":true,"errorCode":"0","errorMsg":"","data":{"waybillNo":"` + data.WaybillNo + `","needRetry":false}}`))
		return
	case "htky":
		var (
			content                = c.Ctx.PostValue("bizData")
			j, _                   = simplejson.NewJson([]byte(content))
			billTraceScanRecord, _ = j.Get("billTraceScanRecord").Array()
			mailNo                 = j.Get("mailNo").MustString()
		)
		data.WaybillNo = mailNo
		data.Name = "百世快递"
		if data.WaybillNo != "" {
			for _, v := range billTraceScanRecord {
				var (
					billData = v.(map[string]interface{})
					jsons, _ = json.Marshal(billData)
				)
				data.ID = 0
				data.OriginalData = string(jsons)
				data.OptProvince = billData["acceptAddress"].(string)
				if billData["udf1"] != nil {
					data.OptCity = billData["udf1"].(string)
				}
				data.OptTime, _ = time.Parse("2006-01-02 15:04:05", billData["acceptTime"].(string))
				data.Content = billData["remark"].(string)
				data.StatusStr = billData["scanType"].(string)
				data.Status = statusToStr(data.StatusStr)
			//	saveExpress(data)
			}

			//go service.DB.Model(&structure.WmsOutDeliver{}).Where("express_code = ? AND express_no = ?", data.Code, data.WaybillNo).Update("express_status_str", data.StatusStr)

		}
		//service.LOG.Sugar().Info("-------------express-POST-ht---------------", c.Ctx.FormValues())
		c.Ctx.Write([]byte(`{"result":true,"remark":"","errorCode":null,"errorDescription":null}`))
		return
	case "sffw":
		// body, _ := c.Ctx.GetBody()
		// service.LOG.Sugar().Info("-------------express-POST-sffw---------------", string(body))
		var fwfwFromData fwfwFromData
		if err := c.Ctx.ReadJSON(&fwfwFromData); err != nil {
			return
		}
		data.WaybillNo = fwfwFromData.MailNo
		data.Name = "丰网速递"
		for _, v := range fwfwFromData.BillTraceScanRecord {
			data.ID = 0
			data.StatusStr = v.ScanType

			data.Content = v.Remark
			data.OptTime, _ = time.Parse("2006-01-02 15:04:05", v.AcceptTime)
			data.OptProvince = v.AcceptAddress
			data.OptCity = v.AcceptAddress + "转运中心"
			odJson, _ := json.Marshal(&v)
			data.OriginalData = string(odJson)
		//	saveExpress(data)

			c.Ctx.Write([]byte(`{"result":true,"errorCode":"0","errorDescription":"success"}`))
			return
		}

	case "yto":
		type ytoBack struct {
			XMLName            xml.Name `xml:"UpdateInfo"`
			LogisticProviderID string   `xml:"logisticProviderID"`
			ClientID           string   `xml:"clientID"`
			TxLogisticID       string   `xml:"txLogisticID"`
			MailNo             string   `xml:"mailNo"`
			Weight             float64  `xml:"weight"`
			InfoType           string   `xml:"infoType"`
			City               string   `xml:"city"`
			District           string   `xml:"district"`
			InfoContent        string   `xml:"infoContent"`
			AcceptTime         string   `xml:"acceptTime"`
			OrgName            string   `xml:"orgName"`
			OrgCode            string   `xml:"orgCode"`
			OrgPhone           string   `xml:"orgPhone"`
			EmpName            string   `xml:"empName"`
			EmpCode            string   `xml:"empCode"`
			Remark             string   `xml:"remark"`
		}
		var (
			logistics_interface = c.Ctx.PostValueTrim("logistics_interface")
			ytoBackData         ytoBack
		)
		err := xml.Unmarshal([]byte(logistics_interface), &ytoBackData)
		data.Name = "圆通速递"
		if err == nil {
			data.WaybillNo = ytoBackData.MailNo

			switch ytoBackData.InfoContent {
			case "GOT":
				data.Status = 1
				data.StatusStr = "已收件"
				data.Weight = ytoBackData.Weight
			case "NOT_SEND":
				data.Status = -1
				data.StatusStr = "揽收失败"
			case "ARRIVAL":
				data.Status = 3
				data.StatusStr = "已收入"
			case "DEPARTURE":
				data.Status = 4
				data.StatusStr = "已发出"
			case "PACKAGE":
				data.Status = 5
				data.StatusStr = "已打包"
			case "SENT_SCAN":
				data.Status = 6
				data.StatusStr = "派件"
			case "INBOUND":
				data.Status = 7
				data.StatusStr = "自提柜入柜"
			case "SIGNED":
				data.Status = 200
				data.StatusStr = "签收成功"
			case "FAILED":
				data.Status = -1
				data.StatusStr = "签收失败"
			default:
				data.Status = -1
				data.StatusStr = ytoBackData.InfoContent
			}
			data.OptTime, _ = time.Parse("2006-01-02 15:04:05", strings.ReplaceAll(ytoBackData.AcceptTime, " CST", ""))
			data.OptProvince = ytoBackData.City
			data.OptCity = ytoBackData.District
			data.Content = ytoBackData.Remark
			odJson, _ := json.Marshal(&ytoBackData)
			data.OriginalData = string(odJson)
		//	saveExpress(data)

			c.Ctx.Write([]byte(`<Response>
			<logisticProviderID>YTO</logisticProviderID>
			<txLogisticID>` + ytoBackData.TxLogisticID + `</txLogisticID>
			<success>true</success>
	  </Response>`))
			return
		}
		c.Ctx.Write([]byte(`<Response>
		<logisticProviderID>YTO</logisticProviderID>
		<txLogisticID>` + ytoBackData.TxLogisticID + `</txLogisticID>
		<success>false</success>
  </Response>`))
		return
	}

}

func statusToStr(statusStr string) int {

	switch statusStr {
	case "收件":
		return 1
	case "发件":
		return 2
	case "到件":
		return 3
	case "派件":
		return 4
	case "第三方代派":
		return 4
	case "驿站代收":
		return 5
	case "用户提货":
		return 5
	case "入库/入柜":
		return 5
	case "入柜/入库":
		return 5
	case "问题件":
		return 6
	case "退回件":
		return 7
	case "签收":
		return 200
	case "快件取出":
		return 200
	default:
		return -1
	}
}

// func saveExpress(data structure.WmsExpress) error {
// 	var expressData structure.WmsExpress
// 	service.DB.Where("code = ? AND waybill_no = ? AND opt_time = ? AND content = ? AND status_str = ?", data.Code, data.WaybillNo, data.OptTime, data.Content, data.StatusStr).Find(&expressData)

// 	if expressData != (structure.WmsExpress{}) {
// 		return nil
// 	}

// 	if service.DB.Create(&data).Error == nil {
// 		go service.DB.Model(&structure.WmsOutDeliver{}).Where("express_code = ? AND express_no = ?", data.Code, data.WaybillNo).Update("express_status_str", data.StatusStr)
// 		if data.Status == 1 && data.Weight > -1 && data.Code == "yto" {
// 			go service.DB.Model(&structure.WmsOutDeliver{}).Where("express_code = ? AND express_no = ?", data.Code, data.WaybillNo).Update("express_weight", data.Weight)
// 		}
// 		if data.Status == 1 && data.Weight > -1 && data.Code == "sto" {
// 			go service.DB.Model(&structure.WmsOutDeliver{}).Where("express_code = ? AND express_no = ?", data.Code, data.WaybillNo).Update("express_weight", data.Weight*1000)
// 		}
// 	}
// 	return nil
// }

// func (c *Callback) GetBy(express_code string) {
// 	service.LOG.Sugar().Info("-------------express-GET--------------", c.Ctx.FormValues())
// }
*/
