/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-03-18 18:32:59
 * @LastEditTime: 2022-05-12 17:31:10
 * @FilePath: \PithyGo\addons\gopay\init.go
 */
package gopay

import (
	"PithyGo/addons/gopay/alipay"
	"PithyGo/addons/gopay/alipay/core"
	"PithyGo/service"
)

func AutoLoad() {
	var aliErr error
	alipay.Client, aliErr = core.NewClient()

	if aliErr != nil {
		service.LOG.Sugar().Fatal("%s", aliErr)
	}

}
