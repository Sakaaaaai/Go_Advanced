//go:build ignore

package main

import (
	"database/sql"
	"fmt"
)

func TransferFunds(tx *sql.Tx, fromAccount, toAccount string, amount int) error {
	_, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccount)
	if err != nil {
		return fmt.Errorf("fromAccountの更新に失敗: %w", err)
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccount)
	if err != nil {
		return fmt.Errorf("toAccountの更新に失敗: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("This is a package for Transaction logic. Run tests to verify the behavior.")
}
