package SQL_translation

import (
	"errors"
	"fmt"
	"gorm-base/DO"
	"gorm.io/gorm"
)

/*
*题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func Run(db *gorm.DB) {
	createTable(db)
	insertRecord(db)
	transferMoney(db, 1, 2, 100)
}

func transferMoney(db *gorm.DB, fromId, toId uint, balance float64) {
	db.Transaction(func(tx *gorm.DB) error {
		var fromAccount DO.Accounts
		if err := tx.Select("accounts").Select("balance").Where("accounts.id = ?", fromId).Find(&fromAccount).Error; err != nil {
			panic(err.Error())
		}
		if fromAccount.Balance < 100 {
			return errors.New("余额不足")
		}
		result := tx.Model(&DO.Accounts{}).Where("id = ?", toId).UpdateColumn("Balance", gorm.Expr("balance+?", balance))
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("更新账户未成功")
		}
		result = tx.Model(&DO.Accounts{}).Where("id = ?", fromId).UpdateColumn("Balance", gorm.Expr("balance-?", balance))
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("更新账户未成功")
		}
		result = tx.Create(&DO.Transactions{FromAccountId: fromId, ToAccountId: toId, Amount: balance})
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("生成交易记录未成功")
		}

		fmt.Println("交易成功")
		return nil
	})
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func insertRecord(db *gorm.DB) {
	db.Create(&DO.Accounts{Name: "张三", Balance: 1000})
	db.Create(&DO.Accounts{Name: "李四", Balance: 1000})
	//db.Create(&DO.Transactions{FromAccountId: "李四", Age: 25, Grade: "三年级"})
}
func createTable(db *gorm.DB) {
	db.AutoMigrate(DO.Accounts{})
	db.AutoMigrate(DO.Transactions{})
}
