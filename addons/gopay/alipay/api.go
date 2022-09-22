/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-03-16 05:33:32
 * @LastEditTime: 2022-04-16 17:36:01
 * @FilePath: \PithyGo\addons\gopay\alipay\api.go
 */
package alipay

import (
	"PithyGo/addons/gopay/alipay/core"
	"context"
	"errors"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

var Client *alipay.Client

func FundTransUniTransfer(order, acc, name, remark string, amount float64) (*alipay.FundTransUniTransferResponse, error) {

	bm := make(gopay.BodyMap)

	bm.Set("out_biz_no", order).
		Set("trans_amount", amount).
		Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("remark", remark).
		SetBodyMap("payee_info", func(b gopay.BodyMap) {
			b.Set("identity", acc).Set("identity_type", "ALIPAY_LOGON_ID").Set("name", name)
		})

	aliRsp, err := Client.FundTransUniTransfer(context.Background(), bm)
	if err != nil {
		// xlog.Error("err:", err)
		// println(aliRsp.Response.Code)
		return nil, err
	}
	return aliRsp, nil
}

// 支付宝回执单下载
func DataBillEreceiptDown(fundOrderId string) (string, error) {

	var (
		bmEA = make(gopay.BodyMap)
		bmEQ = make(gopay.BodyMap)
	)

	// biz_content
	bmEA.SetBodyMap("biz_content", func(bz gopay.BodyMap) {
		bz.Set("type", "FUND_DETAIL")
		bz.Set("key", fundOrderId)
	})

	aliEAPsq := new(core.DataBillEreceiptApplyResponse)

	if err := Client.PostAliPayAPISelfV2(context.Background(), bmEA, "alipay.data.bill.ereceipt.apply", aliEAPsq); err != nil {
		return "", err
	}

	if condition := aliEAPsq.Response.Code; condition != "10000" {
		return "", errors.New(aliEAPsq.Response.Msg)
	}

	bmEQ.SetBodyMap("biz_content", func(bz gopay.BodyMap) {
		bz.Set("file_id", aliEAPsq.Response.FileId)
	})

	aliEQPsp := new(core.DataBillEreceiptQueryResponse)
	if err := Client.PostAliPayAPISelfV2(context.Background(), bmEQ, "alipay.data.bill.ereceipt.query", aliEQPsp); err != nil {
		return "", err
	}
	if aliEQPsp.Response.Code != "10000" {
		return "", errors.New(aliEQPsp.Response.Msg)
	}
	// if aliEQPsp.Response.DownloadUrl == "" {
	// 	return DataBillEreceiptDown(fundOrderId)
	// }
	return aliEQPsp.Response.DownloadUrl, nil
}
