/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-10 13:13:26
 * @LastEditTime: 2022-09-13 18:24:52
 * @FilePath: \PithyGo\app\admin\model\Upload.go
 */
package model

import (
	"PithyGo/addons/tencentyun/cos"
	"PithyGo/common"
	"PithyGo/models"
	"PithyGo/models/structure"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/kataras/iris/v12"
)

func Upload(c iris.Context) (int, string, map[string]interface{}) {
	file, info, _ := c.FormFile("file")
	defer file.Close()
	var (
		err      error
		formData structure.Attachment
		null     = map[string]interface{}{"src": nil}
	)

	if fileTypeMatch(info.Filename) == false {
		return -1, "不支持该文件类型", null
	}
	//	defer file.Close()
	//获取后缀
	filesuffix := path.Ext(info.Filename)
	dest := filepath.Join(common.GetCurrentPath(), "/runtime/"+filesuffix[1:]+"/")
	if err := common.AutoCreatePath(dest); err != nil {
		return -1, err.Error(), null
	}

	body, _ := ioutil.ReadAll(file)

	fileSah1 := fileHash(body)

	formData = models.GetAttachmentBySha1(fileSah1)
	if formData != (structure.Attachment{}) {
		return 200, "success", map[string]interface{}{"src": formData.Url, "fileName": info.Filename}
	}

	timeStr := time.Now()
	var (
		yDate  = timeStr.Format("2006")
		mDate  = timeStr.Format("01")
		dyDate = timeStr.Format("02")
	)

	formData.Name = info.Filename
	formData.Size = uint(info.Size)
	info.Filename = fileNameReplace(info.Filename, common.GetNowRandOrder(6))
	formData.Path = "attachment/" + yDate + "/" + mDate + "/" + dyDate + "/" + info.Filename
	formData.Url = "https://cos-o-1313199494.cos.ap-shanghai.myqcloud.com/" + formData.Path
	formData.Sha1 = fileSah1
	formData.UserId = common.SessManager.Start(c).Get("USER_ID").(uint)

	if _, err = c.UploadFormFiles(dest); err != nil {
		return -1, err.Error(), null
	}
	if err = cos.UploadFormPut(formData.Path, info); err != nil {
		return -1, err.Error(), null
	}

	if err = models.CreateAttachment(formData); err != nil {
		return -1, err.Error(), null
	}

	return 200, "success", map[string]interface{}{"src": formData.Url, "fileName": info.Filename}

}

func fileNameReplace(fileName string, newName string) string {
	rep, _ := regexp.Compile(`[\s\S]+\.`)
	return rep.ReplaceAllString(fileName, newName+".")
}

//	func fileNameFind(fileName string) string {
//		rep, _ := regexp.Compile(`^(.*)\.png|jpg|jpeg|gif|mp4|bmp1$`)
//		str := rep.FindSubmatch([]byte(fileName))
//		return string(str[1])
//	}
func fileTypeMatch(fileName string) bool {
	rep, _ := regexp.Compile(`^(.*)\.png|jpg|jpeg|gif|mp4|bmp|xlsx|xls$`)
	return rep.MatchString(fileName)
}

// FileHash 求数据的MD5值
func fileHash(data []byte) string {
	m := sha1.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}
