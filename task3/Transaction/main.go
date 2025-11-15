package main

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Accounts 映射数据库中的 accounts 表，表示账户信息
type Accounts struct {
	Id      int `gorm:"primaryKey"` // 账户ID，主键
	Name    string                // 账户名称
	Balance int                   // 账户余额
}

// Transactions 映射数据库中的 transactions 表，表示转账交易记录
type Transactions struct {
	Id              int `gorm:"primaryKey"` // 交易ID，主键
	From_account_id int                    // 转出账户ID
	To_account_id   int                    // 转入账户ID
	Amount          int                    // 转账金额
}

// Transfer 实现转账功能的事务操作
// 参数:
//   - db: 数据库连接对象
//   - fromAccountName: 转出账户名称
//   - toAccountName: 转入账户名称
//   - amount: 转账金额
// 返回值:
//   - error: 如果转账过程中出现错误则返回相应的错误信息，否则返回nil
func Transfer(db *gorm.DB, fromAccountName string, toAccountName string, amount int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Accounts

		// 查询转出账户
		if err := tx.Where("name = ?", fromAccountName).First(&fromAccount).Error; err != nil {
			return err
		}

		// 查询转入账户
		if err := tx.Where("name = ?", toAccountName).First(&toAccount).Error; err != nil {
			return err
		}

		// 检查余额是否足够
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}

		// 扣除转出账户余额
		fromAccount.Balance -= amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}

		// 增加转入账户余额
		toAccount.Balance += amount
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}

		// 记录转账交易
		transaction := Transactions{
			From_account_id: fromAccount.Id,
			To_account_id:   toAccount.Id,
			Amount:          amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

// addRecord 添加初始测试数据到accounts表中
func addRecord(db *gorm.DB) {
	db.Create(&Accounts{Name: "A", Balance: 600})
	db.Create(&Accounts{Name: "B", Balance: 800})
}

func main() {
	dsn := "root:123456@tcp(192.168.100.121:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	if err != nil {
		fmt.Println("数据库连接失败") // 打印错误信息
	}

	// db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})
	// addRecord(db)

	err = Transfer(db, "A", "B", 100)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功")
	}

	// 验证结果
	var accountA, accountB Accounts
	db.Where("name = ?", "A").First(&accountA)
	db.Where("name = ?", "B").First(&accountB)
	fmt.Printf("账户A余额: %d\n", accountA.Balance)
	fmt.Printf("账户B余额: %d\n", accountB.Balance)
}