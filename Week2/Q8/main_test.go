package main

import (
	"database/sql"
	"testing"
)

// 実際のDB通信を回避するため、モックや構文チェックを簡易的に行うのは標準ライブラリだけでは難しいですが、
// syntaxが通るかどうかの最低限のテストを配置しておきます。
func TestTransferFundsSignature(t *testing.T) {
	// TransferFundsが存在し、正しい引数を取るかコンパイルレベルでのチェック
	var _ func(*sql.Tx, string, string, int) error = TransferFunds
}
