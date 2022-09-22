/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-19 16:19:43
 * @LastEditTime: 2022-09-16 09:42:05
 * @FilePath: \PithyGo\models\member\cash.go
 */
package member

import (
	"PithyGo/common"
	"PithyGo/models/structure"
	"PithyGo/service"
	"errors"

	"gorm.io/gorm"
)

func (c *Content) GetMyCash() structure.MemberCash {
	var cash structure.MemberCash

	if service.DB.Last(&cash, "status < 1 AND member_id = ?", c.MemberID).RowsAffected == 0 {
		var member structure.Member
		service.DB.Select("realname,acc_alipay").First(&member, c.MemberID)
		return structure.MemberCash{
			Realname:       member.Realname,
			PaymentAccount: member.AccAlipay,
		}
	}

	return cash
}

func (c *Content) CashApply(opt structure.MemberCash) error {
	var member structure.Member
	if err := service.DB.First(&member, c.MemberID).Error; err != nil {
		return common.ReturnErrorToDB("查询用户失败", err)
	}

	if member.Realname == nil {
		service.DB.Model(&member).Update("realname", opt.Realname)
	}
	if member.AccAlipay == nil || *opt.PaymentAccount != *member.AccAlipay {
		service.DB.Model(&member).Update("acc_alipay", opt.PaymentAccount)
	}

	opt.MemberId = c.MemberID
	opt.Realname = member.Realname
	opt.PaymentAccount = member.AccAlipay

	if opt.PaymentAmount > member.Amount {
		return errors.New("提现金额超出了可提现余额")
	}
	//启动事务
	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			//return errors.New("出错啦")
		}
	}()
	opt.Ordersn = common.GetNowRandOrder(6)
	if err := tx.Create(&opt).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := structure.AddMemberAmount(tx, c.MemberID, 2, opt.PaymentAmount, "余额提现"); err != nil {
		tx.Rollback()
		return err
	}
	if res := tx.Model(&member).Where("amount >= ?", opt.PaymentAmount).UpdateColumn("amount", gorm.Expr("amount -?", opt.PaymentAmount)); res.Error != nil {
		tx.Rollback()
		return res.Error
	} else if res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("余额扣除失败")
	}
	return tx.Commit().Error
}
