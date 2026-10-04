package main

import "fmt"

// ========= モック環境 =========
type gormDB struct{}
type gormModel struct { ID uint }

func (db *gormDB) Find(dest interface{}, conds ...interface{}) *gormDB { return db }
func (db *gormDB) Preload(column string, conditions ...interface{}) *gormDB { return db }
func (db *gormDB) Model(value interface{}) *gormDB { return db }
func (db *gormDB) Association(column string) *gormAssociation { return &gormAssociation{} }

type gormAssociation struct{}
func (a *gormAssociation) Find(out interface{}, conds ...interface{}) error { return nil }

type Order struct {
	gormModel
	UserID uint
	Item   string
}

type User struct {
	gormModel
	Name   string
	// 1対多のリレーション。このユーザーが持つ複数の注文。
	Orders []Order
}
// =============================


// 悪い例（N+1問題）のシミュレーション
func GetUsersNPlus1(db *gormDB) []User {
	var users []User
	// 1回のクエリで全ユーザーを取得 (1)
	db.Find(&users)

	// forループの中で各ユーザーごとにクエリを投げてしまう (N)
	for i := range users {
		db.Model(&users[i]).Association("Orders").Find(&users[i].Orders)
	}
	return users
}

// 良い例（Eager Loading）
func GetUsersEager(db *gormDB) []User {
	var users []User
	
	// Preload を使って、取得時に "Orders" も事前に一括取得する
	db.___("___").___(&users)
	
	return users
}

func main() {
	fmt.Println("This is a package for Eager Loading logic.")
}
