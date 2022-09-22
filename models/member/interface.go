/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-19 16:18:36
 * @LastEditTime: 2022-08-23 17:39:44
 * @FilePath: \tborder_wxq\models\member\interface.go
 */
package member

import "PithyGo/models/structure"

type Interface interface {
	GetMyCash() structure.MemberCash
	CashApply(opt structure.MemberCash) error
	GetMemberByMID() (structure.MemberInfo, error)
}
type AdminInterface interface {
	GetMemberList(pageSize, page int, keyword string) ([]structure.MemberList, int64)
	GetMemberCash(pageSize, page int, keyword string) ([]structure.MemberCashList, int64)
}

type AdminContent struct {
	UserID uint
}
type Content struct {
	MemberID uint
}
