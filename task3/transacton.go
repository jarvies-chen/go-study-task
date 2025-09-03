// 题目2：事务语句
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint
	Balance float64
}

type Transaction struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

// Transfer 转账函数 - 使用事务
func Transfer(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询转出账户
		var fromAccount Account
		if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
			return fmt.Errorf("查询转出账户失败: %w", err)
		}

		// 2. 检查余额是否足够
		if fromAccount.Balance < amount {
			return fmt.Errorf("账户 %d 余额不足，当前余额: %.2f，转账金额: %.2f",
				fromAccountID, fromAccount.Balance, amount)
		}

		// 3. 查询转入账户
		var toAccount Account
		if err := tx.First(&toAccount, toAccountID).Error; err != nil {
			return fmt.Errorf("查询转入账户失败: %w", err)
		}

		// 4. 更新转出账户余额
		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error; err != nil {
			return fmt.Errorf("更新转出账户余额失败: %w", err)
		}

		// 5. 更新转入账户余额
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error; err != nil {
			return fmt.Errorf("更新转入账户余额失败: %w", err)
		}

		// 6. 记录转账信息
		transaction := Transaction{
			FromAccountId: fromAccountID,
			ToAccountId:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("记录转账信息失败: %w", err)
		}

		fmt.Printf("转账成功: 账户 %d 向账户 %d 转账 %.2f 元\n",
			fromAccountID, toAccountID, amount)
		fmt.Printf("账户 %d 余额: %.2f -> %.2f\n",
			fromAccountID, fromAccount.Balance, fromAccount.Balance-amount)
		fmt.Printf("账户 %d 余额: %.2f -> %.2f\n",
			toAccountID, toAccount.Balance, toAccount.Balance+amount)
		return nil
	})
}

// 初始化账户数据
func initAccounts(db *gorm.DB) error {
	accounts := []Account{
		{ID: 1, Balance: 1000},
		{ID: 2, Balance: 500},
	}

	for _, account := range accounts {
		var existing Account
		result := db.Where("id = ?", account.ID).First(&existing)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&account).Error; err != nil {
					return err
				}
				fmt.Printf("创建账户 %d，初始余额: %.2f\n", account.ID, account.Balance)
			} else {
				return result.Error
			}
		}
	}
	return nil
}

// 打印所有账户信息
func printAccounts(db *gorm.DB) {
	var accounts []Account
	db.Find(&accounts)
	for _, account := range accounts {
		fmt.Printf("账户 %d: %.2f 元\n", account.ID, account.Balance)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}
	fmt.Println("数据库连接成功")

	db.AutoMigrate(&Account{}, &Transaction{})

	//init DB
	initAccounts(db)
	printAccounts(db)

	// 测试转账 - 正常情况
	fmt.Println("\n=== 测试转账 100 元 ===")
	err = Transfer(db, 1, 2, 100.00)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功")
	}
	// 显示转账后账户余额
	printAccounts(db)

	// 测试转账 - 余额不足的情况
	fmt.Println("\n=== 测试余额不足的转账 ===")
	err = Transfer(db, 1, 2, 2000.00) // 尝试转账2000元
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功")
	}
	// 显示转账后账户余额
	printAccounts(db)

	// 显示转账记录
	fmt.Println("\n=== 转账记录 ===")
	var transactions []Transaction
	db.Find(&transactions)
	if len(transactions) == 0 {
		fmt.Println("暂无转账记录")
	} else {
		for _, tx := range transactions {
			fmt.Printf("ID: %d, 从账户 %d 向账户 %d 转账 %.2f 元\n",
				tx.ID, tx.FromAccountId, tx.ToAccountId, tx.Amount)
		}
	}
}
