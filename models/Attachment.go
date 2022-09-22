/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-10 13:10:37
 * @LastEditTime: 2022-04-24 16:39:13
 * @FilePath: \PithyGo\models\Attachment.go
 */
package models

import (
	"PithyGo/models/structure"
	"PithyGo/service"
)

func GetAttachmentBySha1(sha1 string) structure.Attachment {
	var data structure.Attachment
	service.DB.Where("sha1 = ?", sha1).Find(&data)
	return data
}

func CreateAttachment(data structure.Attachment) error {
	return service.DB.Create(&data).Error
}
