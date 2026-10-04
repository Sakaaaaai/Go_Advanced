package main

import (
	"fmt"

	// 実際のプロジェクトでは "gorm.io/gorm" をインポートしますが、
	// ここでは環境構築を省くため gorm のインターフェースっぽいものをモックとして定義します
)

// ========= モック環境（GORMのふりをするための定義） =========
type gormDB struct{}
type gormModel struct {
	ID        uint
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (db *gormDB) Create(value interface{}) *gormDB { return db }
func (db *gormDB) Where(query interface{}, args ...interface{}) *gormDB { return db }
func (db *gormDB) First(dest interface{}, conds ...interface{}) *gormDB { return db }
func (db *gormDB) Error() error { return nil }

// ==========================================================

// 1. gormModel (本来は gorm.Model) を埋め込んだ User 構造体を定義
type User struct {
	___       // gormModel を埋め込む
	Name string
	Age  int
}

// 2. 新しいユーザーを作成・保存する
func CreateUser(db *gormDB, name string, age int) error {
	user := User{Name: ___, Age: ___}
	// db.Create にポインタを渡して保存
	result := db.___(&user)
	return result.Error()
}

// 3. 名前でユーザーを検索・取得する
func GetUserByName(db *gormDB, name string) (User, error) {
	var user User
	// db.Where().First() メソッドチェーンを使って取得
	result := db.___("name = ?", ___).___(&user)
	return user, result.Error()
}

func main() {
	fmt.Println("This is a package for GORM logic.")
}
