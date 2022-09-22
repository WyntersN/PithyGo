/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-04-16 15:48:57
 * @LastEditTime: 2022-04-16 16:04:27
 * @FilePath: \PithyGo\addons\gopay\alipay\core\m.go
 */
package core

type DataBillEreceiptApplyResponse struct {
	Response DataBillEreceiptApply `json:"alipay_data_bill_ereceipt_apply_response"`
	Sign     string                `json:"sign"`
}
type DataBillEreceiptApply struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	FileId string `json:"file_id"`
}

type DataBillEreceiptQueryResponse struct {
	Response DataBillEreceiptQuery `json:"alipay_data_bill_ereceipt_query_response"`
	Sign     string                `json:"sign"`
}

type DataBillEreceiptQuery struct {
	Code         string `json:"code"`
	Msg          string `json:"msg"`
	Status       string `json:"status"`
	DownloadUrl  string `json:"download_url"`
	ErrorMessage string `json:"error_message"`
}
