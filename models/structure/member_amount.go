/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-17 13:15:57
 * @LastEditTime: 2022-08-17 13:24:58
 * @FilePath: \tborder_wxq\models\structure\member_amount.go
 */
package structure

import "gorm.io/gorm"

type MemberAmount struct {
	gorm.Model
	MemberId uint    `json:"member_id"` //False
	Type     uint    `json:"type"`      //False
	Amount   float32 `json:"amount"`    //False
	Reason   string  `json:"reason"`    //False
}

func AddMemberAmount(tx *gorm.DB, member_id, Type uint, amount float32, reason string) error {
	return tx.Create(&MemberAmount{MemberId: member_id, Type: Type, Amount: amount, Reason: reason}).Error
}
