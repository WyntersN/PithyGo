/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-16 09:14:43
 * @LastEditTime: 2022-09-22 13:47:03
 * @FilePath: \PithyGo\addons\tencentyun\cos\api.go
 */
package cos

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var cosClient *cos.Client

const BaseURL = "https://cos-o-1313199494.cos.ap-shanghai.myqcloud.com"

func NewCos() {
	//f, _ := os.OpenFile("F:/Users/Wynters/GO/src/PithyGo/addons/tencentyun/log.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	u, _ := url.Parse(BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	cosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "", // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: "", // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			/*Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
				Writer:         f,
			},*/
		},
	})
}

func UploadNetPut(name string, fd io.Reader) error {
	_, err := cosClient.Object.Put(context.Background(), name, fd, nil)
	return err
}

func UploadFormPut(name string, fd *multipart.FileHeader) error {
	file, err := fd.Open()
	if err != nil {
		return err
	}
	_, err = cosClient.Object.Put(context.Background(), name, file, nil)
	defer file.Close()

	return err
}

// 2.通过本地文件上传对象
func UploadLoadPut(name string, path string) error {
	_, err := cosClient.Object.PutFromFile(context.Background(), name, path, nil)
	return err
}

func DowndownFile(key, local string) error {
	// 2. 下载对象到本地文件
	u, _ := url.Parse(key)
	_, err := cosClient.Object.GetToFile(context.Background(), u.Path, local, nil)
	return err
}
