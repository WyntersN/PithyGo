/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-11-26 06:06:26
 * @LastEditTime: 2021-11-26 06:08:12
 * @FilePath: \PithyGo\Import\Img.go
 */
package Import

import "github.com/ebar-go/curl"

func GetImg() string {
	req, _ := curl.Get("https://thispersondoesnotexist.com/image")
	return req.String()
}
