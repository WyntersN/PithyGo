/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-04-19 17:42:13
 * @LastEditTime: 2022-04-19 17:56:28
 * @FilePath: \PithyGo\addons\tencentyun\ocr\api.go
 */
package ocr

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

func GeneralAccurate(image string) (string, error) {

	credential := common.NewCredential(
		"id",
		"key",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-shanghai", cpf)

	request := ocr.NewGeneralAccurateOCRRequest()

	request.ImageBase64 = common.StringPtr(image)

	response, err := client.GeneralAccurateOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", nil
	}
	return *response.Response.TextDetections[0].DetectedText, nil
}
