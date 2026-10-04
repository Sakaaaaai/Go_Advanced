# 問10｜ORMの活用（GORM）

## 目標
GoのORM（Object-Relational Mapper）ライブラリ `GORM` を使い、生のSQLを書かずに構造体でデータベース操作を行う方法を学びます。

## 今回の演習で学ぶこと
この問題では、Goの代表的なORM（Object-Relational Mapper）である「GORM」を使い、生のSQLを書かずにデータベースを操作する方法を、以下の3つの重要ポイントに注目して学びます。

**1. モデル構造体の定義と gorm.Model**
データベースの「テーブル」の形を、Go言語の「構造体」の形として定義し、橋渡しをする工程です。
```go
type User struct {
    gorm.Model // IDや作成日時などの便利フィールドを自動追加
    Name string
    Age  int
}
```
- GoでORMを使う最大のメリットは、「いちいちSQLのテーブル定義を書かなくても、構造体を定義するだけでデータの形を取り決められる」点にあります。
- `gorm.Model`: GORMが用意してくれている便利な魔法のパーツです。これを構造体に一行埋め込むだけで、自動的に `ID`（主キー）、`CreatedAt`（作成日時）、`UpdatedAt`（更新日時）、`DeletedAt`（論理削除用）という必須級のフィールドが全て備わります。

**2. db.Create によるレコード作成**
Goのコード上で作成したデータを、魔法のようにデータベースへ保存する操作です。
```go
newU := User{Name: "Go太郎", Age: 25} // Goの世界でデータを作る
db.Create(&newU) // そのデータをDBの世界へ保存する
```
- 開発者が `INSERT INTO users (name, age) VALUES ('Go太郎', 25)` といった生のSQLを組み立てる必要はありません。
- `db.Create(&newU)` と関数を呼ぶだけで、GORMが裏側で自動的に最適なSQLを生成し、実行してくれます。引数には必ず変数のアドレス（`&`）を渡します。

**3. メソッドチェーンによるデータ取得（Where / First）**
条件を指定して、欲しいデータをデータベースから引っ張り出す工程です。
```go
var found User // データを受け取るための「空の箱」用意
db.Where("name = ?", "Go太郎").First(&found)
```
- `Where(...)`: SQLの `WHERE name = 'Go太郎'` の部分に相当します。`?`（プレースホルダ）を使うことで、SQLインジェクションというセキュリティ攻撃を自動的に防いでくれます。
- `First(&found)`: 条件に一致する「最初の1件」を取得し、用意した空の箱（`&found`）にデータを詰め込みます。このように関数をどんどん繋げていく書き方を「メソッドチェーン」と呼び、直感的で読みやすいコードになります。




**GORM vs 生SQL の使い分け**
日々の開発やシンプルなCRUD操作にはGORMが圧倒的に便利ですが、非常に複雑な結合（Join）やパフォーマンスが最優先される箇所では、あえて生のSQLを書くという選択肢も重要です。状況に応じて賢く使い分けましょう。

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `gorm.Model` を埋め込んだ `User` 構造体を `Name string` と `Age int` フィールドで定義してください。
2. `CreateUser(db, name, age)` 関数：`db.Create` で新しいユーザーを作成してください。
3. `GetUserByName(db, name)` 関数：`db.Where` と `First` でユーザーを取得してください。

### 実行方法
```bash
cd /workspaces/Go_Advanced/Week2/Q10
go run main.go
```

### テストでの確認
```bash
cd /workspaces/Go_Advanced/Week2/Q10
go test main.go main_test.go
```
## ヒント
- `gorm.Model` を埋め込むと `ID`, `CreatedAt` 等が自動付与される
- `db.Where("name = ?", name).First(&user)` でSQL の WHERE 句を書ける

## 達成基準
- [ ] GORM の基本的なCreate/Readが書ける
- [ ] `gorm.Model` の意味と提供するフィールドを説明できる
- [ ] 生SQL と ORM の使い分けを説明できる
