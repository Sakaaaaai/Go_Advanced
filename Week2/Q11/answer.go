//go:build ignore

package main

import "fmt"

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
	Orders []Order
}

func GetUsersNPlus1(db *gormDB) []User {
	var users []User
	db.Find(&users)

	for i := range users {
		db.Model(&users[i]).Association("Orders").Find(&users[i].Orders)
	}
	return users
}

func GetUsersEager(db *gormDB) []User {
	var users []User
	db.Preload("Orders").Find(&users)
	return users
}

func main() {
	fmt.Println("This is a package for Eager Loading logic.")
}
