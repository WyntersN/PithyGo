/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-03-16 05:31:11
 * @LastEditTime: 2022-09-22 13:48:28
 * @FilePath: \PithyGo\addons\gopay\alipay\core\d.go
 */
package core

import (
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func NewClient() (*alipay.Client, error) {
	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	Client, err := alipay.NewClient("appId", "privateKey", true)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	// 打开Debug开关，输出日志，默认关闭
	//client.DebugSwitch = gopay.DebugOn

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = Client.SetCertSnByPath("./addons/gopay/alipay/crt/appCertPublicKey_2021003125612036.crt", "./addons/gopay/alipay/crt/alipayRootCert.crt", "./addons/gopay/alipay/crt/alipayCertPublicKey_RSA2.crt")
	if err != nil {
		return nil, err
	}

	Client = Client.SetCharset("utf-8") //.SetReturnUrl("https://www.fmm.ink").SetNotifyUrl("https://www.fmm.ink")

	return Client, nil

}
