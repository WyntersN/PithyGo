/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-17 13:38:23
 * @LastEditTime: 2022-08-23 17:35:26
 * @FilePath: \tborder_wxq\models\structure\member_cash.go
 */
package structure

import (
	"time"

	"gorm.io/gorm"
)

type MemberCash struct {
	gorm.Model
	MemberId       uint       `json:"member_id"`       //False
	Realname       *string    `json:"realname"`        //False
	Ordersn        string     `json:"ordersn"`         //False
	PaymentMethod  uint       `json:"payment_method"`  //False
	PaymentAccount *string    `json:"payment_account"` //False
	PaymentAmount  float32    `json:"payment_amount"`  //False
	PaymentOrdersn *string    `json:"payment_ordersn"` //False
	PaymentTime    *time.Time `json:"payment_time"`    //False
	PaymentRemark  *string    `json:"payment_remark"`  //False
	PaymentError   *string    `json:"payment_error"`   //False
	Status         int        `json:"status"`          //False
}

type MemberCashList struct {
	Id             uint      `json:"id"`              //
	MemberId       uint      `json:"member_id"`       //False
	Realname       string    `json:"realname"`        //False
	Ordersn        string    `json:"ordersn"`         //False
	PaymentMethod  uint      `json:"payment_method"`  //False
	PaymentAccount string    `json:"payment_account"` //False
	PaymentAmount  float32   `json:"payment_amount"`  //False
	PaymentOrdersn string    `json:"payment_ordersn"` //False
	PaymentTime    time.Time `json:"payment_time"`    //False
	PaymentRemark  string    `json:"payment_remark"`  //False
	PaymentError   string    `json:"payment_error"`   //False
	Status         int       `json:"status"`          //False
	CreatedAt      time.Time `json:"createdAt"`       //False
}
