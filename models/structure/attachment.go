/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-10 14:19:04
 * @LastEditTime: 2021-10-10 16:28:56
 * @FilePath: \PithyGo\models\structure\attachment.go
 */
package structure

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Size   uint   `json:"size"`
	Url    string `json:"url"`
	Path   string `json:"path"`
	Type   uint   `json:"type"`
	Sha1   string `json:"sha1"`
}
