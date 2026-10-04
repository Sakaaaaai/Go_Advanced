# 問8｜高度なdatabase/sqlとトランザクション

## 目標
`database/sql` を使って複数のDB操作をアトミックに（全成功または全失敗に）扱う「トランザクション」を学びます。

## 今回の演習で学ぶこと
この問題では、複数のデータベース操作を「ひとまとめ」にして、データの整合性を守る「トランザクション」の実装方法を、以下の3つの重要ポイントに注目して学びます。

**1. tx.Exec による更新処理（送金元）**
トランザクション（取引の一連の処理）の中で、安全にクエリを実行する手順です。
```go
_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, from)
if err != nil {
    return err // ここで失敗すれば、即座に中止を報告する
}
```
- `tx.Exec`: 通常の `db.Exec` ではなく、トランザクション専用の変数 `tx` を使ってSQLを発行します。これにより「この更新はまだ仮の状態で止めておいてね」とDBに指示できます。
- まずは「Aさんの口座からお金を引く」という第一段階の処理を行います。

**2. tx.Exec による更新処理（送金先）**
一連の操作の続きとして、絶対に行わなければならない次の更新です。
```go
_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, to)
if err != nil {
    return err // ここで失敗した場合も、即座に中止を報告
}
```
- Aさんから引いたお金を「Bさんの口座に足す」という処理です。
- もし仮にここで「Bさんの口座が存在しない」などのエラーが発生した場合でも、呼び出し元で設定された「安全網（Rollback）」が発動し、Aさんのお金だけが減ったままになるという最悪の事態（データの不整合）を防ぎます。

**3. コミットとロールバックの流れ**
データの確定（Commit）と取り消し（Rollback）を制御する、システム保護の「安全網」です。
```go
// 処理の最初に「もしダメだったら全部なかったことにしてね」と予約しておく
defer tx.Rollback()

// ... 全ての処理が【成功】したら
return tx.Commit()
```
- `defer tx.Rollback()`: Go言語ならではの超強力なイディオムです。途中でエラーが起きて `return` したり、まさかのパニック（強制終了）が発生しても、関数を抜ける瞬間に必ずこのRollbackが呼ばれ、変更が取り消されます。
- `tx.Commit()`: この行に到達したということは「全て成功した」証拠です。ここで初めてDBに変更を「確定」させます（Commitが成功した後のRollbackは無視されるため安全です）。




**ACIDの原則**
トランザクションを使うことで、システムが不意に停止しても、お金が増えたり消えたりするようなデータ不整合（整合性の破壊）を確実に防ぐことができます。

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

`TransferFunds(tx *sql.Tx, fromAccount, toAccount string, amount int) error` を完成させてください：
1. `tx.Exec` で `fromAccount` の残高を `amount` 分減らしてください。
2. 同様に `toAccount` の残高を `amount` 分増やしてください。
3. エラー発生時はすぐにエラーを返してください（ロールバックは呼び出し元で行う想定）。

### 実行方法
```bash
cd /workspaces/Go_Advanced/Week2/Q8
go run main.go
```

### テストでの確認
```bash
cd /workspaces/Go_Advanced/Week2/Q8
go test main.go main_test.go
```
## ヒント
- `tx.Exec("UPDATE ... SET balance = balance - ? WHERE id = ?", amount, fromAccount)` の形
- 関数がエラーを返したら呼び出し元で `tx.Rollback()` を実行する設計

## 達成基準
- [ ] トランザクションの必要性（ACID）を説明できる
- [ ] `db.Begin() / tx.Exec / tx.Commit / tx.Rollback` の流れを実装できる
- [ ] `defer tx.Rollback()` パターンを使った安全な実装ができる
