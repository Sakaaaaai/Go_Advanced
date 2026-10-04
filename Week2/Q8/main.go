package main

import (
	"database/sql"
	"fmt"
)

// TransferFunds は、fromAccountからtoAccountへamount分の資金を移動します
func TransferFunds(tx *sql.___, fromAccount, toAccount string, amount int) error {
	// 1. fromAccountの残高を減らす
	_, err := tx.___("UPDATE accounts SET balance = balance - ? WHERE id = ?", ___, ___)
	if err != nil {
		return fmt.Errorf("fromAccountの更新に失敗: %w", err)
	}

	// 2. toAccountの残高を増やす
	_, err = tx.___("UPDATE accounts SET balance = balance + ? WHERE id = ?", ___, ___)
	if err != nil {
		return fmt.Errorf("toAccountの更新に失敗: %w", err)
	}

	return nil
}

// （参考）呼び出し元のイメージ
// func processTransfer(db *sql.DB) error {
//     tx, err := db.Begin()
//     if err != nil { return err }
//
//     err = TransferFunds(tx, "A", "B", 1000)
//     if err != nil {
//         tx.Rollback() // 失敗したら元に戻す
//         return err
//     }
//     return tx.Commit() // 成功したら確定する
// }

func main() {
	fmt.Println("This is a package for Transaction logic. Run tests to verify the behavior.")
}
